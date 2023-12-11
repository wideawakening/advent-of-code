package challenge

import (
	"fmt"
	"image"
	"math"
	"os"
	"slices"
	"strings"
)

var totalRows int
var totalColumns int

func Star1(inputFileName string) int {

	// TODO manual validation. sample = 9
	galaxies := GetGalaxyLocation(inputFileName)

	// TODO manual validation. sample = 2
	emptyRows := GetEmptyRows(galaxies)

	// TODO manual validation. sample = 3
	emptyColumns := GetEmptyColumns(galaxies)

	combinations := map[image.Point]map[image.Point]int{}

	totalDistance := 0
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1.X != galaxy2.X || galaxy1.Y != galaxy2.Y {

				if combinations[galaxy1][galaxy2] != 0 {
					//fmt.Printf("pre-calculated. galaxy1: %v, galaxy2: %v, distance: %v\n", galaxy1, galaxy2, combinations[galaxy1][galaxy2])
					continue
				}
				if combinations[galaxy2][galaxy1] != 0 {
					//fmt.Printf("pre-calculated. galaxy1: %v, galaxy2: %v, distance: %v\n", galaxy1, galaxy2, combinations[galaxy2][galaxy1])
					continue
				}

				// same row
				if galaxy1.Y == galaxy2.Y {
					distance := int(math.Abs(float64(galaxy1.X - galaxy2.X)))
					expandedDistance := GetEmptyColumnsBetweenTwoPoints(galaxy1, galaxy2, emptyColumns)
					fmt.Printf("same row. galaxy1: %v, galaxy2: %v, distance: %v, expandedDistance: %v\n", galaxy1, galaxy2, distance, len(expandedDistance))
					totalDistance += distance + len(expandedDistance)

					if combinations[galaxy1] == nil {
						combinations[galaxy1] = map[image.Point]int{}
					}
					combinations[galaxy1][galaxy2] = distance + len(expandedDistance)
					continue

				}

				// same column
				if galaxy1.X == galaxy2.X {
					distance := int(math.Abs(float64(galaxy1.Y - galaxy2.Y)))
					expandedDistance := GetEmptyColumnsBetweenTwoPoints(galaxy1, galaxy2, emptyRows)
					fmt.Printf("same column. galaxy1: %v, galaxy2: %v, distance: %v, expandedDistance: %v\n", galaxy1, galaxy2, distance, len(expandedDistance))
					totalDistance += distance + len(expandedDistance)

					if combinations[galaxy1] == nil {
						combinations[galaxy1] = map[image.Point]int{}
					}
					combinations[galaxy1][galaxy2] = distance + len(expandedDistance)
					continue
				}

				// different row and column
				if galaxy1.Y != galaxy2.Y || galaxy1.X != galaxy2.X {
					expandedDistanceRows := GetEmptyRowsBetweenTwoPoints(galaxy1, galaxy2, emptyRows)
					expandedDistanceColumns := GetEmptyColumnsBetweenTwoPoints(galaxy1, galaxy2, emptyColumns)

					distance :=
						int(math.Abs(float64(galaxy1.X-(galaxy2.X+len(expandedDistanceColumns))))) +
							int(math.Abs(float64(galaxy1.Y-(galaxy2.Y+len(expandedDistanceRows)))))

					fmt.Printf("different. galaxy1: %v, galaxy2: %v, distance: %v\n", galaxy1, galaxy2, distance)
					totalDistance += distance

					if combinations[galaxy1] == nil {
						combinations[galaxy1] = map[image.Point]int{}
					}
					combinations[galaxy1][galaxy2] = distance
					continue
				}
			}
		}
	}
	return totalDistance
}

func GetEmptyColumnsBetweenTwoPoints(galaxy1 image.Point, galaxy2 image.Point, emptyColumns []int) []int {
	emptyColumnsBetweenTwoPoints := []int{}
	for _, emptyColumn := range emptyColumns {
		if emptyColumn > galaxy1.X && emptyColumn < galaxy2.X {
			emptyColumnsBetweenTwoPoints = append(emptyColumnsBetweenTwoPoints, emptyColumn)
		}
	}
	return emptyColumnsBetweenTwoPoints
}

func GetEmptyRowsBetweenTwoPoints(galaxy1 image.Point, galaxy2 image.Point, emptyRows []int) []int {
	emptyRowsBetweenTwoPoints := []int{}
	for _, emptyRow := range emptyRows {
		if emptyRow > galaxy1.Y && emptyRow < galaxy2.Y {
			emptyRowsBetweenTwoPoints = append(emptyRowsBetweenTwoPoints, emptyRow)
		}
	}
	return emptyRowsBetweenTwoPoints
}

func GetGalaxyLocation(inputFileName string) []image.Point {
	points := []image.Point{}
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	rows := strings.Split(string(data), "\n")
	totalRows = len(rows)
	totalColumns = len(rows[0])

	for x, dataLine := range rows {
		//fmt.Printf("%s", dataLine)
		for y, dataChar := range dataLine {
			if dataChar == '#' {
				points = append(points, image.Point{x, y})
			}
		}
	}
	return points
}

func GetEmptyRows(galaxies []image.Point) []int {
	emptyRows := []int{}

	galaxyRows := []int{}
	for _, galaxy := range galaxies {
		galaxyRows = append(galaxyRows, galaxy.Y)
	}

	for i := 0; i < totalRows; i++ {
		if !slices.Contains(galaxyRows, i) {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func GetEmptyColumns(galaxies []image.Point) []int {
	emptyColumns := []int{}

	galaxyColumns := []int{}
	for _, galaxy := range galaxies {
		galaxyColumns = append(galaxyColumns, galaxy.X)
	}

	for i := 0; i < totalColumns; i++ {
		if !slices.Contains(galaxyColumns, i) {
			emptyColumns = append(emptyColumns, i)
		}
	}
	return emptyColumns
}
