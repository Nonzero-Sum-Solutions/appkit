package config

import (
	"fmt"
	"os"

	"github.com/Nonzero-Sum-Solutions/appkit/model"
	"github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit"
	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

func LoadConfig() (appkit.DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	url := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")
	userEmail := os.Getenv("SUPABASE_USER_EMAIL")
	userPwd := os.Getenv("SUPABASE_USER_PWD")

	return model.NewDBConfig(url, apiKey, userEmail, userPwd, 0), nil
}

func LoadConfigWithViper() (appkit.DBConfig, error) {
	url := viper.GetString("SUPABASE_URL")
	apiKey := viper.GetString("SUPABASE_API_KEY")
	userEmail := viper.GetString("SUPABASE_USER_EMAIL")
	userPwd := viper.GetString("SUPABASE_USER_PWD")

	if url == "" || apiKey == "" || userEmail == "" || userPwd == "" {
		return nil, fmt.Errorf("missing Supabase config: set SUPABASE_URL, SUPABASE_API_KEY, SUPABASE_USER_EMAIL, and SUPABASE_USER_PWD")
	}

	return model.NewDBConfig(url, apiKey, userEmail, userPwd, 0), nil
}
