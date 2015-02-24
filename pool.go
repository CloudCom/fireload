package fireload

import (
	"container/ring"
	"errors"
	"math/rand"
)

// Pool .
type Pool struct {
	Nodes *ring.Ring
	Strategy
}

// NewPool .
func NewPool(nodes ...Namespace) (*Pool, error) {
	r := ring.New(len(nodes))
	for _, node := range nodes {
		r.Value = node
		r = r.Next()
	}

	return &Pool{Nodes: r}, nil
}

// ErrInvalidStrategy .
var ErrInvalidStrategy = errors.New("Invalid strategy")

// SetStrategy .
func (p *Pool) SetStrategy(strategy Strategy) error {
	switch strategy {
	case StrategyRandom, StrategyRoundRobin:
		p.Strategy = strategy
	default:
		return ErrInvalidStrategy
	}

	return nil
}

// Next .
func (p *Pool) Next() Namespace {
	switch p.Strategy {
	case StrategyRoundRobin:
		return p.NextRoundRobin()
	default:
		return p.NextRandom()
	}
}

// NextRandom .
func (p *Pool) NextRandom() Namespace {
	n := rand.Intn(p.Nodes.Len())
	p.Nodes = p.Nodes.Move(n)
	return p.Nodes.Value.(Namespace)
}

// NextRoundRobin .
func (p *Pool) NextRoundRobin() Namespace {
	p.Nodes = p.Nodes.Next()
	return p.Nodes.Value.(Namespace)
}
