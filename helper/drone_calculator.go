package helper

import (
	"math"

	"github.com/SawitProRecruitment/UserService/repository"
)

type Result struct {
	TotalDistance int
	FinalRow      int
	FinalCol      int
	FinalHeight   int
}

// CalculateDroneDistance calculates the total distance the drone needs to travel with a max distance limit and returns the resting point.
func CalculateDroneWithMaxDistance(width, height, maxDistance int, trees []repository.EstateTree) Result {
	// Create a map for quick lookup of tree heights by their plot location.
	treeMap := make(map[[2]int]int)
	for _, tree := range trees {
		treeMap[[2]int{tree.X, tree.Y}] = tree.Height
	}

	horizontalDistance := 0
	verticalDistance := 0
	totalDistance := 0

	currentHeight := 0
	finalRow := 1
	finalCol := 1

	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			// Calculate horizontal movement.
			if col > 1 {
				horizontalDistance += 10
			}
			if maxDistance != 0 && totalDistance+horizontalDistance+verticalDistance > maxDistance {
				// If max distance is reached, return the current position and distance.
				return Result{TotalDistance: totalDistance, FinalRow: finalRow, FinalCol: finalCol, FinalHeight: currentHeight}
			}

			// Calculate vertical movement.
			targetHeight := 1 // Default to 1 meter above ground.
			if height, exists := treeMap[[2]int{row, col}]; exists {
				targetHeight = height + 1 // 1 meter above the tree.
			}
			verticalDistance += int(math.Abs(float64(targetHeight - currentHeight)))
			currentHeight = targetHeight

			if maxDistance != 0 && totalDistance+horizontalDistance+verticalDistance > maxDistance {
				// If max distance is reached, return the current position and distance.
				return Result{TotalDistance: totalDistance, FinalRow: finalRow, FinalCol: finalCol, FinalHeight: currentHeight}
			}

			finalRow = row
			finalCol = col
		}

		// Move north to the next row.
		if row != height {
			horizontalDistance += 10 // Move north by one plot (10 meters).
			if maxDistance != 0 && totalDistance+horizontalDistance+verticalDistance > maxDistance {
				// If max distance is reached, return the current position and distance.
				return Result{TotalDistance: totalDistance, FinalRow: finalRow, FinalCol: finalCol, FinalHeight: currentHeight}
			}
		}
	}

	// Final descent to ground level.
	verticalDistance += currentHeight
	totalDistance += horizontalDistance + verticalDistance

	return Result{TotalDistance: totalDistance, FinalRow: finalRow, FinalCol: finalCol, FinalHeight: 0}
}
