package actor

import "log"

func Supervise(task func() error) error {
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

func do(task func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	return task()
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
