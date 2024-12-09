package main

import (
	"log/slog"
	"os"
	"sort"
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

		lv []int
		rv []int
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			matches := strings.SplitN(row, "   ", 2)

			l, err := strconv.Atoi(matches[0])
			if err != nil {
				slog.Error("failed to parse left", "row", row, "val", matches[0], "err", err)
				os.Exit(1)
			}
			lv = append(lv, l)

			r, err := strconv.Atoi(matches[1])
			if err != nil {
				slog.Error("failed to parse right", "row", row, "val", matches[1], "err", err)
				os.Exit(1)
			}
			rv = append(rv, r)
		}
	})

	sort.Ints(lv)
	sort.Ints(rv)

	for i := 0; i < len(lv); i++ {
		if lv[i] > rv[i] {
			result += lv[i] - rv[i]
		} else {
			result += rv[i] - lv[i]
		}
	}

	return strconv.Itoa(result)
}
