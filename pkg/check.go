package pkg

import (
	"log/slog"
	"os"
)

func Check(e error) {
	if e != nil {
		slog.Error("", "err", e)
		os.Exit(1)
	}
}
