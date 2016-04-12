package actor

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestSupervise(t *testing.T) {
	Supervise(func(session *Session) {
		time.Sleep(1 * time.Second)
		defer log.Println("Cleaned up")
		panic(errors.New("Omg Panic"))
	})
}
