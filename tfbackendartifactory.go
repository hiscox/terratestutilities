package util

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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

// TfBackendArtifactoryGet downloads a file from Artifactory
func TfBackendArtifactoryGet(url string, user string, pwd string, repo string, subpath string) (exists bool) {
	// transport controls settings such as proxy and timeouts
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    15 * time.Second,
		DisableCompression: true, // allows us to read the raw body
	}
	// client controls redirects, headers, etc
	// ? might need a redirect func to add headers when a redirect occurs
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url+"/"+repo+"/"+subpath+"/terraform.tfstate", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	switch resp.StatusCode {
	case 200:
		log.Printf("State file found at " + url + "/" + repo + "/" + subpath)
	case 404:
		log.Printf("WARNING: 404. Either no state file exists or you repo/subpath are incorrect!")
		log.Printf("Carrying on as if this is the first time terraform has been run")
	case 403:
		log.Fatal("Forbidden, check your proxy/networking?")
	case 401:
		log.Fatal("Unauthorised, please check your credentials and/or endpoint")
	default:
		log.Printf(string(resp.StatusCode))
		log.Fatal("Unhandled failure")
	}
	defer resp.Body.Close() // even if later methods return err, run Close() at the end
	out, err := os.Create("../terraform.tfstate")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	return true
	// ! write some tests for this stuff!!
}

// TfBackendArtifactoryPost publishes a file to Artifactory
func TfBackendArtifactoryPost(url string, user string, pwd string, repo string, subpath string) {
	// transport controls settings such as proxy and timeouts
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    15 * time.Second,
		DisableCompression: true, // allows us to read the raw body
	}
	// client controls redirects, headers, etc
	// ? might need a redirect func to add headers when a redirect occurs
	client := &http.Client{
		Transport: tr,
	}
	data, err := os.Open("../terraform.tfstate")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	req, err := http.NewRequest("POST", url+"/"+repo+"/"+subpath+"/terraform.tfstate", data)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, pwd)
	//req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	switch resp.StatusCode {
	case 200:
		log.Printf("State replaced at " + url + "/" + repo + "/" + subpath)
	case 201:
		log.Printf("State created at " + url + "/" + repo + "/" + subpath)
	case 403:
		log.Fatal("Forbidden, check your proxy/networking?")
	case 401:
		log.Fatal("Unauthorised, please check your credentials and/or endpoint")
	default:
		log.Printf(string(resp.StatusCode))
		log.Fatal("Unhandled failure")
	}
	// ! write some tests for this stuff!!
}
