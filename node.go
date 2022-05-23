package tinyman

import (
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"go.uber.org/ratelimit"
)

// Node ...
type Node struct {
	ratelimit.Limiter

	ac *algod.Client
	ai *indexer.Client
}

// NewNode ...
func NewNode(ac *algod.Client, ai *indexer.Client) *Node {
	node := &Node{}
	node.Limiter = ratelimit.New(50)
	node.ac = ac
	node.ai = ai

	return node
}
