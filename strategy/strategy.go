package strategy

import "github.com/mpoegel/bomberman-ai/bomberman"

// Strategy defines approach for choosing the next move
type Strategy interface {
	Execute(*bomberman.Message) string
}
