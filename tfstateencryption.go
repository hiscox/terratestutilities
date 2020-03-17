package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// TfStateEncrypt calls the terrahelp binary for encrypting a statefile with AES
func TfStateEncrypt() {
	es := os.Getenv("ENCRYPTION_SECRET")
	if len(es) == 0 {
		log.Fatal("Env var ENCRYPTION_SECRET has not been defined")
	}

	c := "/bin/bash"
	args := []string{"terrahelp", "encrypt",
		"-simple-key=", es,
		"-file=", "../terraform.tfstate"}

	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Unable encrypt state file")
		log.Fatal(err)
	}
	// ! write some tests for this stuff!!
}

// TfStateDecrypt calls the terrahelp binary for decrypting a statefile with AES
func TfStateDecrypt() {
	es := os.Getenv("ENCRYPTION_SECRET")
	if len(es) == 0 {
		log.Fatal("Env var ENCRYPTION_SECRET has not been defined")
	}

	content, err := ioutil.ReadFile("../terraform.tfstate")
	if !strings.Contains(string(content), "@terrahelp-enc") {
		log.Printf("Your statefile is already unencrypted, it will be encrypted at the end")
		return
	}

	c := "/bin/bash"
	args := []string{"terrahelp", "decrypt",
		"-simple-key=", es,
		"-file=", "../terraform.tfstate"}

	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Printf("Unable to decrypt state file")
		log.Fatal(err)
	}
	return
	// ! write some tests for this stuff!!
}
