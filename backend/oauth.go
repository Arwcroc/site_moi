package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OauthToken struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	CreatedAt        int    `json:"created_at"`
	SecretValidUntil int    `json:"secret_valid_until"`
	ExpiresDate      time.Time
}

type OauthClient struct {
	Token	OauthToken
	Url		string
	Uid		string
	Secret	string
}

func (c *OauthClient) CheckToken() bool {
	if c.Token.AccessToken == "" {
		return false
	}
	if c.Token.ExpiresDate.Before(time.Now()) {
		return false
	}
	return true
}

func (c *OauthClient) GrabToken() error {
	requestUrl := fmt.Sprintf(
		"%s?grant_type=client_credentials&client_id=%s&client_secret=%s",
		c.Url,
		c.Uid,
		c.Secret,
	)

	resp, err := http.Post(requestUrl, "", bytes.NewBuffer([]byte("")))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &c.Token)
	if err != nil {
		return err
	}
	c.Token.ExpiresDate = time.Now().Add(time.Second * time.Duration(c.Token.ExpiresIn))
	return nil
}

func (c *OauthClient) RefreshToken() error {
	if c.CheckToken() {
		return nil
	}
	return c.GrabToken()
}
