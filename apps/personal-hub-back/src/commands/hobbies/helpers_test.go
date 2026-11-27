package hobbies_commands_test

import "time"

var (
	date1 = time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2026, 2, 20, 0, 0, 0, 0, time.UTC)
)

func ptr[T any](v T) *T { return &v }
