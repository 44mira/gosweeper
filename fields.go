package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Tile struct {
	IsMine    bool
	IsClose   bool
	IsFlagged bool
}

// A Field is represented as a 2d slice
type Field struct {
	Tiles      [][]Tile
	TotalMines uint
	LiveMines  uint
}

// Display is a function that takes a Field struct, and displays the current
// state of the board based on the individual Tiles.
func (f *Field) Display() {
	tiles := f.Tiles
	adjMatrix := f.GetAdjacencyMatrix().Cells

	for i := range tiles {
		for j := range tiles[i] {
			currentTile := tiles[i][j]

			if currentTile.IsClose {
				fmt.Print(" ⬛")
				continue
			}

			if currentTile.IsFlagged {
				fmt.Print(BrightMagenta + " " + Reset)
				continue
			}

			if currentTile.IsMine {
				fmt.Print(BrightCyan + " 󰷚 " + Reset)
				continue
			}

			if adjMatrix[i][j] == 0 {
				fmt.Print(" 󱁐 ") // empty
				continue
			}

			fmt.Printf("  %v%v"+Reset, numberColor(adjMatrix[i][j]), adjMatrix[i][j])
		}
		fmt.Println()
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
	}

	// Insert mines
	insertedMines := make(map[int]bool) // hashmap for inserted mines

	for mines > 0 {
		mine := rand.IntN(x * y)

		if insertedMines[mine] {
			continue
		}
		insertedMines[mine] = true // put value in hashmap

		// convert the int into a coordinate
		tiles[mine/x][mine%y].IsMine = true

		mines--
	}

	return Field{tiles, uint(mineCount), uint(mineCount)}, nil
}

func numberColor(n uint) string {
	switch n {
	case 0:
		return BrightCyan
	case 1:
		return BrightBlue
	case 2:
		return Green
	case 3:
		return BrightRed
	case 4:
		return Blue
	case 5:
		return Red
	case 6:
		return Cyan
	case 7:
		return Magenta
	case 8:
		return Gray
	}

	return "ERROR"
}
