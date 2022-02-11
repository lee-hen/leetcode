package robot_room_cleaner

import (
	"math/rand"
	"time"
)

// Robot
// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
	row, col int
}

// Move
// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}

// TurnLeft
// Robot will stay in the same cell after calling TurnLeft
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft() {
}

// TurnRight
// Robot will stay in the same cell after calling TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnRight() {
}

// Clean the current cell.
func (robot *Robot) Clean() {
}
