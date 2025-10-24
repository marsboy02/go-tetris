package main

// Color represents the color of a tetromino
type Color int

const (
	ColorNone Color = iota
	ColorCyan    // I
	ColorYellow  // O
	ColorPurple  // T
	ColorGreen   // S
	ColorRed     // Z
	ColorBlue    // J
	ColorOrange  // L
)

// TetrominoType represents the type of tetromino
type TetrominoType int

const (
	TetrominoI TetrominoType = iota
	TetrominoO
	TetrominoT
	TetrominoS
	TetrominoZ
	TetrominoJ
	TetrominoL
)

// Tetromino represents a tetromino piece
type Tetromino struct {
	Type     TetrominoType
	Shape    [][]int
	Color    Color
	X, Y     int
	Rotation int
}

// GetShapes returns all rotation states for each tetromino type
func GetShapes(t TetrominoType) [][][]int {
	switch t {
	case TetrominoI:
		return [][][]int{
			{{1, 1, 1, 1}},
			{{1}, {1}, {1}, {1}},
			{{1, 1, 1, 1}},
			{{1}, {1}, {1}, {1}},
		}
	case TetrominoO:
		return [][][]int{
			{{1, 1}, {1, 1}},
			{{1, 1}, {1, 1}},
			{{1, 1}, {1, 1}},
			{{1, 1}, {1, 1}},
		}
	case TetrominoT:
		return [][][]int{
			{{0, 1, 0}, {1, 1, 1}},
			{{1, 0}, {1, 1}, {1, 0}},
			{{1, 1, 1}, {0, 1, 0}},
			{{0, 1}, {1, 1}, {0, 1}},
		}
	case TetrominoS:
		return [][][]int{
			{{0, 1, 1}, {1, 1, 0}},
			{{1, 0}, {1, 1}, {0, 1}},
			{{0, 1, 1}, {1, 1, 0}},
			{{1, 0}, {1, 1}, {0, 1}},
		}
	case TetrominoZ:
		return [][][]int{
			{{1, 1, 0}, {0, 1, 1}},
			{{0, 1}, {1, 1}, {1, 0}},
			{{1, 1, 0}, {0, 1, 1}},
			{{0, 1}, {1, 1}, {1, 0}},
		}
	case TetrominoJ:
		return [][][]int{
			{{1, 0, 0}, {1, 1, 1}},
			{{1, 1}, {1, 0}, {1, 0}},
			{{1, 1, 1}, {0, 0, 1}},
			{{0, 1}, {0, 1}, {1, 1}},
		}
	case TetrominoL:
		return [][][]int{
			{{0, 0, 1}, {1, 1, 1}},
			{{1, 0}, {1, 0}, {1, 1}},
			{{1, 1, 1}, {1, 0, 0}},
			{{1, 1}, {0, 1}, {0, 1}},
		}
	}
	return nil
}

// GetColor returns the color for a tetromino type
func GetColor(t TetrominoType) Color {
	switch t {
	case TetrominoI:
		return ColorCyan
	case TetrominoO:
		return ColorYellow
	case TetrominoT:
		return ColorPurple
	case TetrominoS:
		return ColorGreen
	case TetrominoZ:
		return ColorRed
	case TetrominoJ:
		return ColorBlue
	case TetrominoL:
		return ColorOrange
	}
	return ColorNone
}

// NewTetromino creates a new tetromino of the given type
func NewTetromino(t TetrominoType) *Tetromino {
	shapes := GetShapes(t)
	return &Tetromino{
		Type:     t,
		Shape:    shapes[0],
		Color:    GetColor(t),
		X:        3,
		Y:        0,
		Rotation: 0,
	}
}

// Rotate rotates the tetromino clockwise
func (t *Tetromino) Rotate() {
	shapes := GetShapes(t.Type)
	t.Rotation = (t.Rotation + 1) % len(shapes)
	t.Shape = shapes[t.Rotation]
}

// RotateBack rotates the tetromino counter-clockwise
func (t *Tetromino) RotateBack() {
	shapes := GetShapes(t.Type)
	t.Rotation = (t.Rotation - 1 + len(shapes)) % len(shapes)
	t.Shape = shapes[t.Rotation]
}
