package supervisor

import (
	"sync"

	"github.com/ironbay/delta/uuid"
)

type Process struct {
	key  string
	task func(*Process)
	once sync.Once
	kill chan error
}

func NewProcess(task func(*Process)) *Process {
	return &Process{
		key:  uuid.Ascending(),
		task: task,
		kill: make(chan error),
	}
}

func (this *Process) Kill(err error) {
	this.once.Do(func() {
		this.kill <- err
	})
}

func (this *Process) Run() {
}
