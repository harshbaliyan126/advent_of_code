package dec

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Dec02AocPart1() int {

	workingDir, _ := os.Getwd()
	filePath, _ := filepath.Abs(filepath.Join(workingDir, "internals", "in.txt"))
	file, _ := os.Open(filePath)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 1

	totalCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		format := fmt.Sprintf("Game %d:", i)
		parts := strings.SplitAfterN(line, format, 2)

		gameSession := strings.Split(parts[1], ";")

		isPossible := true

		colorCount := map[string]int{"red": 0, "blue": 0, "green": 0}

		for _, session := range gameSession {

			ballCounts := strings.Split(strings.TrimSpace(session), ",")

			for _, ballCount := range ballCounts {
				ballCount := strings.TrimSpace(ballCount)

				var ballColor []string = strings.Split(ballCount, " ")
				num, _ := strconv.Atoi(ballColor[0])

				colorCount[ballColor[1]] = num

			}

			if colorCount["red"] > 12 || colorCount["blue"] > 14 || colorCount["green"] > 13 {
				isPossible = false
				break
			}

			colorCount["red"] = 0
			colorCount["blue"] = 0
			colorCount["green"] = 0

		}

		if isPossible {
			totalCount += i
		}

		// fmt.Println(i, " ", colorCount)
		// if  colorCount["red"] <= 12 && colorCount["blue"] <= 14 && colorCount["green"] <= 13 {
		//   totalCount += i
		//   fmt.Println("i: ", i)
		// }

		colorCount["red"] = 0
		colorCount["blue"] = 0
		colorCount["green"] = 0

		i++
	}

	return totalCount

}
