package httprequest

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var c = http.Client{Timeout: time.Duration(10) * time.Second}

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
		log.Panicf("error occured on request creation: %v", err)
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
	if resp.StatusCode >= 400 {
		err = fmt.Errorf("error occured when calling request: %v", resp.StatusCode)
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
