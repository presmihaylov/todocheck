package gitlab

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

func Test_Task_GetStatus(t *testing.T) {
	t.Parallel()

	overdueDate := Date(time.Now().Add(-24 * time.Hour))

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
		testname := tt.task.State
		t.Run(testname, func(t *testing.T) {
			status, _ := tt.task.GetStatus()
			if status != tt.want {
				t.Errorf("got %d, want %d", status, tt.want)
			}
		})
	}

}

func Test_Task_Marshalling(t *testing.T) {
	t.Parallel()

	overDueTestDate := Date(time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC))

	futureDate := Date(time.Now().Add(2 * 24 * time.Hour))

	var tests = []struct {
		name string
		task Task
		want string
	}{
		{
			name: "closed_task",
			task: Task{State: "closed"},
			want: `{"state":"closed"}`,
		},
		{
			name: "open_task",
			task: Task{State: "open"},
			want: `{"state":"open"}`,
		},
		{
			name: "reopened_task",
			task: Task{State: "reopened"},
			want: `{"state":"reopened"}`,
		},
		{
			name: "task_with_due_date_over_due",
			task: Task{State: "open", DueDate: &overDueTestDate},
			want: `{"state":"open","due_date":"2023-12-25"}`,
		},
		{
			name: "task_with_due_date_not_over_due",
			task: Task{State: "open", DueDate: &futureDate},
			want: fmt.Sprintf(`{"state":"open","due_date":"%s"}`, futureDate.AsTime().Format("2006-01-02")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.task)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}
			if string(got) != tt.want {
				t.Errorf("json.Marshal() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
