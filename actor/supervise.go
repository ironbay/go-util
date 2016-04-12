package actor

import (
	"log"
	"sync"
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

func Supervise(task func(*Session)) error {
	for {
		err := do(task)
		if err == nil {
			return nil
		}
		if ae, ok := err.(*ActorError); ok {
			return ae
		}
		log.Println("Restarting")
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
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	go task(session)
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
