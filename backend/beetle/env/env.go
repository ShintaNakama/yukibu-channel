package env

import "os"

type ENV struct {
	ENV string
}

func NewENV() *ENV {
	//env := os.Getenv("APP_ENV")
	return &ENV{
		ENV: os.Getenv("APP_ENV"),
	}
}

func IsProd(e *ENV) bool {
	return e.ENV == "prod"
}
