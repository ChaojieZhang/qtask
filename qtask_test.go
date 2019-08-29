package qtask

import (
	"fmt"
	"testing"
	"time"
)

type TestTask struct {
	Number int
}

func (task *TestTask) Process() error {
	fmt.Printf("task: %d start\n", task.Number)
	time.Sleep(1 * time.Second)
	fmt.Printf("task: %d end\n", task.Number)
	return nil
}

func TestRun(t *testing.T) {
	qt, err := Create(2)
	if err != nil {
		t.Error(err)
	}
	qt.Run()
	for i := 0; i <= 10; i++ {
		qt.Add(&TestTask{
			Number: i,
		})
	}
	time.Sleep(2 * 10 * time.Second)
}
