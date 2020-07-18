package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var board [4][4]int

const (
	seperator = "+------+------+------+------+\n"
)

func main() {
	initBoard()

	reader := bufio.NewReader(os.Stdin)
	for {
		insertNew()
		printBoard()
		for {
			fmt.Printf("Enter a move: a:left, s:down, d:right, w:up\n")
			input, _ := reader.ReadString('\n')

			success := false

			switch strings.ToUpper(string([]byte(input)[0])) {
			case "A":
				success = moveLeft()
			case "D":
				success = moveRight()
			case "W":
				success = moveUp()
			case "S":
				success = moveDown()
			default:
				fmt.Printf("Unknown input!\n")
				continue
			}

			if success {
				fmt.Printf("Moved successfully\n")
				break
			} else {
				fmt.Printf("Move was not successful\n")
				printBoard()
			}
		}

		if gameOver() {
			break
		}
	}

	fmt.Printf("Game over\n")
}

func initBoard() {
	board = [4][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
}

func printBoard() {
	fmt.Printf(seperator)
	for row := 0; row < 4; row++ {
		line := "|"
		for col := 0; col < 4; col++ {
			if board[row][col] == 0 {
				line += "      |"
			} else {
				line += fmt.Sprintf(" %4d |", board[row][col])
			}
		}
		fmt.Println(line)
		fmt.Printf(seperator)
	}
}

// insertNew inserts a random 2 or 4 to the an empty cell
func insertNew() {
	n := rand.Intn(2)*2 + 2
	for {
		row := rand.Intn(4)
		col := rand.Intn(4)
		if board[row][col] == 0 {
			fmt.Printf("Inserting [%d][%d] %d\n", row, col, n)
			board[row][col] = n
			return
		}
	}
}

// gameOver Checks the board to see if no space left, i.e game over
func gameOver() bool {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if board[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

// findValue finds a non-zero value, and return value & col
func findValue(row, col int) (int, int) {
	for ; col < 4; col++ {
		if board[row][col] > 0 {
			return board[row][col], col
		}
	}
	return 0, col
}

func moveLeft() bool {
	success := false

	for row := 0; row < 4; row++ {
		headCell := 0
		valLeft := 0
		valRight := 0
		for currentCell := 0; currentCell < 4; currentCell++ {
			valLeft, headCell = findValue(row, headCell)
			if valLeft == 0 { // If not found, cell is zero
				board[row][currentCell] = 0
				continue
			}

			if headCell > currentCell { // Meaning we had empty cells or merged cells in between
				success = true
			}
			board[row][currentCell] = valLeft

			headCell++
			valRight, headCell = findValue(row, headCell)
			if valLeft == valRight { // Merge
				board[row][currentCell] = valLeft * 2
				board[row][headCell] = 0
				success = true
			}
		}
	}

	return success
}

func moveRight() bool {
	rotateBoard()
	rotateBoard()
	success := moveLeft()
	rotateBoard()
	rotateBoard()
	return success
}

func moveUp() bool {
	rotateBoard()
	rotateBoard()
	rotateBoard()
	success := moveLeft()
	rotateBoard()
	return success
}

func moveDown() bool {
	rotateBoard()
	success := moveLeft()
	rotateBoard()
	rotateBoard()
	rotateBoard()
	return success
}

// Rotate board clockwise 90 degrees
// Before:
// 00 01 02 03
// 10 11 12 13
// 20 21 22 23
// 30 31 32 33
// After:
// 30 20 10 00
// 31 21 11 01
// 32 22 12 02
// 33 23 13 03
func rotateBoard() {
	newBoard := [4][4]int{}
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			newBoard[row][col] = board[3-col][row]
		}
	}
	board = newBoard
}
