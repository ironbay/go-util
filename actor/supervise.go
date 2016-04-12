package actor

import "log"

type Session struct {
	Stop chan error

	cleanup []func()
}

func NewSession() *Session {
	return &Session{
		Stop: make(chan error, 1),
	}
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
	session := NewSession()
	defer session.clean()
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	go task(session)
	return <-session.Stop
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
