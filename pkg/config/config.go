package config

import (
	"encoding/json"
	"io"
	"os"
)

type Socials struct {
	Facebook  string `json:"facebook" validate:"required,url"`
	Twitter   string `json:"twitter" validate:"required,url"`
	Instagram string `json:"instagram" validate:"required,url"`
	Youtube   string `json:"youtube" validate:"required,url"`
}

type Colors struct {
	MainBackground          string `json:"mainBackground" validate:"iscolor"`
	DarkBackground          string `json:"darkBackground" validate:"iscolor"`
	PrimaryBackground       string `json:"primaryBackground" validate:"iscolor"`
	PrimaryBackgroundActive string `json:"primaryBackgroundActive" validate:"iscolor"`
	Secondary               string `json:"secondary" validate:"iscolor"`
	SecondaryLight          string `json:"secondaryLight" validate:"iscolor"`
	LoaderBackground        string `json:"loaderBackground" validate:"iscolor"`
}

type Images struct {
	HeaderLogo        string `json:"headerLogo" validate:"required"`
	FooterLogo        string `json:"footerLogo" validate:"required"`
	DefaultBackground string `json:"defaultBackground" validate:"required"`
	MainLogo          string `json:"mainLogo" validate:"required"`
	Favicon           string `json:"favicon" validate:"required"`
}

type Contact struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

type Didomi struct {
	ApiKey   string `json:"apiKey" validate:"required"`
	NoticeId string `json:"noticeId" validate:"required"`
}

type Firebase struct {
	Candidate  FirebaseConfig `json:"candidate" validate:"required"`
	Staging    FirebaseConfig `json:"staging" validate:"required"`
	Production FirebaseConfig `json:"production" validate:"required"`
}

type FirebaseConfig struct {
	ApiKey            string `json:"apiKey" validate:"required"`
	AuthDomain        string `json:"authDomain" validate:"required"`
	DatabaseURL       string `json:"databaseURL" validate:"required"`
	ProjectId         string `json:"projectId" validate:"required"`
	StorageBucket     string `json:"storageBucket" validate:"required"`
	MessagingSenderId string `json:"messagingSenderId" validate:"required"`
	AppId             string `json:"appId" validate:"required"`
	MeasurementId     string `json:"measurementId" validate:"required"`
}

type Config struct {
	AppName            string `json:"appName" validate:"required"`
	AppId              string
	ShortName          string   `json:"shortName" validate:"required"`
	DirName            string   `json:"dirName" validate:"required"`
	GtmId              string   `json:"gtmId" validate:"required"`
	BrandWebsiteUrl    string   `json:"brandWebsiteUrl" validate:"required,url"`
	MenuSource         string   `json:"menuSource" validate:"required"`
	BackgroundVideoUrl string   `json:"backgroundVideoUrl" validate:"required,url"`
	Images             Images   `json:"images" validate:"required"`
	Socials            Socials  `json:"socials" validate:"required"`
	Contact            Contact  `json:"contact" validate:"required"`
	Colors             Colors   `json:"colors" validate:"required"`
	Didomi             Didomi   `json:"didomi" validate:"required"`
	Firebase           Firebase `json:"firebase" validate:"required"`
}

func NewConfig(path string) *Config {
	configFile, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	byteConfig, err := io.ReadAll(configFile)

	if err != nil {
		panic(err)
	}

	var config Config
	json.Unmarshal(byteConfig, &config)

	return &config
}