package main

import "math/rand"

type Player struct {
	name string
	strategy Strategy
}


/* Return a decision for player */
func makeDecision(player Player) int {
	/* Get Player's Strategy Probability */
	numCo := int(player.strategy.cooperate * 10)
	numDe := int(player.strategy.defect * 10)

	/* Create a Slice with certain number of choices for outcomes */

	/* Number of Co-operate choice */
	decisions := make([]int, numCo+numDe)
	for i := 0; i < numCo; i++ {
		decisions = append(decisions, 0)
	}

	/* Number of Defect choice */
	for i := 0; i < numDe; i++ {
		decisions = append(decisions, 1)
	}

	/* Return a randomly picked Decision */
	return decisions[rand.Intn(len(decisions))]
}
