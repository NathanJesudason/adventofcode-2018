package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	count := 0
	fabric := [1000][1000]int{}
	content, _ := ioutil.ReadFile(inputFile)
	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, line := range lines {
		if line != "" {
			part := re.FindAllStringSubmatch(line, -1)
			dx, _ := strconv.Atoi(part[0][2])
			dy, _ := strconv.Atoi(part[0][3])
			x, _ := strconv.Atoi(part[0][4])
			y, _ := strconv.Atoi(part[0][5])
			for j := dy; j < y+dy; j++ {
				for i := dx; i < x+dx; i++ {
					fabric[i][j]++
					if fabric[i][j] == 2 {
						count++
					}
				}
			}

		}

	}
	println("overlapping area: ", count)
claims:
	for _, line := range lines {
		if line != "" {
			part := re.FindAllStringSubmatch(line, -1)
			claim, _ := strconv.Atoi(part[0][1])
			dx, _ := strconv.Atoi(part[0][2])
			dy, _ := strconv.Atoi(part[0][3])
			x, _ := strconv.Atoi(part[0][4])
			y, _ := strconv.Atoi(part[0][5])
			for j := dy; j < y+dy; j++ {
				for i := dx; i < x+dx; i++ {
					if fabric[i][j] > 1 {
						continue claims
					}
				}
			}
			println("The only good claim: ", claim)
		}
	}
}
