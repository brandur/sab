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

type history struct {
	Name *string `json:"name"`
	Size int     `json:"bytes"`
}

type histories struct {
	Inner *historiesInner `json:"history"`
}

type historiesInner struct {
	Histories []*history `json:"slots"`
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
	apiKey     string
	httpClient *http.Client
	url        string
}

type status struct {
	State *string `json:"state"`
}

func newSabClient(url string, apiKey string) *sabClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &sabClient{
		apiKey:     apiKey,
		httpClient: &http.Client{Transport: transport},
		url:        url,
	}
}

func (c *sabClient) buildApiUrl(mode string, extra string) string {
	url := fmt.Sprintf("%v/api?mode=%v&output=json&apikey=%v&%v", c.url, mode, c.apiKey, extra)
	printDebug("url: %v", url)
	return url
}

func (c *sabClient) getHistories(limit int) ([]*history, error) {
	url := c.buildApiUrl("history", fmt.Sprintf("limit=%v", limit))
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := readAndCheck(resp)
	if err != nil {
		return nil, err
	}

	var h histories
	err = json.Unmarshal(data, &h)
	if err != nil {
		return nil, err
	}

	return h.Inner.Histories, nil
}

func (c *sabClient) getJobs() ([]*job, error) {
	url := c.buildApiUrl("qstatus", "")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := readAndCheck(resp)
	if err != nil {
		return nil, err
	}

	var j jobs
	err = json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	return j.Jobs, nil
}

func (c *sabClient) getStatus() (*status, error) {
	url := c.buildApiUrl("qstatus", "")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := readAndCheck(resp)
	if err != nil {
		return nil, err
	}

	var s status
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func readAndCheck(resp *http.Response) ([]byte, error) {
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
	return data, nil
}
