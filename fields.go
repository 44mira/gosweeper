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

