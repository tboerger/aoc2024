package main

import (
	"log/slog"
	"os"
	"regexp"
	"strconv"

	"github.com/tboerger/aoc2024/pkg"
)

var (
	mulRegex  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex   = regexp.MustCompile(`\bdo\(\)`)
	dontRegex = regexp.MustCompile(`\bdon't\(\)`)
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
		result   = 0
		multiply = true
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			for len(row) > 0 {
				if match := doRegex.FindStringIndex(row); match != nil && match[0] == 0 {
					slog.Debug("do", "match", match)
					multiply = true
					row = row[match[1]:]
				} else if match := dontRegex.FindStringIndex(row); match != nil && match[0] == 0 {
					slog.Debug("dont", "match", match)
					multiply = false
					row = row[match[1]:]
				} else if match := mulRegex.FindStringSubmatchIndex(row); match != nil && match[0] == 0 {
					if multiply {
						l, err := strconv.Atoi(row[match[2]:match[3]])
						if err != nil {
							slog.Error("failed to parse left", "val", row[match[2]:match[3]], "err", err)
							os.Exit(1)
						}

						r, err := strconv.Atoi(row[match[4]:match[5]])
						if err != nil {
							slog.Error("failed to parse right", "val", row[match[4]:match[5]], "err", err)
							os.Exit(1)
						}

						slog.Debug("multiply", "left", row[match[2]:match[3]], "right", row[match[4]:match[5]])
						result += l * r
					}

					row = row[match[1]:]
				} else {
					row = row[1:]
				}
			}
		}
	})

	return strconv.Itoa(result)
}
