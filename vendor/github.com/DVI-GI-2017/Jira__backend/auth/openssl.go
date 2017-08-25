package auth

import "os/exec"

// Generates private key and saves it in the "privKey" file.
func generatePrivateKey(privKey string) error {
	return openssl("genrsa", "-out", privKey, "1024")
}

// Generates public key for private key from "privKey" and saves it in the "pubKey" file.
func generatePublicKey(privKey, pubKey string) error {
	return openssl("rsa", "-in", privKey, "-out", pubKey, "-pubout")
}

// Wraps openssl command
func openssl(args ...string) error {
	command := &exec.Cmd{
		Path: getOpenssl(),
		Args: args,
	}
	return command.Run()
}

// Returns path to openssl binary
func getOpenssl() string {
	return "/usr/bin/openssl"
}
