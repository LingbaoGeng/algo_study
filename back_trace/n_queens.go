package back_trace

var ans1 [][]string

func SolveNQueens(n int) [][]string {
	ans1 = make([][]string, 0)
	chessBoard := make([][]byte, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = '.'
		}
	}
	backTrace1(chessBoard, n, 0)
	return ans1
}

func backTrace1(chessBoard [][]byte, n int, row int) {
	if row == n {
		tempBoard := make([]string, 0)
		for _, v := range chessBoard {
			tempBoard = append(tempBoard, string(v))
		}
		ans1 = append(ans1, tempBoard)
		return
	}
	for col := 0; col < n; col++ {
		if !isValid(chessBoard, n, row, col) {
			continue
		}
		chessBoard[row][col] = 'Q'
		backTrace1(chessBoard, n, row+1)
		chessBoard[row][col] = '.'
	}
}

func isValid(chessBoard [][]byte, n, row, col int) bool {
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == 'Q' {
			return false
		}
	}
	i := row - 1
	j := col + 1

	for i >= 0 && j < n {
		if chessBoard[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if chessBoard[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}
