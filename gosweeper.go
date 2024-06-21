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

	game, err := Initialize(*y, *x, *mines)

	if err != nil {
		log.Fatalf("%+v", err)
	}

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	game.Display(s)

	// catch panic and rethrow them after calling s.Fini()
	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	for {
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			}
		}
	}
}
