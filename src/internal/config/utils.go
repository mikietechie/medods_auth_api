/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package config

import "os"

func GetEnvOrDef(key, def string) string {
	env := os.Getenv(key)
	if len(env) != 0 {
		return env
	}
	return def
}
