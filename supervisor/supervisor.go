package supervisor

type Supervisor struct {
	processes map[string]*Process
}

func New() *Supervisor {
	return &Supervisor{
		processes: make(map[string]*Process),
	}
}

func (this *Supervisor) Spawn(task func(*Process)) *Process {
	process := NewProcess(task)
	this.processes[process.key] = process
	return process
}
