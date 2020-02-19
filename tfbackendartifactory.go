package util

import (
	"errors"
	"os"
)

// TfBackendArtifactory ensures environment variables have been set for using Artifactory
// as a remote terraform backend
func TfBackendArtifactory() (url string, user string, pwd string, repo string, subpath string, err error) {
	url = os.Getenv("TFBACKEND_URL")
	if len(url) == 0 {
		return "", "", "", "", "", errors.New("Env var TFBACKEND_URL not set")
	}
	u := os.Getenv("TFBACKEND_USER")
	if len(u) == 0 {
		return "", "", "", "", "", errors.New("Env var TFBACKEND_USER not set")
	}
	p := os.Getenv("TFBACKEND_PASSWORD")
	if len(p) == 0 {
		return "", "", "", "", "", errors.New("Env var TFBACKEND_PASSWORD not set")
	}
	repo = os.Getenv("TFBACKEND_REPO")
	if len(repo) == 0 {
		return "", "", "", "", "", errors.New("Env var TFBACKEND_REPO not set")
	}
	subpath = os.Getenv("TFBACKEND_SUBPATH")
	if len(subpath) == 0 {
		return "", "", "", "", "", errors.New("Env var TFBACKEND_SUBPATH not set")
	}
	return url, u, p, repo, subpath, nil
}
