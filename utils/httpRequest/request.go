package httprequest

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/viper"
)

func regenerateToken(configFile string) {
	token, err := HTTPCreateToken(viper.GetString("client_id"), viper.GetString("client_secret"), viper.GetString("grant_type"), viper.GetString("scope"), viper.GetInt("expiration"))
	if err != nil {
		log.Fatalf("%s", err)
	}
	if token == "" {
		log.Fatal("client_id or client_secret not valid")
	} else {
		fmt.Fprintln(os.Stdout, "Token generated successfully")
		config.WriteToken(configFile, token)
	}
}

var c = http.Client{Timeout: time.Duration(10) * time.Second}
var counter = false

type PageResult struct {
	Items      json.RawMessage `json:"items"`
	TotalCount int             `json:"total_count"`
}

func HTTPRequest(method string, resource string, body []byte) ([]byte, error) {
	var bodyIO io.Reader = nil
	if body != nil {
		bodyIO = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, resource, bodyIO)
	if err != nil {
		log.Panicf("error occurred on request creation: %v", err)
	}

	if !strings.Contains(resource, "token") && viper.GetString("account_id") == "" && viper.GetString("account_environment_id") == "" {
		log.Fatalf("account_id or account_environment_id required, Please configure your CLI")
	}

	if strings.Contains(resource, "token") && viper.GetString("client_id") == "" && viper.GetString("client_secret") == "" {
		log.Fatalf("client_id or client_secret required, Please configure your CLI")
	}

	if !strings.Contains(resource, "token") && viper.GetString("token") == "" {
		regenerateToken(config.CredentialsFile)
	}

	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Authorization", "Bearer "+viper.GetString("token"))
	req.Header.Add("Accept-Encoding", `gzip, deflate, br`)
	if body != nil {
		req.Header.Add("Content-Type", `application/json`)
	}

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
	respBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 && resp.StatusCode != 403 {
		err = errors.New(string(respBody))
	}
	if resp.StatusCode == 403 && !counter {
		counter = true
		regenerateToken(config.CredentialsFile)
		return HTTPRequest(method, resource, body)
	}
	return respBody, err
}

func HTTPGetItem[T any](resource string) (T, error) {
	var result T
	respBody, err := HTTPRequest(http.MethodGet, resource, nil)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(respBody, &result)
	return result, err
}

func HTTPGetAllPages[T any](resource string) ([]T, error) {
	currentPage := 1
	results := []T{}
	for {
		respBody, err := HTTPRequest(http.MethodGet, fmt.Sprintf("%s?_page=%d&_max_per_page=100", resource, currentPage), nil)
		if err != nil {
			return nil, err
		}
		pageResult := &PageResult{}
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
