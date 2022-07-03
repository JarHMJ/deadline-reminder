package cmd

import (
	"testing"
	"time"
)

func Test_calculateDays(t *testing.T) {
	type args struct {
		now      *time.Time
		deadline *time.Time
	}
	now := time.Date(2022, 5, 1, 12, 3, 41, 123, time.Local)
	deadline := time.Date(2022, 5, 2, 12, 3, 41, 123, time.Local)

	now1 := time.Date(2022, 5, 1, 12, 3, 41, 123, time.Local)
	deadline1 := time.Date(2022, 5, 1, 12, 3, 41, 123, time.Local)

	now2 := time.Date(2022, 5, 1, 12, 3, 41, 123, time.Local)
	deadline2 := time.Date(2022, 6, 1, 12, 3, 41, 123, time.Local)

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "one day", args: args{now: &now, deadline: &deadline}, want: 1},
		{name: "zero day", args: args{now: &now1, deadline: &deadline1}, want: 0},
		{name: "31 day", args: args{now: &now2, deadline: &deadline2}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDays(tt.args.now, tt.args.deadline); got != tt.want {
				t.Errorf("calculateDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
