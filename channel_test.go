package growler

import (
	"sync"
	"testing"
)

var mutex = sync.Mutex{}
var ch = make(chan bool, 1)

func UseMutex() {
	mutex.Lock()
	mutex.Unlock()
}
func UseChan() {
	ch <- true
	<-ch
}

func BenchmarkUseMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseMutex()
	}
}
func BenchmarkUseChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseChan()
	}
}
