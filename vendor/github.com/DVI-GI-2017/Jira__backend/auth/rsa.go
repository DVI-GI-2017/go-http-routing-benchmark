package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

const (
	keysDir = "rsa"

	privKey = "app.rsa"
	pubKey  = "app.rsa.pub"

	privKeyPath = keysDir + "/" + privKey
	pubKeyPath  = keysDir + "/" + pubKey
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	if !isExists(keysDir) {
		err := generateKeys()
		if err != nil {
			log.Panicf("can not generate keys: %v", err)
		}
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Panicf("can not read sign bytes from file: %v", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Panicf("can not parse rsa private key: %v", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Panicf("can not read verify bytes from file: %v", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Panicf("can not parse verify key: %v", err)
	}
}

// Generates rsa keys.
func generateKeys() error {
	err := os.Mkdir(keysDir, 0777)
	if err != nil {
		return fmt.Errorf("can not create directory for keys: %v", err)
	}

	err = generatePrivateKey(privKeyPath)
	if err != nil {
		return fmt.Errorf("can not generate private key: %v", err)
	}

	err = generatePublicKey(privKeyPath, pubKeyPath)
	if err != nil {
		return fmt.Errorf("can not generate public key: %v", err)
	}
	return nil
}

// Check if dir, file or symlink exists
func isExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
