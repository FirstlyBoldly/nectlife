package must

import (
	"log/slog"
	"strconv"
)

func AtoiEnforce(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		slog.Error("Value " + str + " could not be converted into type integer")
	}

	return val
}
