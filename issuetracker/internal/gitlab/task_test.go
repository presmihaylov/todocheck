package gitlab

import (
	"testing"
	"time"

	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

func Test_Task_GetStatus(t *testing.T) {

	t.Parallel()

	overdueDate := time.Now().Add(-24 * time.Hour) 

	var tests = []struct {
		task Task
		want taskstatus.TaskStatus
	}{
		{Task{State: "closed"}, taskstatus.Closed},
		{Task{State: "open"}, taskstatus.Open},
		{Task{State: "reopened"}, taskstatus.Open},
		{Task{State: "open", DueDate: &overdueDate}, taskstatus.Overdue},
	}

	for _, tt := range tests {
		testname :=  tt.task.State
		t.Run(testname, func(t *testing.T) {
			status, _ := tt.task.GetStatus() 
			if status != tt.want {
				t.Errorf("got %d, want %d", status, tt.want)
			}

		})
	}

}
