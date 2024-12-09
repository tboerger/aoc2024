package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/tboerger/aoc2024/pkg"
)

func main() {
	testing := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "--testing":
			testing = true
		}
	}

	slog.Info("finished", "result", run(testing))
}

func run(testing bool) string {
	var (
		result = 0
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			slog.Info(row)
		}
	})

	return strconv.Itoa(result)
}
