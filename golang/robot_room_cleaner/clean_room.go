package robot_room_cleaner

import "fmt"

func CleanRoom(robot *Robot) {
	visited1 := make(map[string]bool)
	dfs(robot, visited1)

	visited2 := make(map[string]bool)
	backtrack(robot, 0, 0, 0, visited2)
}

func dfs(robot *Robot, visited map[string]bool) {
	key := fmt.Sprintf("%d-%d", robot.row, robot.col)
	if visited[key] {
		return
	}

	robot.Clean()
	visited[key] = true

	if robot.Move() {
		dfs(robot, visited)

		robot.TurnRight()
		robot.TurnRight()
		robot.Move()
		robot.TurnRight()
		robot.TurnRight()
	}

	robot.TurnLeft()
	if robot.Move() {
		dfs(robot, visited)
		robot.TurnRight()
		robot.TurnRight()
		robot.Move()
		robot.TurnLeft()
	} else {
		robot.TurnRight()
	}

	robot.TurnRight()
	if robot.Move() {
		dfs(robot, visited)
		robot.TurnLeft()
		robot.TurnLeft()
		robot.Move()
		robot.TurnRight()
	} else {
		robot.TurnLeft()
	}

	robot.TurnLeft()
	robot.TurnLeft()
	if robot.Move() {
		dfs(robot, visited)
		robot.TurnRight()
		robot.TurnRight()
		robot.Move()
	} else {
		robot.TurnRight()
		robot.TurnRight()
	}
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func backtrack(robot *Robot, row, col, d int, visited map[string]bool) {
	key := fmt.Sprintf("%d-%d", row, col)

	robot.Clean()
	visited[key] = true

	// going clockwise : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
	for i := 0; i < 4; i++ {
		newD := (d + i) % 4
		newRow := row + directions[newD][0]
		newCol := col + directions[newD][1]

		newKey := fmt.Sprintf("%d-%d", newRow, newCol)

		if !visited[newKey] && robot.Move() {
			backtrack(robot, newRow, newCol, newD, visited)
			goBack(robot)
		}

		// turn the robot following chosen direction : clockwise
		robot.TurnRight()
	}
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
