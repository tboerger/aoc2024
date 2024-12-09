package main

import (
	"log/slog"
	"os"
	"slices"
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
		result  = 0
		rules   []string
		updates [][]int
		order   = make(map[int][]int)
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			if strings.Contains(row, "|") {
				rules = append(rules, row)
			}

			if strings.Contains(row, ",") {
				update := []int{}

				for _, val := range strings.Split(row, ",") {
					ival, _ := strconv.Atoi(val)
					update = append(update, ival)
				}

				updates = append(updates, update)
			}
		}
	})

	for _, rule := range rules {
		parts := strings.SplitN(rule, "|", 2)

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			slog.Error("failed to parse left", "val", parts, "err", err)
			os.Exit(1)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			slog.Error("failed to parse right", "val", parts, "err", err)
			os.Exit(1)
		}

		order[x] = append(order[x], y)
	}

	for _, update := range updates {
		if valid(order, update) {
			continue
		}

		ordered := make([]int, len(update))
		copy(ordered, update)

		slices.SortFunc(ordered, func(a, b int) int {
			for _, ord := range order[a] {
				if ord == b {
					return -1
				}
			}

			return 1
		})

		slog.Debug("", "update", update, "ordered", ordered)

		if valid(order, ordered) {
			result += ordered[len(ordered)/2]
		}
	}

	return strconv.Itoa(result)
}

func valid(order map[int][]int, update []int) bool {
	pos := make(map[int]int)

	for i, page := range update {
		pos[page] = i
	}

	for x, after := range order {
		if posX, exists := pos[x]; exists {
			for _, y := range after {
				if posY, exists := pos[y]; exists {
					if posX >= posY {
						return false
					}
				}
			}
		}
	}

	return true
}
