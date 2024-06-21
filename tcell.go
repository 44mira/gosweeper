package main

import (
	"github.com/gdamore/tcell/v2"
)

const ZERO = 48

func DrawTile(s tcell.Screen, x, y int, style tcell.Style, tile Tile, neighbors uint, dimensions [2]int) {

	xmax, ymax := s.Size()

	x += xmax/2 - dimensions[0]/2*3 - 1 // offset by width and last space
	y += ymax/2 - dimensions[1]/2

	switch {
	case tile.IsClose:
		s.SetContent(x, y, ' ', []rune{' '}, style.Foreground(tcell.ColorPaleGreen))
	case tile.IsFlagged:
		s.SetContent(x, y, ' ', []rune{''}, style.Foreground(tcell.ColorDarkMagenta))
	case tile.IsMine:
		s.SetContent(x, y, ' ', []rune{'󰷚'}, style.Foreground(tcell.ColorSteelBlue))
	case neighbors == 0:
		s.SetContent(x, y, ' ', []rune{'󱁐'}, style.Foreground(tcell.ColorLightSkyBlue))
	default:
		s.SetContent(x, y, ' ', []rune{rune(neighbors) + ZERO}, style.Foreground(numberColor(neighbors)))
	}
	s.SetContent(x+1, y, ' ', nil, style.Foreground(numberColor(neighbors)))

}

func numberColor(n uint) tcell.Color {
	switch n {
	case 0:
		return tcell.ColorLightBlue
	case 1:
		return tcell.ColorCornflowerBlue
	case 2:
		return tcell.ColorGreen
	case 3:
		return tcell.ColorRed
	case 4:
		return tcell.ColorMediumPurple
	case 5:
		return tcell.ColorOrangeRed
	case 6:
		return tcell.ColorDarkCyan
	case 7:
		return tcell.ColorLightGray
	case 8:
		return tcell.ColorCornsilk
	}

	return tcell.ColorFireBrick
}
