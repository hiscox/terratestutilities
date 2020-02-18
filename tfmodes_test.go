package util

import (
	"os"
	"testing"
)

func TestTfModes(t *testing.T) {
	t.Run("default environment variables are used", func(t *testing.T) {
		_, _, _, err := TfModes()
		if err != nil {
			t.Errorf("got %v, expected %v", err, "nil")
		}
	})
	t.Run("invalid environment variable syntax", func(t *testing.T) {
		os.Setenv("TF_DESTROY", "foobar")
		_, _, _, err := TfModes()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "an error to be raised")
		}
	})
	t.Run("explicit environment variables are set", func(t *testing.T) {
		os.Setenv("TF_DESTROY", "false")
		os.Setenv("TF_PLAN", "false")
		os.Setenv("TF_APPLY", "true")
		_, _, _, err := TfModes()
		if err != nil {
			t.Errorf("got %v, expected %v", err, "nil")
		}
	})
}
