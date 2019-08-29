package qtask

import (
	"fmt"
)

// QTask Qtask is queue task
type QTask struct {
	taskChan        chan Task
	runingCountChan chan int
	Limit           int
	RunningCount    int
	Queue           queue
	Tasks           []*Task
}

// Create create new QTask
func Create(limit int) (*QTask, error) {
	qtask := QTask{
		taskChan:        make(chan Task),
		runingCountChan: make(chan int),
		Limit:           limit,
		Queue:           queue{0, nil, nil},
	}
	return &qtask, nil
}

// Add add task
func (qtask *QTask) Add(task Task) {
	qtask.taskChan <- task
}

func (qtask *QTask) handleTaskChan() {
	for {
		task := <-qtask.taskChan
		qtask.Queue.Push(task)
		fmt.Println("add task", task)
	}
}

func (qtask *QTask) handleRunningCountChan() {
	for {
		inc := <-qtask.runingCountChan
		qtask.RunningCount += inc
		fmt.Println(qtask.RunningCount, inc)
	}
}

func (qtask *QTask) handleTask() {
	for {
		if qtask.Queue.Count > 0 && qtask.RunningCount < qtask.Limit {
			qtask.runingCountChan <- 1
			task := qtask.Queue.Pop()
			go func() {
				task.Process()
				qtask.runingCountChan <- -1
			}()
		}
	}
}

// Run run QTask
func (qtask *QTask) Run() {
	fmt.Printf("QTask start run\nTask concurrent limit %d\n", qtask.Limit)
	go qtask.handleRunningCountChan()
	go qtask.handleTaskChan()
	go qtask.handleTask()
}
