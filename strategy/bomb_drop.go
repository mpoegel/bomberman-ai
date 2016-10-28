package strategy

import "github.com/mpoegel/bomberman-ai/bomberman"

// BombDropStrategy is a straightforward strategy to place bombs next to soft walls, move out of the
// blast zone, and repeat until ?
type BombDropStrategy struct {
	Name      string
	moveCount int
}

// Execute calculates the next move given this strategy
func (strategy *BombDropStrategy) Execute(msg *bomberman.Message) string {
	var move string
	if strategy.moveCount == 0 {
		move = "b"
	} else {
		move = ""
	}
	strategy.moveCount++
	return move
}

// NewBombDropStrategy creates a new BombDropStrategy
func NewBombDropStrategy() *BombDropStrategy {
	s := BombDropStrategy{"Bomb-Drop", 0}
	return &s
}
