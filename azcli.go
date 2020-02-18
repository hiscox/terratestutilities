package util

import (
	"errors"
	"os"
)

// AzCliAuth ensures environment variables are set for Azure CLI authentication
func AzCliAuth() (cSecret string, cID string, tenID string, subID string, err error) {
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return "", "", "", "", errors.New("Env var ARM_CLIENT_SECRET not set")
	}
	clientID := os.Getenv("ARM_CLIENT_ID")
	if len(clientID) == 0 {
		return "", "", "", "", errors.New("Env var ARM_CLIENT_ID not set")
	}
	tenantID := os.Getenv("ARM_TENANT_ID")
	if len(tenantID) == 0 {
		return "", "", "", "", errors.New("Env var ARM_TENANT_ID not set")
	}
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		return "", "", "", "", errors.New("Env var ARM_SUBSCRIPTION_ID not set")
	}
	return clientSecret, clientID, tenantID, subscriptionID, nil
}
