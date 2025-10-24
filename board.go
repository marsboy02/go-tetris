package main

import (
	"math/rand"
	"time"
)

const (
	BoardWidth  = 10
	BoardHeight = 20
)

// Board represents the game board
type Board struct {
	Grid          [][]Color
	Current       *Tetromino
	Next          *Tetromino
	Score         int
	Lines         int
	GameOver      bool
	Level         int
	rand          *rand.Rand
}

// NewBoard creates a new game board
func NewBoard() *Board {
	grid := make([][]Color, BoardHeight)
	for i := range grid {
		grid[i] = make([]Color, BoardWidth)
	}

	b := &Board{
		Grid:  grid,
		Score: 0,
		Lines: 0,
		Level: 1,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	b.Next = b.randomTetromino()
	b.SpawnTetromino()
	return b
}

// randomTetromino generates a random tetromino
func (b *Board) randomTetromino() *Tetromino {
	types := []TetrominoType{
		TetrominoI, TetrominoO, TetrominoT,
		TetrominoS, TetrominoZ, TetrominoJ, TetrominoL,
	}
	t := types[b.rand.Intn(len(types))]
	return NewTetromino(t)
}

// SpawnTetromino spawns a new tetromino
func (b *Board) SpawnTetromino() {
	b.Current = b.Next
	b.Next = b.randomTetromino()

	// Check if game over
	if b.CheckCollision(b.Current, b.Current.X, b.Current.Y) {
		b.GameOver = true
	}
}

// CheckCollision checks if a tetromino collides with the board
func (b *Board) CheckCollision(t *Tetromino, x, y int) bool {
	for row := 0; row < len(t.Shape); row++ {
		for col := 0; col < len(t.Shape[row]); col++ {
			if t.Shape[row][col] == 0 {
				continue
			}

			newX := x + col
			newY := y + row

			// Check bounds
			if newX < 0 || newX >= BoardWidth || newY >= BoardHeight {
				return true
			}

			// Check if position is occupied (only if within bounds)
			if newY >= 0 && b.Grid[newY][newX] != ColorNone {
				return true
			}
		}
	}
	return false
}

// MoveTetromino moves the current tetromino
func (b *Board) MoveTetromino(dx, dy int) bool {
	newX := b.Current.X + dx
	newY := b.Current.Y + dy

	if !b.CheckCollision(b.Current, newX, newY) {
		b.Current.X = newX
		b.Current.Y = newY
		return true
	}
	return false
}

// RotateTetromino rotates the current tetromino
func (b *Board) RotateTetromino() bool {
	b.Current.Rotate()
	if b.CheckCollision(b.Current, b.Current.X, b.Current.Y) {
		b.Current.RotateBack()
		return false
	}
	return true
}

// LockTetromino locks the current tetromino to the board
func (b *Board) LockTetromino() {
	for row := 0; row < len(b.Current.Shape); row++ {
		for col := 0; col < len(b.Current.Shape[row]); col++ {
			if b.Current.Shape[row][col] == 0 {
				continue
			}

			x := b.Current.X + col
			y := b.Current.Y + row

			if y >= 0 && y < BoardHeight && x >= 0 && x < BoardWidth {
				b.Grid[y][x] = b.Current.Color
			}
		}
	}

	b.ClearLines()
	b.SpawnTetromino()
}

// ClearLines clears completed lines
func (b *Board) ClearLines() {
	linesCleared := 0

	for row := BoardHeight - 1; row >= 0; row-- {
		full := true
		for col := 0; col < BoardWidth; col++ {
			if b.Grid[row][col] == ColorNone {
				full = false
				break
			}
		}

		if full {
			// Remove the line
			b.Grid = append(b.Grid[:row], b.Grid[row+1:]...)
			// Add new empty line at top
			b.Grid = append([][]Color{make([]Color, BoardWidth)}, b.Grid...)
			linesCleared++
			row++ // Check the same row again
		}
	}

	if linesCleared > 0 {
		b.Lines += linesCleared
		// Scoring: 100 for 1 line, 300 for 2, 500 for 3, 800 for 4
		scores := []int{0, 100, 300, 500, 800}
		b.Score += scores[linesCleared] * b.Level

		// Level up every 10 lines
		b.Level = b.Lines/10 + 1
	}
}

// DropTetromino drops the current tetromino one step
func (b *Board) DropTetromino() bool {
	if b.MoveTetromino(0, 1) {
		return true
	}
	b.LockTetromino()
	return false
}

// HardDrop drops the tetromino instantly
func (b *Board) HardDrop() {
	for b.MoveTetromino(0, 1) {
	}
	b.LockTetromino()
}
