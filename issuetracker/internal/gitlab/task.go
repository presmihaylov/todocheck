package gitlab

import (
	"encoding/json"
	"time"

	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // Remove quotes

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *Date) AsTime() time.Time {
	return time.Time(*d)
}

// Task model for gitlab tasks
type Task struct {
	State   string `json:"state"`
	DueDate *Date  `json:"due_date,omitempty"`
}

// GetStatus of gitlab task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.State {
	case "closed":
		return taskstatus.Closed, nil
	default:
		if t.DueDate != nil && t.DueDate.AsTime().Before(time.Now()) {
			return taskstatus.Overdue, nil
		}
		return taskstatus.Open, nil
	}
}
