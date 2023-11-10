package client

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// var client = lark.NewClient(config.Bot.Lark.VerificationToken, config.Bot.Lark.EncryptKey)

func GetDefaultHeader() *http.Header {
	h := &http.Header{}
	h.Add("Content-Type", "application/json; charset=utf-8")
	return h
}

func GetTenentAccessTokenHeader() (*http.Header, error) {
	token, err := GetTenentAccessToken()
	if err != nil {
		return nil, err
	}
	h := GetDefaultHeader()
	h.Add("Authorization", "Bearer "+token)
	return h, nil
}

func doHTTP(reqBody []byte, method, url string, header *http.Header) (respBody []byte, err error) {

	httpReq, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header = *header

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	respBody, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if len(respBody) == 0 {
		return nil, errors.New("doHTTP len(respBody)=0")
	}

	return
}
