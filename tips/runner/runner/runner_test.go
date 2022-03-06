package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	a := assert.New(t)

	{
		runner := New(&Config{
			Timeout:  2,
			Interval: 4,
			Cmd:      "sleep 10",
		})
		result, err := runner.Run()
		expected := Result{
			Cmd:    "sleep 10",
			Output: "command timeout: ",
			Status: StatusTimeout,
		}
		a.Equal(expected, *result)
		if err == nil || !a.Contains(err.Error(), "killed") {
			t.Fatalf("UnexpectedErr")
		}
	}

	{
		runner := New(&Config{
			Timeout:  2,
			Interval: 4,
			Cmd:      "sh -c 'sleep 10'",
			UseShell: true,
		})
		result, err := runner.Run()
		expected := Result{
			Cmd:    "sh -c 'sleep 10'",
			Output: "command timeout: ",
			Status: StatusTimeout,
		}
		a.Equal(expected, *result)
		if err == nil || !a.Contains(err.Error(), "killed") {
			t.Fatalf("UnexpectedErr")
		}
	}
}
