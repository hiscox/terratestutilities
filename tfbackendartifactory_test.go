package util

import (
	"os"
	"testing"
)

func TestTfBackendArtifactory(t *testing.T) {
	os.Setenv("TFBACKEND_URL", "https://foo.bar.io")
	os.Setenv("TFBACKEND_USER", "qwerty")
	os.Setenv("TFBACKEND_PASSWORD", "complex")
	os.Setenv("TFBACKEND_REPO", "repo-of-states")
	os.Setenv("TFBACKEND_SUBPATH", "nested-dir")

	t.Run("all environment variables are set", func(t *testing.T) {
		_, _, _, _, _, err := TfBackendArtifactory()
		if err != nil {
			t.Errorf("got %v, expected %v", err, "nil")
		}
	})

	t.Run("TFBACKEND_URL has not been set", func(t *testing.T) {
		os.Unsetenv("TFBACKEND_URL")
		_, _, _, _, _, err := TfBackendArtifactory()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var TFBACKEND_URL not set")
		}
	})
	t.Run("TFBACKEND_USER has not been set", func(t *testing.T) {
		os.Setenv("TFBACKEND_URL", "https://foo.bar.io")
		os.Unsetenv("TFBACKEND_USER")
		_, _, _, _, _, err := TfBackendArtifactory()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var TFBACKEND_USER not set")
		}
	})
	t.Run("TFBACKEND_PASSWORD has not been set", func(t *testing.T) {
		os.Setenv("TFBACKEND_USER", "qwerty")
		os.Unsetenv("TFBACKEND_PASSWORD")
		_, _, _, _, _, err := TfBackendArtifactory()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "Env var TFBACKEND_PASSWORD not set")
		}
	})
	t.Run("TFBACKEND_REPO has not been set", func(t *testing.T) {
		os.Setenv("TFBACKEND_PASSWORD", "complex")
		os.Unsetenv("TFBACKEND_REPO")
		_, _, _, _, _, err := TfBackendArtifactory()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "TFBACKEND_REPO not set")
		}
	})
	t.Run("TFBACKEND_SUBPATH has not been set", func(t *testing.T) {
		os.Setenv("TFBACKEND_REPO", "repo-of-states")
		os.Unsetenv("TFBACKEND_SUBPATH")
		_, _, _, _, _, err := TfBackendArtifactory()
		if err == nil {
			t.Errorf("got %v, expected %v", "nil", "TFBACKEND_SUBPATH not set")
		}
	})
}
