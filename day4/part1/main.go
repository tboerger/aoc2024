package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/tboerger/aoc2024/pkg"
)

const (
	term       = "XMAS"
	termLength = len(term)
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
		grid   []string

		dirs = [][]int{
			{0, 1},   // right
			{0, -1},  // left
			{1, 0},   // down
			{-1, 0},  // up
			{1, 1},   // down-right
			{-1, -1}, // up-left
			{1, -1},  // down-left
			{-1, 1},  // up-right
		}
	)

	pkg.Parse(testing, func(rows []string) {
		grid = rows
	})

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {

			for _, d := range dirs {
				if match(grid, r, c, d[0], d[1]) {
					result += 1
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func match(grid []string, r, c, dr, dc int) bool {
	for i := 0; i < termLength; i++ {
		nr, nc := r+dr*i, c+dc*i

		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] != term[i] {
			return false
		}
	}

	return true
}
