package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-android/sdk"
)

// SignatureType ..
type SignatureType string

// SignatureType values
const (
	KeystoreSignatureType SignatureType = "keystore"
)

// KeystoreSignatureConfiguration ..
type KeystoreSignatureConfiguration struct {
	keystorePth          string
	keystorePassword     string
	aliasPassword        string
	alias                string
	nextKeystorePth      string
	nextKeystorePassword string
	nextAliasPassword    string
	nextAlias            string
	lineagePth           string
}

// SignatureConfiguration ...
type SignatureConfiguration struct {
	apkSigner             string
	signerScheme          string
	debuggablePermitted   string
	signatureType         SignatureType
	keystoreConfiguration *KeystoreSignatureConfiguration
}

func buildAPKSignerPath() (string, error) {
	androidHome := os.Getenv("ANDROID_HOME")
	androidSDK, err := sdk.New(androidHome)
	if err != nil {
		return "", fmt.Errorf("could not locate Android SDK using environment variable ANDROID_HOME='%s': %s", androidHome, err)
	}

	signer, err := androidSDK.LatestBuildToolPath("apksigner")

	if err != nil {
		return "", fmt.Errorf("failed to create sdk model: %s", err)
	}

	return signer, err
}

// NewKeystoreSignatureConfiguration ...
func NewKeystoreSignatureConfiguration(
	keystore string, keystorePassword string, alias string,
	nextKeystore string, nextKeystorePassword string, nextAlias string,
	aliasPassword string, debuggablePermitted string, signerScheme string) (SignatureConfiguration, error) {
	apkSigner, err := buildAPKSignerPath()

	if err != nil {
		return SignatureConfiguration{}, err
	}

	keystoreConfig := KeystoreSignatureConfiguration{
		keystorePth:          keystore,
		keystorePassword:     keystorePassword,
		alias:                alias,
		nextKeystorePth:      nextKeystore,
		nextKeystorePassword: nextKeystorePassword,
		nextAlias:            nextAlias,
		aliasPassword:        aliasPassword,
	}

	return SignatureConfiguration{
		apkSigner:             apkSigner,
		debuggablePermitted:   debuggablePermitted,
		signerScheme:          signerScheme,
		signatureType:         KeystoreSignatureType,
		keystoreConfiguration: &keystoreConfig,
	}, nil
}
