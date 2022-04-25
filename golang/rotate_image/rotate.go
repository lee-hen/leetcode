package rotate_image

func Rotate(matrix [][]int)  {
	for layer := 0; layer < len(matrix)/2; layer++ {
		first := layer
		last := len(matrix) - 1 - layer

		for i := layer; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			// left -> top
			matrix[first][i] = matrix[last-offset][first]

			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			// top -> right
			matrix[i][last] = top // right <- saved top
		}
	}
}
