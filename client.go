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

type job struct {
	Filename *string `json:"filename"`
	Size     float64 `json:"mb"`
	SizeLeft float64 `json:"mbleft"`
	TimeLeft *string `json:"timeleft"`
}

type jobs struct {
	Jobs []*job `json:"jobs"`
}

type sabClient struct {
	apiKey     *string
	httpClient *http.Client
	url        *string
}

type status struct {
	State *string `json:"state"`
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

func (c *sabClient) getJobs() ([]*job, error) {
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

	var j jobs
	err = json.Unmarshal(data, &j)
	return j.Jobs, nil
}

func (c *sabClient) getStatus() (*status, error) {
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

	var s status
	err = json.Unmarshal(data, &s)
	return &s, nil
}