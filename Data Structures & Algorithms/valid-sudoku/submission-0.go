func isValidSudoku(board [][]byte) bool {
    for j := 0; j < 9; j++ {
        seen := make(map[byte]bool)
        for i := 0; i < 9; i++ {
            val := board[i][j]
            if val == '.' {
                continue
            }
            if seen[val] {
                return false
            }
            seen[val] = true
        }
    }

    for i := 0; i < 9; i++ {
        seen := make(map[byte]bool)
        for j := 0; j < 9; j++ {
            val := board[i][j]
            if val == '.' {
                continue
            }
            if seen[val] {
                return false
            }
            seen[val] = true
        }
    }

    for boxRow := 0; boxRow < 9; boxRow += 3 {
        for boxCol := 0; boxCol < 9; boxCol += 3 {

            seen := make(map[byte]bool)

            for i := 0; i < 3; i++ {
                for j := 0; j < 3; j++ {

                    val := board[boxRow+i][boxCol+j]
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