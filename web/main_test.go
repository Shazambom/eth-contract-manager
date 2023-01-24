package web

import (
	"os"
	"sync"
	"testing"
)

var portCounter = 11080
var mu sync.Mutex

func getTestPort() int {
	mu.Lock()
	defer mu.Unlock()
	portCounter++
	return portCounter
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
