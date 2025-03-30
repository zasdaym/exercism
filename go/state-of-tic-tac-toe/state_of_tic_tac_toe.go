package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	var (
		crosses int
		circles int
	)

	for _, row := range board {
		for _, col := range row {
			switch col {
			case 'X':
				crosses++
			case 'O':
				circles++
			}
		}
	}

	var (
		crossWins  = checkLine(board, 'X')
		circleWins = checkLine(board, 'O')
	)

	if crossWins && circleWins {
		return "", fmt.Errorf("invalid board")
	}

	if crossWins || circleWins {
		return Win, nil
	}

	if crosses+circles == 9 {
		return Draw, nil
	}

	// Cross goes first, so it's impossible we have more circles than crosses.
	if crosses < circles {
		return "", fmt.Errorf("invalid board")
	}

	// Turns taken alternately
	if crosses-circles > 1 {
		return "", fmt.Errorf("invalid board")
	}

	return Ongoing, nil
}

func checkLine(board []string, player byte) bool {
	lines := [][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
	}

	for _, line := range lines {
		if board[line[0][0]][line[0][1]] == player &&
			board[line[1][0]][line[1][1]] == player &&
			board[line[2][0]][line[2][1]] == player {
			return true
		}
	}

	return false
}
