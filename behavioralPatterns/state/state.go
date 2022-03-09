package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
STATE: Directly related to FSM.

OBJECTIVE:
- Develop FSM
- Have a type that alters its own behavior when some internal things have changed.
- Model compĺex graphs and pipelines can be upgraded easily by adding more states and rerouting their output states.

EXAMPLE: Guess number name,
- Number guessing game where we'll have to guess some number between 0 and 10 with few attempts.


ACCEPTANCE CRITERIA:
- Game will ask the player how many tries they will have before losing the game.
- The number to guess must be between 0 and 10.
- Every time a player enters a number to guess, the number of retries drop by one.
- If the number of retries reaches zero and the number is still incorrect, the game finishes and the players has lost.
- If the player guesses the number, the player wins.
*/
type GameState interface {
	ExecuteState(*GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

type WinState struct {
}

func (s *WinState) ExecuteState(c *GameContext) bool {
	println("Congrats, you won")
	return false
}

type LoseState struct {
}

func (s *LoseState) ExecuteState(c *GameContext) bool {
	fmt.Printf("You lose. The correct number was: %d\n", c.SecretNumber)
	return false
}

type StartState struct{}

func (s *StartState) ExecuteState(c *GameContext) bool {
	c.Next = &AskState{}
	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number of retries to set difficulty: ")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)
	return true
}

type AskState struct{}

func (a *AskState) ExecuteState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)
	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1
	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true
}

type FinishState struct{}

func (a *FinishState) ExecuteState(c *GameContext) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}
	return true
}

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	for game.Next.ExecuteState(&game) {

	}
}
