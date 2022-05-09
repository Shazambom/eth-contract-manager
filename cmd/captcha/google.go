package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GoogleResponse struct {
	Event GoogleEvent `json:"event"`
	Name string `json:"name"`
	Reasons []string `json:"reasons"`
	Score float64 `json:"score"`
	Token GoogleTokenProps `json:"tokenProperties"`

}

type GoogleTokenProps struct {
	Action string `json:"action"`
	CreateTime time.Time `json:"createTime"`
	HostName string `json:"hostname"`
	InvalidReason string `json:"invalidReason"`
	Valid bool `json:"valid"`
}

type GoogleEvent struct {
	Token string `json:"token"`
	SiteKey string `json:"siteKey"`
	ExpectedAction string `json:"expectedAction"`
	IPAddress string `json:"userIpAddress"`
	UserAgent string `json:"userAgent"`
}

type GoogleReq struct {
	Event GoogleEvent `json:"event"`
}



func VerifyCaptcha(gURL, projectID, apiKey, siteKey, token, IP, userAgent, referer string) (*GoogleResponse, error) {
	url := gURL + "projects/" + projectID + "/assessments?key=" + apiKey

	gReq := GoogleReq{GoogleEvent{
		Token:          token,
		SiteKey:        siteKey,
		ExpectedAction: "challenge",
		IPAddress: IP,
		UserAgent: userAgent,
	}}

	marshGReq, marshalErr := json.Marshal(gReq)
	if marshalErr != nil {
		return nil, marshalErr
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshGReq))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", referer)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		fmt.Println(reqErr.Error())
		return nil, reqErr
	}
	defer resp.Body.Close()

	body, parseErr := ioutil.ReadAll(resp.Body)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return nil, parseErr
	}

	var googleResponse GoogleResponse
	unmarshalErr := json.Unmarshal(body, &googleResponse)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return &googleResponse, nil
}
