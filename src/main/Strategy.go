package main

import "fmt"

type Strategy struct {
	cooperate float32
	defect float32
	display string
}


/* Create Strategy list for player to choose */
func getStrategyList() []Strategy {

	strategies := make([]Strategy, 5)


	strategies[0] = Strategy{1, 0, "Fully trust on opponent and remain silent!"}
	fmt.Println("1. Fully trust on opponent and remain silent!")

	strategies[1] = Strategy{0.7, 0.3, "Mostly trust opponent, but sometimes defect!"}
	fmt.Println("2. Mostly trust opponent, but sometimes defect!")

	strategies[2] = Strategy{0.5, 0.5, "Half trust opponent, and randomly pick its decision!"}
	fmt.Println("3. Half trust opponent, and randomly pick its decision!")

	strategies[3] = Strategy{0.3, 0.7, "Mostly defect opponent, but sometimes remain silent"}
	fmt.Println("4. Mostly defect opponent, but sometimes remain silent")

	strategies[4] = Strategy{0, 1, "Doesn't trust opponent and always back door defect!"}
	fmt.Println("5. Doesn't trust opponent and always back door defect!")

	return strategies
}