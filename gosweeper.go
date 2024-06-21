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

	// Initialize Field
	game, err := Initialize(*y, *x, *mines)
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
				game.Dig(x, y)
			case tcell.Button2: // Secondary click
				game.Flag(x, y)
			}
		}
	}
}

func (f *Field) Dig(x, y int) {
	// Check bounds
	if x >= len(f.Tiles)*2 && x < 0 && y >= len(f.Tiles[0]) && y < 0 {
		return
	}

	// Can't open flagged
	if !f.Tiles[x/2][y].IsFlagged {
		f.Tiles[x/2][y].IsClose = false
	}
}

func (f *Field) Flag(x, y int) {
	// Check bounds
	if x >= len(f.Tiles)*2 && x < 0 && y >= len(f.Tiles[0]) && y < 0 {
		return
	}

	// Can't flag opened
	if f.Tiles[x/2][y].IsClose {
		f.Tiles[x/2][y].IsFlagged = !f.Tiles[x/2][y].IsFlagged
	}
}
