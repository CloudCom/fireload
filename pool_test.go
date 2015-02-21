package fireload

import "testing"

func testPool() *Pool {
	nodes := []*Namespace{
		{Domain: "node-1"},
		{Domain: "node-2"},
		{Domain: "node-3"},
	}
	p, _ := NewPool(nodes)
	return p
}

func Test_NewPool(t *testing.T) {
	p := testPool()

	if p == nil {
		t.Fail()
	}
}

func Test_NewPool_NilPointer(t *testing.T) {
	nodes := []*Namespace{
		{Domain: "node-1"},
		{Domain: "node-2"},
		{Domain: "node-3"},
		nil,
	}

	if _, err := NewPool(nodes); err != ErrNilPointer {
		t.Fatalf("Expected NewPool with nil pointer to return %s but got %v", ErrNilPointer, err)
	}
}

func Test_Pool_DefaultStrategy(t *testing.T) {
	p := testPool()

	if p.Strategy != StrategyRandom {
		t.Fatalf("Expected default strategy to be %d but got %d", StrategyRandom, p.Strategy)
	}
}

func Test_Pool_SetStrategy(t *testing.T) {
	p := testPool()
	if err := p.SetStrategy(StrategyRoundRobin); err != nil {
		t.Fatal(err)
	}

	if p.Strategy != StrategyRoundRobin {
		t.Fatalf("Expected default strategy to be %d but got %d", StrategyRoundRobin, p.Strategy)
	}
}

func Test_Pool_SetStrategy_Invalid(t *testing.T) {
	p := testPool()

	if err := p.SetStrategy(Strategy(-1)); err != ErrInvalidStrategy {
		t.Fatalf("Expected SetStrategy with invalid strategy to return %s but got %v", ErrNilPointer, err)
	}
}

func Test_Pool_NextRandom(t *testing.T) {
	p := testPool()

	for i := 0; i <= 3*len(p.Nodes); i++ {
		if got := p.NextRandom(); got == nil {
			t.Fatalf("Expected NextRandom() not to yield a nil pointer")
		}
	}
}

func Test_Pool_Next_StrategyRandom(t *testing.T) {
	p := testPool()
	if err := p.SetStrategy(StrategyRandom); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 3*len(p.Nodes); i++ {
		if got := p.Next(); got == nil {
			t.Fatalf("Expected Next() not to yield a nil pointer")
		}
	}
}

func Test_Pool_NextRoundRobin(t *testing.T) {
	p := testPool()

	for i := 0; i <= 3; i++ {
		for j, expected := range p.Nodes {
			if got := p.NextRoundRobin(); got != expected {
				t.Fatalf("Expected NextRoundRobin() to yield %v but got %v", expected, got)
			}

			if p.lastIndex != j {
				t.Fatalf("Expected NextRoundRobin() to update lastIndex to %d, but got %d", j, p.lastIndex)
			}
		}
	}
}

func Test_Pool_Next_StrategyRoundRobin(t *testing.T) {
	p := testPool()
	if err := p.SetStrategy(StrategyRoundRobin); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 3; i++ {
		for j, expected := range p.Nodes {
			if got := p.Next(); got != expected {
				t.Fatalf("Expected Next() to yield %v but got %v", expected, got)
			}

			if p.lastIndex != j {
				t.Fatalf("Expected Next() to update lastIndex to %d, but got %d", j, p.lastIndex)
			}
		}
	}
}
