
# GOSWEEPER

Play minesweeper straight from your command line!

## Dependencies

- A [*Nerd Font*](https://www.nerdfonts.com/)
- Not a strict dependency, but only tested on the [`kitty`](https://sw.kovidgoyal.net/kitty/) terminal
    - Might have formatting issues on other terminal emulators

## TODO

- [x] Generate field
    - [x] Handle invalid field parameters
    - [x] Distribute mines pseudo-randomly
- [x] Display field 
    - [x] Use Nerdfont Icons
    - [x] Color tiles
- [x] Number tiles based on neighbor mines
- [ ] Flagging
- [ ] Command line arguments
    - [ ] Board dimensions
    - [ ] Theme
        - [ ] Colors
        - [ ] Icons
    - [ ] Mines
    - [ ] Timer
        - [ ] Stopwatch
        - [ ] Time trial
- [ ] Controls
    - [ ] Arrow keys
    - [ ] HJKL
    - [ ] Flag (space)
    - [ ] Dig (enter)
- [ ] Timer
    - [ ] Stopwatch
    - [ ] Time trial
- [ ] Score
    - [ ] based on time
    - [ ] based on mine count

## Optional features

- [ ] Better generation algo (don't know yet)
- [ ] Mouse controls (Tcell)
