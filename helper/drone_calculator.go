package helper

import (
	"math"

	"github.com/SawitProRecruitment/UserService/repository"
)

// CalculateDroneDistance calculates the total distance the drone needs to travel with a max distance limit.
func CalculateDroneDistance(width, height, maxDistance int, trees []repository.EstateTree) int {
	// Create a map for quick lookup of tree heights by their plot location.
	treeMap := make(map[[2]int]int)
	for _, tree := range trees {
		treeMap[[2]int{tree.X, tree.Y}] = tree.Height
	}

	horizontalDistance := 0
	verticalDistance := 0
	totalDistance := 0

	currentHeight := 0
	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			// Calculate horizontal movement
			if col > 1 {
				horizontalDistance += 10
			}
			if totalDistance+horizontalDistance > maxDistance {
				// If max distance is reached, return to ground level
				verticalDistance += currentHeight
				totalDistance += horizontalDistance + verticalDistance
				return totalDistance
			}

			// Calculate vertical movement
			targetHeight := 1 // Default to 1 meter above ground.
			if height, exists := treeMap[[2]int{row, col}]; exists {
				targetHeight = height + 1 // 1 meter above the tree.
			}
			verticalDistance += int(math.Abs(float64(targetHeight - currentHeight)))
			currentHeight = targetHeight

			if totalDistance+horizontalDistance+verticalDistance > maxDistance {
				// If max distance is reached, return to ground level
				verticalDistance += currentHeight
				totalDistance += horizontalDistance + verticalDistance
				return totalDistance
			}
		}

		// Move north to the next row
		if row != height {
			horizontalDistance += 10 // Move north by one plot (10 meters).
			if totalDistance+horizontalDistance > maxDistance {
				// If max distance is reached, return to ground level
				verticalDistance += currentHeight
				totalDistance += horizontalDistance + verticalDistance
				return totalDistance
			}
		}
	}

	// Final descent to ground level
	verticalDistance += currentHeight
	totalDistance += horizontalDistance + verticalDistance
	return totalDistance
}
