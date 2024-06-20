package main

import (
	"fmt"
	"math/rand/v2"
)

type Tile struct {
	IsMine  bool
	IsClose bool
}

// A Field is represented as a 2d slice
type Field struct {
	Tiles [][]Tile
}

// Display is a function that takes a Field struct, and displays the current
// state of the board based on the individual Tiles.
func (f *Field) Display() {
	tiles := f.Tiles

	for i := range tiles {
		for j := range tiles[i] {
			currentTile := tiles[i][j]

			if currentTile.IsClose {
				fmt.Print(" ⬛")
				continue
			}

			if currentTile.IsMine {
				fmt.Print(" 󰷚 ")
				continue
			}

			// TODO: logic for neighbor mines

			fmt.Print(" ⬜") // empty
		}
		fmt.Println()
	}
}

// Initialize takes the x * y dimensions of a field, followed by its number of
// mines. Mines are pseudo-randomly distributed onto the field after the
// initialization.
func Initialize(x, y, mines int) Field {

	// Create matrix
	tiles := make([][]Tile, x) // Create columns
	for i := range tiles {
		tiles[i] = make([]Tile, y) // Create rows
	}

	// Insert mines
	insertedMines := make(map[int]bool, x*y) // hashmap for inserted mines

	for mines > 0 {
		mine := rand.IntN(x * y)

		if insertedMines[mine] {
			continue
		}

		// convert the int into a coordinate
		tiles[mine/x][mine%x].IsMine = true
		mines--
	}

	return Field{tiles}
}
