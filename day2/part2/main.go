package main

import (
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"

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
			matches := strings.Split(row, " ")
			vals := []int{}

			for _, match := range matches {
				v, err := strconv.Atoi(match)
				if err != nil {
					slog.Error("failed to parse int", "row", row, "val", match, "err", err)
					os.Exit(1)
				}
				vals = append(vals, v)
			}

			if dampener(vals) {
				result += 1
			}
		}
	})

	return strconv.Itoa(result)
}

func dampener(report []int) bool {
	if safe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		if safe(append(append([]int{}, report[:i]...), report[i+1:]...)) {
			return true
		}
	}

	return false
}

func safe(report []int) bool {
	inc := true
	dec := true
	big := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if abs := math.Abs(float64(diff)); abs < 1 || abs > 3 {
			slog.Debug("baddiff", "prev", report[i-1], "curr", report[i], "diff", diff, "report", report)
			big = true
			continue
		}

		if diff > 0 {
			slog.Debug("increment", "prev", report[i-1], "curr", report[i], "diff", diff, "report", report)
			dec = false
		}

		if diff < 0 {
			slog.Debug("decrement", "prev", report[i-1], "curr", report[i], "diff", diff, "report", report)
			inc = false
		}
	}

	slog.Debug("report", "safe", !big && (inc || dec))
	return !big && (inc || dec)
}
