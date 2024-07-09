package main

import (
	"errors"
	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
)

type Tile struct {
	IsMine    bool
	IsClose   bool
	IsFlagged bool
}

// A Field is represented as a 2d slice
type Field struct {
	Tiles     [][]Tile
	AdjMatrix [][]uint

	// game state
	TotalMines uint
	LiveMines  uint
}

// Display is a function that takes a Field struct, and displays the current
// state of the board based on the individual Tiles.
func (f *Field) Display(s tcell.Screen) {
	tileStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	tiles, adjMatrix := f.Tiles, f.AdjMatrix

	for i := range tiles {
		for j := range tiles[i] {
			x, y := i*2, j
			DrawTile(s, x, y, tileStyle, tiles[i][j], adjMatrix[i][j])
		}
	}
}

// Initialize takes the x * y dimensions of a field, followed by its number of
// mines. Mines are pseudo-randomly distributed onto the field after the
// initialization.
func Initialize(x, y, mines int) (Field, error) {

	if x*y <= 0 {
		return Field{}, errors.New("Invalid field dimensions.")
	}

	if mines > x*y {
		return Field{}, errors.New("Invalid mine count.")
	}

	mineCount := mines

	// Create matrix
	tiles := make([][]Tile, x) // Create columns
	for i := range tiles {
		tiles[i] = make([]Tile, y) // Create rows

		for j := range tiles[i] {
			tiles[i][j].IsClose = true
		}
	}

	for mines > 0 {
		mineX := rand.IntN(x)
		mineY := rand.IntN(y)

		if tiles[mineX][mineY].IsMine {
			continue
		}

		// convert the int into a coordinate
		tiles[mineX][mineY].IsMine = true

		mines--
	}
	field := Field{tiles, [][]uint{}, uint(mineCount), uint(mineCount)}
	field.AdjMatrix = field.GetAdjacencyMatrix().Cells

	return field, nil
}
