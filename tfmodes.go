package util

import (
	"errors"
	"log"
	"os"
	"strings"
)

// TfModes validates how terraform will execute and sets sensible defaults.
func TfModes() (d string, p string, a string, err error) {
	destroy := os.Getenv("TF_DESTROY")
	if len(destroy) == 0 {
		log.Printf("Env var TF_DESTROY not set, setting TF_DESTROY to false")
		destroy = "false"
	}
	destroy = strings.ToLower(destroy)

	plan := os.Getenv("TF_PLAN")
	if len(plan) == 0 {
		log.Printf("Env var TF_PLAN not set, setting TF_PLAN to true")
		plan = "true"
	}
	plan = strings.ToLower(plan)

	apply := os.Getenv("TF_APPLY")
	if len(apply) == 0 {
		log.Printf("Env var TF_APPLY not set, setting TF_APPLY to false")
		apply = "false"
	}
	apply = strings.ToLower(apply)

	if (apply == "false" || apply == "true") && (destroy == "false" || destroy == "true") && (plan == "false" || plan == "true") {
		return destroy, plan, apply, nil
	}
	return "", "", "", errors.New("TfModes environment variables are not syntactically correct. Only 'true' or 'false' will be accepted")
}
