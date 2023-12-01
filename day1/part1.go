package day1

import (
	_ "advent2023/logger"
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"unicode"
)

func RunPart1(input string) (int, error) {
	slog.Info("Part 1 -- run")

	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	calibrationSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		first := -1
		last := 0

		for _, r := range line {
			if unicode.IsDigit(r) {
				digit, _ := strconv.Atoi(string(r))
				if first == -1 {
					first = digit
				}
				last = digit
			}
		}

		slog.Debug(fmt.Sprintf("line=%s, first=%d, last=%d, calib=%d",
			line, first, last, first*10+last))

		calibrationSum += first*10 + last
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	slog.Info(fmt.Sprintf("Part 1 -- sum of all of the calibration values: %d", calibrationSum))
	return calibrationSum, nil
}
