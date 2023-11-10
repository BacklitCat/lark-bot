package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"larkbot/internal/config"
	"larkbot/larktype"
	"net/http"
	"time"
)

var (
	TenantAccessToken           string
	TenantAccessTokenExpire     int64
	TenantAccessTokenExpireTime int64
)

// GetTenentAccessToken
// 获取、更新TenentAccessToken
func GetTenentAccessToken() (string, error) {
	// 如果没有token 或者 token 过期， 需要更新token
	time_now := time.Now().Unix()
	if TenantAccessToken == "" || time_now >= TenantAccessTokenExpireTime-config.Bot.Lark.UpdateTokenDeltaTimeSecond {
		token, expire, err := GetTenentAccessTokenFromLark()
		if err != nil {
			return "", err
		}

		// 理论上token, expire必不同，更新
		TenantAccessToken = token
		TenantAccessTokenExpire = expire
		TenantAccessTokenExpireTime = time_now + expire
		// return TenantAccessToken, nil
	}
	// 不需要更新 token，取内存缓存
	return TenantAccessToken, nil
}

// GetTenentAccessTokenFromLark
// 从飞书更新TenentAccessToken
func GetTenentAccessTokenFromLark() (string, int64, error) {

	taReq := &larktype.TenantAccessTokenReq{
		AppID:     config.Bot.Lark.AppID,
		AppSecret: config.Bot.Lark.Secret.AppSecret,
	}

	taResp := &larktype.TenantAccessTokenResp{}

	playload, err := json.Marshal(&taReq)
	if err != nil {
		return "", 0, err
	}
	if len(playload) == 0 {
		return "", 0, errors.New("marshal restult is empty")
	}

	httpReq, err := http.NewRequest(
		config.LarkAPI.TENANT_ACCESS_TOKEN.Method,
		config.LarkAPI.TENANT_ACCESS_TOKEN.Url,
		bytes.NewReader(playload))
	if err != nil {
		return "", 0, err
	}
	httpReq.Header.Add("Content-Type", "application/json; charset=utf-8")

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", 0, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", 0, err
	}

	if err = json.Unmarshal(body, taResp); err != nil {
		return "", 0, err
	}

	if taResp.Code != 0 {
		return "", 0, fmt.Errorf("code=%d, msg=%s", taResp.Code, taResp.Msg)
	}

	return taResp.TenantAccessToken, taResp.Expire, nil
}
