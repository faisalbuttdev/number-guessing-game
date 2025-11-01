
# Number Guessing Game

A simple **CLI-based Number Guessing Game** written in **Go**, featuring multiple difficulty levels, a live timer, and highscore tracking.

This project is designed to demonstrate Go concepts like **structs, maps, slices, goroutines, and channels** in a fun, interactive way.

---

## Features

* Choose between **3 difficulty levels**:

  * Easy (10 guesses)
  * Medium (5 guesses)
  * Hard (3 guesses)

* **Live timer** shows how long you’ve been guessing.

* **Highscore tracking** (in-memory) to keep track of your best scores.

* Replay the game as many times as you like.

---

## Installation

1. Make sure you have **Go** installed (version 1.20+ recommended).
2. Clone the repository:

```bash
git clone https://github.com/yourusername/number-guessing-game.git
cd number-guessing-game
```

3. Run the game:

```bash
go run main.go
```

---

## How to Play

1. Run the game using `go run main.go`.
2. Select a difficulty level by typing `1`, `2`, or `3`.
3. Guess the number between 1 and 100 within the allowed number of tries.
4. Watch the **live timer** while playing.
5. If you guess correctly, your **score (number of tries)** will be added to the highscore list.
6. Decide whether to play again after each round.


## Technologies

* **Go**
* CLI-based user interface
* Goroutines & Channels for live timer

---

## Future Enhancements

* Save **highscores to a file** for persistent tracking
* Add **hints** or range narrowing after wrong guesses
* Color-coded output for wins/losses

---


