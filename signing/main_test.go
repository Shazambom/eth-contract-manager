package signing

import (
	"log"
	"os"
	"testing"
)

var privateKey, address string
var s *SignatureHandler

func TestMain(m *testing.M) {
	s = &SignatureHandler{}
	var err error
	privateKey, address, err = s.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	code := m.Run()

	os.Exit(code)
}
