package httprequest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var c = http.Client{Timeout: time.Duration(10) * time.Second}

func HTTPRequest(method string, resource string, body []byte) ([]byte, error) {
	var bodyIO io.Reader = nil
	if body != nil {
		bodyIO = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest("DELETE", resource, bodyIO)
	if err != nil {
		log.Panicf("error occured on request creation: %v", err)
	}

	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	req.Header.Add("Accept-Encoding", `gzip, deflate, br`)
	if body != nil {
		req.Header.Add("Content-Type", `application/json`)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		err = fmt.Errorf("error occured when calling request: %v", resp.StatusCode)
	}
	return respBody, err
}
