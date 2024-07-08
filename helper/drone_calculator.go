package helper

import (
	"log"
	"math"

	"github.com/SawitProRecruitment/UserService/repository"
)

var DefaultMaxDistance int = 100000000000

// CalculateDroneDistance calculates the total distance the drone needs to travel with a max distance limit.
func CalculateDroneDistance(width, height int, trees []repository.EstateTree) int {
	log.Println(width)
	log.Println(height)

	// Create a map for quick lookup of tree heights by their plot location.
	treeMap := make(map[[2]int]int)
	for _, tree := range trees {
		treeMap[[2]int{tree.X, tree.Y}] = tree.Height
		log.Println(tree.X, tree.Y, tree.Height)
	}

	horizontalDistance := 0
	verticalDistance := 0

	// Calculate horizontal distance.
	totalRows := height
	for row := 0; row < totalRows; row++ {
		horizontalDistance += (width - 1) * 10 // Eastward or westward travel (10 meters per plot).
		if row != totalRows-1 {
			horizontalDistance += 10 // Move north by one plot (10 meters).
		}
	}

	// Add return distance if the final row was traveled east to west.
	if totalRows%2 == 1 {
		horizontalDistance += (width - 1) * 10
	}

	// Calculate vertical distance.
	currentHeight := 0
	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			targetHeight := 1 // Default to 1 meter above ground.
			if height, exists := treeMap[[2]int{row, col}]; exists {
				targetHeight = height + 1 // 1 meter above the tree.
			}
			verticalDistance += int(math.Abs(float64(targetHeight - currentHeight)))
			currentHeight = targetHeight
		}
	}

	// Final descent.
	verticalDistance += currentHeight

	// Total distance.
	totalDistance := horizontalDistance + verticalDistance
	return totalDistance
}
