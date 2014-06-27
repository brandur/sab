package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiError struct {
	Message *string `json:"error"`
}

type sabClient struct {
	apiKey     *string
	httpClient *http.Client
	url        *string
}

func newSabClient(url *string, apiKey *string) *sabClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &sabClient{
		apiKey:     apiKey,
		httpClient: &http.Client{Transport: transport},
		url:        url,
	}
}

func (c *sabClient) getJobs() (map[string]interface{}, error) {
	url := fmt.Sprintf("%v/api?mode=qstatus&output=json&apikey=%v", *c.url, *c.apiKey)
	printDebug("url: %v", url)

	resp, err := c.httpClient.Get(url)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var e *apiError
	err = json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}

	if e.Message != nil {
		return nil, errors.New(*e.Message)
	}

	var thing map[string]interface{}
	err = json.Unmarshal(data, &thing)
	return thing, nil
}
