func isValidSudoku(board [][]byte) bool {
	// Check rows
	for r := 0; r < 9; r++ {
		seen := make(map[byte]bool)

		for c := 0; c < 9; c++ {
			val := board[r][c]

			if val == '.' {
				continue
			}

			if seen[val] {
				return false
			}

			seen[val] = true
		}
	}

	// Check columns
	for c := 0; c < 9; c++ {
		seen := make(map[byte]bool)

		for r := 0; r < 9; r++ {
			val := board[r][c]

			if val == '.' {
				continue
			}

			if seen[val] {
				return false
			}

			seen[val] = true
		}
	}

	// Check 3x3 boxes
	for boxRow := 0; boxRow < 9; boxRow += 3 {
		for boxCol := 0; boxCol < 9; boxCol += 3 {
			seen := make(map[byte]bool)

			for r := boxRow; r < boxRow+3; r++ {
				for c := boxCol; c < boxCol+3; c++ {
					val := board[r][c]

					if val == '.' {
						continue
					}

					if seen[val] {
						return false
					}

					seen[val] = true
				}
			}
		}
	}

	return true
}