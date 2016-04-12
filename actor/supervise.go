package actor

import (
	"log"
	"sync"
	"time"
)

type Session struct {
	Stop func(err error)

	cleanup []func()
}

func (this *Session) Cleanup(cb func()) {
	this.cleanup = append(this.cleanup, cb)
}

func (this *Session) clean() {
	for _, cb := range this.cleanup {
		cb()
	}
}

func Supervise(task func(*Session), strategies ...func()) error {
	for {
		err := do(task)
		if err == nil {
			return nil
		}
		if ae, ok := err.(*ActorError); ok {
			return ae
		}
		log.Println("Restarting:", err)
		for _, s := range strategies {
			s()
		}
	}
}

func do(task func(*Session)) (err error) {
	stop := make(chan error)
	once := sync.Once{}
	session := &Session{
		Stop: func(err error) {
			once.Do(func() {
				stop <- err
			})
		},
	}
	defer session.clean()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
				session.Stop(err)
			}
		}()
		task(session)
	}()
	return <-stop
}

type ActorError struct {
	message string
}

func (*ActorError) Error() string {
	return ""
}

func Error(message string) *ActorError {
	return &ActorError{
		message,
	}
}

func SleepStrategy(seconds int) func() {
	return func() {
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
