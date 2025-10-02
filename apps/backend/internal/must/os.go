package must

import (
	"log/slog"
	"os"
)

func LookupEnvEnforce(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		slog.Error("Key " + val + " does not exist")
	}

	return val
}
