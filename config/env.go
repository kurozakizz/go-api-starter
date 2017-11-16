package config

import "os"

var (
	apiVersion = os.Getenv("API_VERION")
)

func GetAPIPrefix() string {
	return "/v" + apiVersion
}
