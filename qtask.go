package qtask

// QTask Qtask is queue task
type QTask struct {
	Limit       int
	RuningTasks queue
	Queue       queue
	Tasks       []*Task
}

// New create new QTask
func New(limit int) (*QTask, error) {
	qtask := QTask{
		Limit: limit,
	}
	return &qtask, nil
}

// AddTask add task
func (qtask *QTask) AddTask(task Task) error {
	qtask.RuningTasks.Push(task)
	return nil
}

// Run run QTask
func (qtask *QTask) Run() {
	for {
		if qtask.Limit > 0 && qtask.RuningTasks.Count < qtask.Limit {
			queueData := qtask.Queue.Pop()
			task, _ := queueData.(Task)
			qtask.RuningTasks.Push(task)
			task.Process()
		}
	}
}

// Task task interface
type Task interface {
	Process() error
}
