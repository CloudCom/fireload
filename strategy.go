package fireload

// Strategy represents various strategies for load balancing a pool of Firebase Namespaces
type Strategy int

const (
	// StrategyRandom picks a Namespace at random
	StrategyRandom Strategy = iota

	// StrategyRoundRobin returns Namespaces in "round-robin" fashion
	StrategyRoundRobin
)
