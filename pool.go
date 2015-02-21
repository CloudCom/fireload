package fireload

import (
	"errors"
	"math/rand"
	"time"
)

// Pool .
type Pool struct {
	Nodes     []*Namespace
	lastIndex int
	Strategy
	source rand.Source
}

// ErrNilPointer .
var ErrNilPointer = errors.New("Nodes slice cannot contain a nil pointer")

// NewPool .
func NewPool(nodes []*Namespace) (*Pool, error) {
	for _, ptr := range nodes {
		if ptr == nil {
			return nil, ErrNilPointer
		}
	}
	return &Pool{
		Nodes:     nodes,
		source:    rand.NewSource(time.Now().UnixNano()),
		lastIndex: -1,
	}, nil
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
func (p *Pool) Next() *Namespace {
	switch p.Strategy {
	case StrategyRoundRobin:
		return p.NextRoundRobin()
	default:
		return p.NextRandom()
	}
}

// NextRandom .
func (p *Pool) NextRandom() *Namespace {
	p.lastIndex = int(p.source.Int63() % int64(len(p.Nodes)))
	return p.Nodes[p.lastIndex]
}

// NextRoundRobin .
func (p *Pool) NextRoundRobin() *Namespace {
	p.lastIndex = (p.lastIndex + 1) % len(p.Nodes)
	return p.Nodes[p.lastIndex]
}
