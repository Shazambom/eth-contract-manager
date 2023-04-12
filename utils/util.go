package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnvVars() (rdsEndpoint, rdsPwd, secretKey, gURL, siteKey, projectID, contractAddress, signingKey string, abi []string, maxCount, maxIncr int, validOrigins []string, envErr error) {
	if enabled := os.Getenv("ENABLED"); enabled != "true" {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, errors.New("Lambda not Enabled")
	}
	if abiErr := json.Unmarshal([]byte(os.Getenv("ABI")), &abi); abiErr != nil {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, abiErr
	}

	maxCount, envErr = strconv.Atoi(os.Getenv("MAX_COUNT"))
	if envErr != nil {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, envErr
	}
	maxIncr, envErr = strconv.Atoi(os.Getenv("MAX_INCR"))

	return os.Getenv("REDIS_ENDPOINT"),
		os.Getenv("REDIS_PWD"),
		os.Getenv("SECRET_KEY"),
		os.Getenv("GOOGLE_URL"),
		os.Getenv("SITE_KEY"),
		os.Getenv("PROJECT_ID"),
		os.Getenv("CONTRACT_ADDRESS"),
		os.Getenv("SIGNER_PRIVATE_KEY"),
		abi, maxCount, maxIncr,
		strings.Split(os.Getenv("VALID_ORIGINS"), "~"),
		envErr
}

func GetEnvVar(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", errors.New("environment variable doesn't exist")
	}
	return val, nil
}

func GetEnvVarBatch(keys []string, vars ...*string) error {
	for i, key := range keys {
		val, keyErr := GetEnvVar(key)
		if keyErr != nil {
			log.Println("Key: " + key)
			return keyErr
		}
		*vars[i] = val

	}
	return nil
}

func StrInStrList(str string, strList []string) bool {
	for _, val := range strList {
		if val == str {
			return true
		}
	}
	return false
}

func GetNonceBytes() ([]byte, error) {
	UUID := uuid.New()
	return UUID.MarshalBinary()
}

func GetNonce() (string, error) {
	UUID := uuid.New()
	UUIDBytes, uuidErr := UUID.MarshalBinary()
	if uuidErr != nil {
		return "", uuidErr
	}
	return "0x" + hex.EncodeToString(UUIDBytes), nil
}

func MergeChannels(cs ...<-chan string) <-chan string {
	//TODO debug this, I think it causes some hanging to occur
	out := make(chan string)
	for _, c := range cs {
		go func(c <-chan string) {
			out <- <-c
		}(c)
	}
	return out
}
