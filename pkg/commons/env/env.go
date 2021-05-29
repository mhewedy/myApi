package env

import (
	"os"
)

func Get(key string, defValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defValue
	}
	return val
}
