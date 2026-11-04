package leetcode

func snakesAndLadders(board [][]int) int {
	return bfs(1, board, 0)
}

func bfs(square int, board [][]int, step int) int {
	total := len(board) * len(board)
	type state struct {
		square int
		step   int
	}
	queue := []state{{square, step}}
	visited := make([]bool, total+1)
	visited[square] = true

	for head := 0; head < len(queue); head++ {
		current := queue[head]
		if current.square == total {
			return current.step
		}
		for roll := 1; roll <= 6 && current.square+roll <= total; roll++ {
			nextSquare := current.square + roll
			row, col := position(nextSquare, len(board))
			if board[row][col] != -1 {
				nextSquare = board[row][col]
			}
			if !visited[nextSquare] {
				visited[nextSquare] = true
				queue = append(queue, state{nextSquare, current.step + 1})
			}
		}
	}
	return -1
}

func position(square int, n int) (row, col int) {
	bottomRow := (square - 1) / n
	offset := (square - 1) % n

	row = n - 1 - bottomRow
	col = offset
	if bottomRow%2 == 1 {
		col = n - 1 - col
	}
	return
}
