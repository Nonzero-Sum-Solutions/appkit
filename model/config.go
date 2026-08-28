package model

import "github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit"

type configDefault struct {
	url, apiKey, userEmail, userPwd string
	dbTimeoutSeconds                int
}

func (c *configDefault) SupabaseTimeoutSeconds() int {
	return c.dbTimeoutSeconds
}

func (c *configDefault) SupabaseUserEmail() string {
	return c.userEmail
}

func (c *configDefault) SupabaseUserPwd() string {
	return c.userPwd
}

func (c *configDefault) SupabaseAPIKey() string {
	return c.apiKey
}

func (c *configDefault) SupabaseURL() string {
	return c.url
}

func NewDBConfig(url, apiKey, userEmail, userPwd string, dbTimeoutSeconds int) appkit.DBConfig {
	return &configDefault{url, apiKey, userEmail, userPwd, dbTimeoutSeconds}
}
