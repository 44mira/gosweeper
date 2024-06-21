![image](https://github.com/44mira/gosweeper/assets/116419708/cddb945b-6757-4a4c-91c0-94a877775e4e)

Play minesweeper straight from your command line!

## ✴️ Dependencies

- [*Go*](https://go.dev)
- A [*Nerd Font*](https://www.nerdfonts.com/)
- Not a strict dependency, but only tested on the [`kitty`](https://sw.kovidgoyal.net/kitty/) terminal
    - Might have formatting issues on other terminal emulators

## ✴️ TODO

- [x] Generate field
    - [x] Handle invalid field parameters
    - [x] Distribute mines pseudo-randomly
- [x] Display field 
    - [x] Use Nerdfont Icons
    - [x] Color tiles
    - [x] Use a CLI Framework (TCell)
        - [x] Roughly center the board on screen
        - [ ] Recenter on resize
- [x] Number tiles based on neighbor mines
- [x] Flagging
- [ ] Command line arguments
    - [x] Board dimensions
    - [ ] Theme
        - [ ] Colors
        - [ ] Icons
    - [x] Mines
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

## ✴️ Optional features

- [ ] Better generation algo (don't know yet)
- [ ] Mouse controls (Tcell)

## ✴️ Build instructions

1. Clone this repo

```bash
git clone https://github.com/44mira/gosweeper.git

cd gosweeper
```

2. Build the binary

```bash
go build gosweeper
```

3. Run the game and pass command flags

```bash
# macos / linux
./gosweeper -mine=10 -x=5 -y=5      # start a game 5x5 game with 10 mines

# windows
gosweeper.exe
```

| Flag   | Usage               |
| ------ | ------------------- |
| mine   | Number of mines     |
| x      | Width of the field  |
| y      | Height of the field |
