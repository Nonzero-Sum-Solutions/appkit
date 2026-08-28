package db

import (
	"net/http"
	"time"

	"github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit"
	"github.com/google/uuid"
	supabase "github.com/supabase-community/supabase-go"
)

func NewAuthenticatedSupabaseClient(c appkit.DBConfig) (*supabase.Client, uuid.UUID, error) {
	options := supabase.ClientOptions{}
	client, err := supabase.NewClient(c.SupabaseURL(), c.SupabaseAPIKey(), &options)
	if err != nil {
		return nil, uuid.Nil, err
	}

	if c.SupabaseTimeoutSeconds() > 0 {
		customHTTPClient := &http.Client{
			Timeout: time.Duration(c.SupabaseTimeoutSeconds()) * time.Second,
		}

		client.Auth = client.Auth.WithClient(*customHTTPClient)
	}

	session, err := client.SignInWithEmailPassword(c.SupabaseUserEmail(), c.SupabaseUserPwd())
	if err != nil {
		return nil, uuid.Nil, err
	}
	client.UpdateAuthSession(session)
	client.EnableTokenAutoRefresh(session)

	return client, session.User.ID, nil
}
