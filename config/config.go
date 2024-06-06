package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Initialize Struct Db Config
type DbConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbDriver   string
}

// Initialize Struct Token Config
type TokenConfig struct {
	TokenIssue    string `json:"TokenIssue"`
	TokenSecret   []byte `json:"TokenSecret"`
	TokenExpire   time.Duration
	SigningMethod *jwt.SigningMethodHMAC
}

// Initialize Struct Api Config
type ApiConfig struct {
	ApiPort string
}

// Initialize Struct Smtp Config
type SmtpConfig struct {
	EmailName    string
	EmailAppPswd string
}

// Initialize Struct Client Config
type ClientConfig struct {
	ResetPasswordURL url.URL
}

// Initialize Struct Midtrans Config
type MidtransConfig struct {
	ClientKey string
	ServerKey string
}

// Initialize Struct Config
type Config struct {
	DbConfig
	TokenConfig
	ApiConfig
	SmtpConfig
	ClientConfig
	MidtransConfig
	Env string
}

// Configuration
func (c *Config) Configuration() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	c.DbConfig = DbConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbDriver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))

	c.TokenConfig = TokenConfig{
		TokenIssue:    os.Getenv("TOKEN_ISSUE"),
		TokenSecret:   []byte(os.Getenv("TOKEN_SECRET")),
		TokenExpire:   time.Hour * time.Duration(tokenExpire),
		SigningMethod: jwt.SigningMethodHS256,
	}

	c.SmtpConfig = SmtpConfig{
		EmailName:    os.Getenv("EMAIL_NAME"),
		EmailAppPswd: os.Getenv("EMAIL_APP_PASSWORD"),
	}

	c.MidtransConfig = MidtransConfig{
		ClientKey: os.Getenv("CLIENT_KEY"),
		ServerKey: os.Getenv("SERVER_KEY"),
	}

	c.Env = os.Getenv("ENV")
	if c.Env != "development" && c.Env != "staging" && c.Env != "production" {
		return fmt.Errorf("invalid env: %v", c.Env)
	}

	rpURL, err := url.Parse(os.Getenv("RESET_PASSWORD_URL"))
	if err != nil {
		return fmt.Errorf("error parsing reset password url: %v, env may be missing", err)
	}

	c.ClientConfig = ClientConfig{
		ResetPasswordURL: *rpURL,
	}

	if c.DbHost == "" || c.DbPort == "" || c.DbUser == "" || c.DbPassword == "" || c.DbName == "" || c.DbDriver == "" || c.ApiPort == "" || c.TokenIssue == "" || len(c.TokenSecret) == 0 || c.TokenExpire < 0 || c.SigningMethod == nil || c.Env == "" || c.EmailName == "" || c.EmailAppPswd == "" {
		return fmt.Errorf("missing environment variables")
	}

	return nil
}

// Construction to Access Config
func NewConfig() (*Config, error) {
	config := &Config{}

	if err := config.Configuration(); err != nil {
		return nil, err
	}

	return config, nil
}
