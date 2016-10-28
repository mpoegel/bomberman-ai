package strategy

import "github.com/mpoegel/bomberman-ai/bomberman"

// SuicideStrategy is strategy that immediately suicides
type SuicideStrategy struct {
	Name  string
	count int
}

// Execute calculates the next move given this strategy
func (strategy *SuicideStrategy) Execute(msg *bomberman.Message) string {
	var move string
	if strategy.count == 0 {
		move = "b"
	} else {
		move = ""
	}
	strategy.count++
	return move
}

// NewSuicideStrategy creates a new SuicideStrategy
func NewSuicideStrategy() *SuicideStrategy {
	s := SuicideStrategy{"suicide", 0}
	return &s
}
