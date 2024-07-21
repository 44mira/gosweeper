package main

import (
	"github.com/gdamore/tcell/v2"
)

const ZERO = 48

func DrawTile(s tcell.Screen, x, y int, style tcell.Style, tile Tile, neighbors uint) {
	switch {
	case tile.IsFlagged:
		s.SetContent(x, y, '󰈿', nil, style.Foreground(tcell.ColorIndianRed))
	case tile.IsClose:
		s.SetContent(x, y, '󰞍', nil, style.Foreground(tcell.ColorDarkSeaGreen))
	case tile.IsMine:
		s.SetContent(x, y, '󰷚', nil, style.Foreground(tcell.ColorSteelBlue))
	case neighbors == 0:
		s.SetContent(x, y, '∿', nil, style.Foreground(tcell.ColorLightSkyBlue))
	default:
		s.SetContent(x, y, rune(neighbors)+ZERO, nil, style.Foreground(numberColor(neighbors)))
	}

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
