package gameOfLife

var board = createBoard()

func createBoard() [][]bool {
    board := make([][]bool, 100)
    for i := range(board) {
        board[i] =make([]bool, 100)
    }

    board[1][2] = true
    board[2][3] = true
    board[3][1] = true
    board[3][2] = true
    board[3][3] = true

    return board
}

func checkNeighbor(x, y int) bool {
    count := 0
    for i := x-1; i <= x+1; i++ {
        if i < 0 || i > 99 {
            continue
        }
        for j := y-1; j <= y+1; j++ {
            if j < 0 || j > 99 {
                continue
            }

            if i == x && j == y {
                continue
            }

            if board[i][j] {
                count++
            }
        }
    }

    if count < 2 || count > 3 {
        return false
    }

    if count == 3 {
        return true
    }

    return board[x][y]
}

func NextBoardState() [][]bool {
    var nextBoard = createBoard()
    for i := range(board) {
        for j := range(board[i]) {
            nextBoard[i][j] = checkNeighbor(i, j)
        }
    }
    copy(board, nextBoard)
    return nextBoard
}
