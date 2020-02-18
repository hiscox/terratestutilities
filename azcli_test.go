package util

import (
	"os"
	"testing"
)

func TestAzCliAuth(t *testing.T) {
	os.Setenv("ARM_CLIENT_SECRET", "a-guid-here-0")
	os.Setenv("ARM_CLIENT_ID", "a-guid-here-1")
	os.Setenv("ARM_TENANT_ID", "a-guid-here-2")
	os.Setenv("ARM_SUBSCRIPTION_ID", "a-guid-here-3")

	t.Run("all environment variables are set", func(t *testing.T) {
		_, _, _, _, err := AzCliAuth()
		if err != nil {
			t.Errorf("got %v, expected %v", err, "nil")
		}
	})

	t.Run("ARM_CLIENT_SECRET has not been set", func(t *testing.T) {
		os.Unsetenv("ARM_CLIENT_SECRET")
		_, _, _, _, err := AzCliAuth()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var ARM_CLIENT_SECRET not set")
		}
	})
	t.Run("ARM_CLIENT_ID has not been set", func(t *testing.T) {
		os.Setenv("ARM_CLIENT_SECRET", "a-guid-here-0")
		os.Unsetenv("ARM_CLIENT_ID")
		_, _, _, _, err := AzCliAuth()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var ARM_CLIENT_ID not set")
		}
	})
	t.Run("ARM_TENANT_ID has not been set", func(t *testing.T) {
		os.Setenv("ARM_CLIENT_ID", "a-guid-here-1")
		os.Unsetenv("ARM_TENANT_ID")
		_, _, _, _, err := AzCliAuth()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var ARM_TENANT_ID not set")
		}
	})
	t.Run("ARM_SUBSCRIPTION_ID has not been set", func(t *testing.T) {
		os.Setenv("ARM_TENANT_ID", "a-guid-here-2")
		os.Unsetenv("ARM_SUBSCRIPTION_ID")
		_, _, _, _, err := AzCliAuth()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "ARM_SUBSCRIPTION_ID not set")
		}
	})
}
