package common

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
)

var UserAgent string

var c = http.Client{Timeout: time.Duration(10) * time.Second}
var counter = false

type PageResultFE struct {
	Items      json.RawMessage `json:"items"`
	TotalCount int             `json:"total_count"`
}

type ResourceRequest struct {
	AccountID            string `mapstructure:"account_id"`
	AccountEnvironmentID string `mapstructure:"account_environment_id"`
}

func (c *ResourceRequest) Init(cL *RequestConfig) {
	c.AccountEnvironmentID = cL.AccountEnvironmentID
	c.AccountID = cL.AccountID
}

type PageResultWE struct {
	Data       json.RawMessage `json:"_data"`
	Pagination PaginationWE    `json:"_pagination"`
}

type PaginationWE struct {
	Total      int `json:"_total"`
	Pages      int `json:"_pages"`
	Page       int `json:"_page"`
	MaxPerPage int `json:"_max_per_page"`
}

type RequestConfig struct {
	Product               string
	Username              string `mapstructure:"username"`
	ClientID              string `mapstructure:"client_id"`
	ClientSecret          string `mapstructure:"client_secret"`
	AccountID             string `mapstructure:"account_id"`
	AccountEnvironmentID  string `mapstructure:"account_environment_id"`
	Token                 string `mapstructure:"token"`
	Scope                 string `mapstructure:"scope"`
	RefreshToken          string `mapstructure:"refresh_token"`
	CurrentUsedCredential string `mapstructure:"current_used_credential"`
	OutputFormat          string `mapstructure:"output_format"`
}

var cred RequestConfig

func Init(credL RequestConfig) {
	cred = credL
}

func regenerateToken(product, configName string) {
	var authenticationResponse models.TokenResponse
	var err error

	if product == utils.FEATURE_EXPERIMENTATION {
		authenticationResponse, err = HTTPRefreshTokenFE(cred.ClientID, cred.RefreshToken)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
	} else {
		authenticationResponse, err = HTTPCreateTokenWE(cred.ClientID, cred.ClientSecret, cred.AccountID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
	}

	if authenticationResponse.AccessToken == "" {
		log.Fatal("client_id or client_secret not valid")
	}

	cred.RefreshToken = authenticationResponse.RefreshToken
	cred.Token = authenticationResponse.AccessToken
	err = config.RewriteToken(product, configName, authenticationResponse)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

}

func HTTPRequest[T any](method string, url string, body []byte) ([]byte, error) {
	var bodyIO io.Reader = nil
	if body != nil {
		bodyIO = bytes.NewBuffer(body)
	}

	var resource T

	resourceType := reflect.TypeOf(resource)

	if resourceType == reflect.TypeOf(feature_experimentation.Goal{}) || resourceType == reflect.TypeOf(feature_experimentation.CampaignFE{}) {
		if cred.AccountID == "" || cred.AccountEnvironmentID == "" {
			log.Fatalf("account_id or account_environment_id required, Please authenticate your CLI")
		}
	}

	req, err := http.NewRequest(method, url, bodyIO)
	if err != nil {
		log.Panicf("error occurred on request creation: %v", err)
	}

	if cred.Product == utils.FEATURE_EXPERIMENTATION {
		if (cred.Username == "" || cred.AccountID == "") && resourceType != reflect.TypeOf(models.Token{}) {
			log.Fatalf("username and account_id required, Please authenticate your CLI")
		}
		// for resource loader
		if resourceType.String() == "resource.ResourceData" && !strings.Contains(url, "token") && (cred.AccountID == "" || cred.AccountEnvironmentID == "") {
			log.Fatalf("account_id or account_environment_id required, Please authenticate your CLI")
		}

		/* 		if strings.Contains(url, "token") && cred.ClientID == "" && cred.ClientSecret == "" {
			log.Fatalf("client_id or client_secret required, Please authenticate your CLI")
		} */
	}

	if cred.Product == utils.WEB_EXPERIMENTATION {
		if resourceType != reflect.TypeOf(web_experimentation.AccountWE{}) && !strings.Contains(url, "token") && cred.AccountID == "" {
			log.Fatalf("username, account_id required, Please authenticate your CLI")
		}
	}

	if !strings.Contains(url, "token") && cred.Token == "" {
		regenerateToken(cred.Product, cred.CurrentUsedCredential)
	}

	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Authorization", "Bearer "+cred.Token)
	req.Header.Add("Accept-Encoding", `gzip, deflate, br`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}
	respBody, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if (resp.StatusCode == 401) && !counter {
		counter = true
		regenerateToken(cred.Product, cred.CurrentUsedCredential)
		return HTTPRequest[T](method, url, body)
	}

	match, _ := regexp.MatchString("4\\d\\d|5\\d\\d", resp.Status)
	if match {
		err := errors.New(string(respBody))
		return nil, err
	}

	return respBody, err
}

func HTTPGetItem[T any](resource string) (T, error) {
	var result T
	respBody, err := HTTPRequest[T](http.MethodGet, resource, nil)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(respBody, &result)
	return result, err
}

func HTTPGetAllPagesFE[T any](resource string) ([]T, error) {
	currentPage := 1
	results := []T{}
	for {
		respBody, err := HTTPRequest[T](http.MethodGet, fmt.Sprintf("%s?_page=%d&_max_per_page=100", resource, currentPage), nil)
		if err != nil {
			return nil, err
		}
		pageResult := &PageResultFE{}
		err = json.Unmarshal(respBody, pageResult)
		if err != nil {
			return nil, err
		}

		typedItems := []T{}
		err = json.Unmarshal(pageResult.Items, &typedItems)
		if err != nil {
			return nil, err
		}
		results = append(results, typedItems...)

		if len(results) >= pageResult.TotalCount || len(pageResult.Items) == 0 {
			break
		}
		currentPage++
	}
	return results, nil
}

func HTTPGetAllPagesWE[T any](resource string) ([]T, error) {
	currentPage := 1
	results := []T{}
	for {
		respBody, err := HTTPRequest[T](http.MethodGet, fmt.Sprintf("%s?_page=%d&_max_per_page=100", resource, currentPage), nil)
		if err != nil {
			return nil, err
		}
		pageResult := &PageResultWE{}
		err = json.Unmarshal(respBody, pageResult)
		if err != nil {
			return nil, err
		}

		typedItems := []T{}
		err = json.Unmarshal(pageResult.Data, &typedItems)
		if err != nil {
			return nil, err
		}
		results = append(results, typedItems...)

		if len(results) >= pageResult.Pagination.Total || len(pageResult.Data) == 0 {
			break
		}
		currentPage++
	}
	return results, nil
}
