package main

import (
	"log/slog"
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
		blinks = 75
		stones []string
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			stones = strings.Split(row, " ")
		}
	})

	slog.Debug("init", "stones", stones)

	for i := 0; i < blinks; i++ {
		next := []string{}
		slog.Debug("", "blink", i+1)

		for _, stone := range stones {
			switch {
			case stone == "0":
				next = append(next, "1")

			case len(stone)%2 != 0:
				val, _ := strconv.Atoi(stone)
				next = append(next, strconv.Itoa(val*2024))

			case len(stone)%2 == 0:
				middle := len(stone) / 2
				right := strings.TrimLeft(stone[middle:], "0")

				if len(right) == 0 {
					right = "0"
				}

				next = append(next, stone[:middle], right)

			default:
				next = append(next, stone)
			}
		}

		stones = next
	}

	return strconv.Itoa(len(stones))
}
