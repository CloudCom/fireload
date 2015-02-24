package fireload

import (
	"container/ring"
	"testing"

	"github.com/kr/pretty"
)

var testNodes = []Namespace{
	{Domain: "node-1.firebaseio.com"},
	{Domain: "node-2.firebaseio.com"},
	{Domain: "node-3.firebaseio.com"},
	{Domain: "node-4.firebaseio.com"},
	{Domain: "node-5.firebaseio.com"},
}

func Test_NewPool(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	assertSeenSet(t, p.Nodes, testNodes)
}

func assertSeenSet(t *testing.T, nodes *ring.Ring, expected []Namespace) {
	seen := map[string]int{}
	for i := 0; i < nodes.Len(); i++ {
		val := nodes.Value
		ns, ok := val.(Namespace)
		if !ok {
			t.Fatalf("Expected nodes to be a Ring with Namespace values. Got %T", val)
		}
		seen[ns.Domain]++
		nodes = nodes.Next()
	}

	for _, ns := range expected {
		if _, ok := seen[ns.Domain]; !ok {
			t.Fatalf("Expected nodes to contain %s.\n\n%# v", ns.Domain, pretty.Formatter(seen))
		}

		seen[ns.Domain]--
	}

	for domain, count := range seen {
		if count != 0 {
			t.Fatalf("Expected count for %s off by %d", domain, count)
		}
	}
}

func Test_Pool_DefaultStrategy(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	if p.Strategy != StrategyRandom {
		t.Fatalf("Expected default strategy to be %d but got %d", StrategyRandom, p.Strategy)
	}
}

func Test_Pool_SetStrategy(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.SetStrategy(StrategyRoundRobin); err != nil {
		t.Fatal(err)
	}

	if p.Strategy != StrategyRoundRobin {
		t.Fatalf("Expected default strategy to be %d but got %d", StrategyRoundRobin, p.Strategy)
	}
}

func Test_Pool_SetStrategy_Invalid(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.SetStrategy(Strategy(-1)); err != ErrInvalidStrategy {
		t.Fatalf("Expected SetStrategy with invalid strategy to return %s but got %v", ErrInvalidStrategy, err)
	}
}

func Test_Pool_Add_Pass(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	newNS := NewNamespace("node-6.firebaseio.com")
	p.Add(newNS)

	assertSeenSet(t, p.Nodes, append(testNodes, newNS))
}

func Test_Pool_NextRandom(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 2*p.Nodes.Len(); i++ {
		next := p.NextRandom()
		if next.Domain == "" {
			t.Fatalf("Expected NextRandom() not to yield nil value Namespace")
		}
	}
}

func Test_Pool_Next_StrategyRandom(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.SetStrategy(StrategyRandom); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 2*p.Nodes.Len(); i++ {
		if got := p.Next(); got.Domain == "" {
			t.Fatalf("Expected Next() not to yield nil value Namespace")
		}
	}
}

func Test_Pool_NextRoundRobin(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 2*p.Nodes.Len(); i++ {
		node := p.Nodes.Next()

		expected, ok := node.Value.(Namespace)
		if !ok {
			t.Fatalf("Couldn't typecast %v as Namesapce", node.Value)
		}

		if got := p.NextRoundRobin(); got.String() != expected.String() {
			t.Fatalf("Expected NextRoundRobin() to yield %v but got %v", expected, got)
		}

	}
}

func Test_Pool_Next_StrategyRoundRobin(t *testing.T) {
	p, err := NewPool(testNodes...)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.SetStrategy(StrategyRoundRobin); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 2*p.Nodes.Len(); i++ {
		node := p.Nodes.Next()

		expected, ok := node.Value.(Namespace)
		if !ok {
			t.Fatalf("Couldn't typecast %v as Namesapce", node.Value)
		}

		if got := p.Next(); got.String() != expected.String() {
			t.Fatalf("Expected Next() to yield %v but got %v", expected, got)
		}
	}
}
