package main

import (
	"flag"
	"github.com/gdamore/tcell/v2"
	"log"
)

func main() {
	mines := flag.Int("mine", 10, "Number of mines")
	y := flag.Int("y", 5, "Height of the field")
	x := flag.Int("x", 5, "Width of the field")
	flag.Parse()

	GameLoop(*x, *y, *mines)
}

// [[ Game loop ]] {{{
func GameLoop(x, y, mines int) {
	// [[ Initial Boilerplate ]] {{{1
	// Initialize Field
	game, err := Initialize(x, y, mines)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Initialize Screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.Clear()

	// catch panic and rethrow them after calling s.Fini()
	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	game.Display(s)
	// }}}1

	for {
		s.Show()
		game.Display(s)
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			}

		case *tcell.EventMouse:
			x, y := ev.Position()

			switch ev.Buttons() {
			case tcell.Button1: // Primary click
				game.Dig(x/2, y)
			case tcell.Button2: // Secondary click
				game.Flag(x/2, y)
			}
		}
	}
}

// }}}

// [[ Dig and Flag ]] {{{
// Digs a tile on (x, y) given that it is closed and not flagged
// Recursively digs surrounding tiles whenever the current tile has 0 neighbor
// mines
func (f *Field) Dig(x, y int) {
	// Check bounds
	if x >= len(f.Tiles) || x < 0 || y >= len(f.Tiles[0]) || y < 0 || !f.Tiles[x][y].IsClose {
		return
	}

	// Can't open flagged
	if !f.Tiles[x][y].IsFlagged {
		f.Tiles[x][y].IsClose = false
	}

	// Recursively expand the dig
	if f.AdjMatrix[x][y] == 0 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ { // nested loop to dig surrounding
				// skip recursion on current
				if i == 0 && j == 0 {
					continue
				}

				f.Dig(x+i, y+j)
			}
		}
	}
}

func (f *Field) Flag(x, y int) {
	// Check bounds
	if x >= len(f.Tiles) || x < 0 || y >= len(f.Tiles[0]) || y < 0 {
		return
	}

	// Can't flag opened
	if f.Tiles[x][y].IsClose {
		f.Tiles[x][y].IsFlagged = !f.Tiles[x][y].IsFlagged
	}
}
