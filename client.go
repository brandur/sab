package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// used for pause, resume, and shutdown
type action struct {
	Status bool `json:"status"`
}

type apiError struct {
	Message *string `json:"error"`
}

type history struct {
	Name   *string `json:"name"`
	Size   int     `json:"bytes"`
	Status *string `json:"status"`
}

type histories struct {
	Inner *historiesInner `json:"history"`
}

type historiesInner struct {
	Histories []*history `json:"slots"`
}

type job struct {
	Filename   *string `json:"filename"`
	Percentage *string `json:"percentage"`
	Size       *string `json:"mb"`
	Status     *string `json:"status"`
}

type jobs struct {
	Inner *jobsInner `json:"queue"`
}

type jobsInner struct {
	Jobs []*job `json:"slots"`
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
	var h histories
	err := c.apiCall("history", fmt.Sprintf("limit=%v", limit), &h)
	return h.Inner.Histories, err
}

func (c *sabClient) getJobs() ([]*job, error) {
	var j jobs
	err := c.apiCall("queue", "", &j)
	return j.Inner.Jobs, err
}

func (c *sabClient) getStatus() (*status, error) {
	var s status
	err := c.apiCall("qstatus", "", &s)
	return &s, err
}

func (c *sabClient) pause() (*action, error) {
	var a action
	err := c.apiCall("pause", "", &a)
	return &a, err
}

func (c *sabClient) apiCall(mode string, extra string, t interface{}) error {
	url := c.buildApiUrl(mode, extra)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var e *apiError
	err = json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	printDebug("error check: %+v", t)

	if e.Message != nil {
		return errors.New(*e.Message)
	}

	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	printDebug("response: %+v", t)

	return nil
}
