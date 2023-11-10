package client_test

import (
	"larkbot/internal/client"
	"testing"
)

func GetTenentAccessTokenFromLark(t *testing.T) {
	token, expire, err := client.GetTenentAccessTokenFromLark()
	if err != nil {
		t.Error(err)
	}
	t.Logf("token=%s, expire=%d\n", token, expire)
}

func TestGetTenentAccessToken(t *testing.T) {
	// 每隔2s取一次，一共取3次
	t.Logf("token=%s, expire=%d\n", client.TenantAccessToken, client.TenantAccessTokenExpireTime)
	for i := 0; i < 2; i++ {
		token, err := client.GetTenentAccessToken()
		if err != nil {
			t.Error(err)
		}
		t.Logf("token=%s, client.TenantAccessToken=%s, expireTime=%d\n", token, client.TenantAccessToken, client.TenantAccessTokenExpireTime)
	}

}
