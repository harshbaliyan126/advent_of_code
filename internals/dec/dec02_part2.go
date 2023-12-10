package dec

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Dec02AocPart2() int {

	workingDir, _ := os.Getwd()

	filePath, _ := filepath.Abs(filepath.Join(workingDir, "internals", "in.txt"))

	file, _ := os.Open(filePath)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 1

	powerSetSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		format := fmt.Sprintf("Game %d:", i)
		parts := strings.SplitAfterN(line, format, 2)

		gameSession := strings.Split(parts[1], ";")

		colorCount := map[string]int{"red": 0, "blue": 0, "green": 0}

		for _, session := range gameSession {

			ballCounts := strings.Split(strings.TrimSpace(session), ",")

			for _, ballCount := range ballCounts {
				ballCount := strings.TrimSpace(ballCount)

				if ballCount != "" {
					var ballColor []string = strings.Split(ballCount, " ")
					count, _ := strconv.Atoi(ballColor[0])

					colorCount[ballColor[1]] = max(colorCount[ballColor[1]], count)

				}

			}

		}

		powerSetSum += (colorCount["red"] * colorCount["blue"] * colorCount["green"])

		colorCount["red"] = 0
		colorCount["blue"] = 0
		colorCount["green"] = 0

		i++
	}

	return powerSetSum

}
