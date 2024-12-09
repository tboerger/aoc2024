package main

import (
	"log/slog"
	"os"
	"regexp"
	"strconv"

	"github.com/tboerger/aoc2024/pkg"
)

var (
	mulRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
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
			matches := mulRegex.FindAllStringSubmatch(row, -1)
			slog.Debug("matches", "values", matches)

			for _, match := range matches {
				l, err := strconv.Atoi(match[1])
				if err != nil {
					slog.Error("failed to parse left", "val", match[1], "err", err)
					os.Exit(1)
				}

				r, err := strconv.Atoi(match[2])
				if err != nil {
					slog.Error("failed to parse right", "val", match[2], "err", err)
					os.Exit(1)
				}

				result += l * r
			}
		}
	})

	return strconv.Itoa(result)
}
