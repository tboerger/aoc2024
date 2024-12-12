package main

import (
	"fmt"
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
		diskmap []int
		blocks  []int
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			for _, char := range strings.Split(row, "") {
				num, _ := strconv.Atoi(char)
				diskmap = append(diskmap, num)
			}
		}
	})

	fmt.Println(len(diskmap))

	blocks = parseDiskmap(diskmap)
	slog.Debug("parsed", "count", len(blocks), "val", blocks)

	blocks = compactDisk(blocks)
	slog.Debug("compact", "count", len(blocks), "val", blocks)

	return strconv.Itoa(calculateChecksum(blocks))
}

func parseDiskmap(diskmap []int) []int {
	result := []int{}

	for i, val := range diskmap {
		if i%2 == 0 {
			slog.Debug("data", "num", val)
			result = append(result, slices.Repeat([]int{i / 2}, val)...)
		} else {
			slog.Debug("free", "num", val)
			result = append(result, slices.Repeat([]int{-1}, val)...)
		}
	}

	return result
}

func compactDisk(blocks []int) []int {
	free := slices.Index(blocks, -1)

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] == -1 {
			continue
		}

		if free > i {
			break
		}

		slog.Debug("a", "blocks", blocks, "free", free, "i", i)

		if blocks[i] >= 0 {
			blocks[free], blocks[i] = blocks[i], blocks[free]
			free = slices.Index(blocks, -1)
		}

		slog.Debug("a", "blocks", blocks, "free", free, "i", i)
	}

	return blocks
}

func calculateChecksum(blocks []int) int {
	result := 0

	for pos, val := range blocks {
		if val != -1 {
			result += pos * val
		}
	}

	return result
}
