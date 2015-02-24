package fireload

import (
	"container/ring"
	"errors"
	"math/rand"
)

// Pool represents a pool of Firebase Namespaces, with an associated Strategy
// for load balancing in the pool.
type Pool struct {
	Nodes *ring.Ring
	Strategy
}

// NewPool returns a Pool with the given namespaces.
func NewPool(nodes ...Namespace) (*Pool, error) {
	r := ring.New(len(nodes))
	for _, node := range nodes {
		r.Value = node
		r = r.Next()
	}

	return &Pool{Nodes: r}, nil
}

// ErrInvalidStrategy is returned by `Pool.SetStrategy` when the given strategy is invalid.
var ErrInvalidStrategy = errors.New("Invalid strategy")

// SetStrategy sets the strategy of the pool to the given strategy.
func (p *Pool) SetStrategy(strategy Strategy) error {
	switch strategy {
	case StrategyRandom, StrategyRoundRobin:
		p.Strategy = strategy
	default:
		return ErrInvalidStrategy
	}

	return nil
}

// Add inserts the given namespace into the pool.
//
// The Nodes in a Pool are considered to be unordered, so the Namespace is inserted at the
// "next" position (i.e. a subsequent call to NextRoundRobin would return the inserted Namespace).
func (p *Pool) Add(ns Namespace) {
	s := &ring.Ring{Value: ns}
	r := p.Nodes.Link(s)
	s.Link(r)
}

// Drop removes all Namespaces with the given domain from the pool.
func (p *Pool) Drop(domain string) error {
	for i := 0; i < p.Nodes.Len(); i++ {
		ns, ok := p.Nodes.Value.(Namespace)
		if !ok {
			return errors.New("Could not typecast Ring.Value to Namespace")
		}

		if ns.Domain == domain {
			prev := p.Nodes.Prev()
			next := p.Nodes.Next()
			prev.Link(next)
			p.Nodes = nil
			p.Nodes = prev
		}

		p.Nodes = p.Nodes.Next()
	}
	return nil
}

// Next returns from the pool the Namespace deemed to be "next" according to the
// pool's strategy.
func (p *Pool) Next() Namespace {
	switch p.Strategy {
	case StrategyRoundRobin:
		return p.NextRoundRobin()
	default:
		return p.NextRandom()
	}
}

// NextRandom returns a random Namespace from the pool
func (p *Pool) NextRandom() Namespace {
	n := rand.Intn(p.Nodes.Len())
	p.Nodes = p.Nodes.Move(n)
	return p.Nodes.Value.(Namespace)
}

// NextRoundRobin returns a Namespace in "round-robin" fashion
func (p *Pool) NextRoundRobin() Namespace {
	p.Nodes = p.Nodes.Next()
	return p.Nodes.Value.(Namespace)
}
