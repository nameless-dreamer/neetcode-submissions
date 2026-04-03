func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	boardLen := len(board)
	for i := 0; i < boardLen; i++ {
		rows[i] = make(map[byte]bool)
    	columns[i] = make(map[byte]bool)
    	boxes[i] = make(map[byte]bool)

		for j := 0; j < boardLen; j++ {
			square := board[i][j]

			if square == '.' {
				continue
			}

			boxIndex := (i / 3) * 3 + (j / 3)

			if rows[i][square] || columns[j][square] || boxes[boxIndex][square] {
				return false
			}

			rows[i][square] = true
			columns[j][square] = true
			boxes[boxIndex][square] = true
		}
	}

	return true
}
