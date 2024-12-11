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
		blinks = 25
		stones []int
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			for _, number := range strings.Split(row, " ") {
				val, _ := strconv.Atoi(number)
				stones = append(stones, val)
			}
		}
	})

	slog.Debug("init", "stones", stones)

	for i := 0; i < blinks; i++ {
		next := []int{}
		slog.Debug("", "blink", i+1)

		for _, stone := range stones {
			switch {
			case stone == 0:
				next = append(next, 1)

			case len(strconv.Itoa(stone))%2 == 0:
				left, right := splitNumber(stone)
				next = append(next, left, right)

			default:
				next = append(next, stone*2024)
			}
		}

		stones = next
	}

	return strconv.Itoa(len(stones))
}

func splitNumber(stone int) (int, int) {
	splitPoint := countDigits(stone) / 2
	splitDivisor := int(math.Pow10(splitPoint))

	return stone / splitDivisor, stone % splitDivisor
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(float64(num)))) + 1
}
