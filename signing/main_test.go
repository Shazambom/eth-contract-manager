package signing

import (
	"log"
	"os"
	"sync"
	"testing"
)

var privateKey, address string
var s *SignatureHandler

var portCounter = 9080
var mu sync.Mutex

func getTestPort() int {
	mu.Lock()
	defer mu.Unlock()
	portCounter++
	return portCounter
}
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
