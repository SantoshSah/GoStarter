package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Main include info about API settings.
type Main struct {
	Port string `json:"port"`

	Auth        Auth      `json:""`
	DB          DB        `json:"database"`
	FE          AppConfig `json:"frontEnd"`
	DemoBalance float64   `json:"demoBalance"`
	Debug       Debug     `json:"debug"`
}

// Auth includes all authorization settings
type Auth struct {
	AccessToken  string `json:"accessToken"`         // CMS token
	AuthTokenTTL int    `json:"authTokenTtlSeconds"` // Time to life of user authentication link, seconds
	JWTSecret    string `json:"jwtSecret"`           // Secret key for generating user token
	JwtTTL       int    `json:"jwtTokenTtlMinutes"`  // Time to life of user session token
}

// DB include info about database connection.
type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

// Debug include all debugging flags
type Debug struct {
	Log bool `json:"log"`
	DB  bool `json:"db"`
}

// AppConfig includes info, required for front-end
type AppConfig struct {
	APITimeout int `json:"apiTimeout"` // Timeout for front-end requests

	// FIXME(tiabc): Abbreviations don't make the code more readable.
	LRP  string `json:"logoutRedirectPath"`
	DRP  string `json:"depositRedirectPath"`
	FLRP string `json:"firstLinkRedirectPath"`
	SLRP string `json:"secondLinkRedirectPath"`
	FLN  string `json:"firstLinkName"`
	SLN  string `json:"secondLinkName"`
	IJRP string `json:"invalidJwtRedirectPath"`
}

// FromFile read config from path and create configuration struct.
func FromFile(path string) Main {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic("error reading config " + path + ": " + err.Error())
	}
	var conf Main
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic("error parsing config " + path + ": " + err.Error())
	}

	return conf
}

//Save saves config to path
func Save(path string, config Main) error {
	configJSON, _ := json.MarshalIndent(config, "", "  ")
	err := ioutil.WriteFile(path, configJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

// DbConnURL create url to connect with database.
func (c *Main) DbConnURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name)
}
