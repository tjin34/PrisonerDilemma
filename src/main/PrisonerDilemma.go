package main

import (
	"fmt"
	"log"
)

type PDSimulation struct {
	ID int
	player1 Player
	player2 Player
	sentence Sentence
}

type Sentence struct {
	player1 int
	player2 int
}


func play(ID int, player1 Player, player2 Player) (bool, PDSimulation) {
	/* Get player's judgement based on their decision */
	sentence := judgement(makeDecision(player1), makeDecision(player2))

	/* Return the result of current play */
	return true, PDSimulation{ID, player1, player2,sentence}
}


/* Return the judgement for both players */
func judgement(decision1 int, decision2 int) Sentence {

	/* If player1 defects */
	if decision1 == 1 {
		/* And player2 defects, too */
		if decision2 == 1 {
			return Sentence{2,2}
		}
		/* But if player2 remains silent */
		return Sentence{0, 3}
	}

	/* Else if player 1 remains silent */
	/* But Player 2 defect */
	if decision2 == 1 {
		return Sentence{3, 0}
	}

	/* Both of them remain silent */
	return Sentence{1, 1}
}


/* Run a single simulation */
func simulation(statistic Statistic) Statistic{

	/* Channel for all finished simulations */
	simulationChannel := make(chan int, statistic.numberRuns)

	/* Channel for idle slots for future simulation */
	idleChannel := make(chan int)

	/* For test sake, only 100 simulation at the same time */
	for i := 0; i < 100 ; i++  {
		go func(i int) {
			idleChannel <- i
		}(i)
	}

	/* Loop through num, and execute simulations */
	for i := 0; i < statistic.numberRuns; i++ {
		go func(i int) {
			pass,result := play(i, statistic.player1, statistic.player2)
			channel := <- idleChannel

			/* If simulation finished as true, record stats and release slots for future simulation */
			if pass {
				statistic.sentence1 += result.sentence.player1
				statistic.sentence2 += result.sentence.player2
				simulationChannel <- i
				idleChannel <- channel
			}
		}(i)
	}

	/* Wait for all simulation to be finished */
	for i := 0; i < statistic.numberRuns; i++ {
		<- simulationChannel
	}

	return statistic
}

/* Manually run a simulation */
func manualSimulation() Statistic {
	fmt.Println()
	fmt.Println("Firstly, please choose a strategy for player1:")
	strategies := getStrategyList()
	fmt.Print("Enter your choice of strategy for player 1: ")
	var strategy1 int
	_, err := fmt.Scanf("%d", &strategy1)
	if err != nil {
		log.Fatal(err)
	}
	strategy1 -= 1

	fmt.Println()
	fmt.Println("Secondly, please choose a strategy for player2:")
	strategies = getStrategyList()
	fmt.Print("Enter your choice of strategy for player 2: ")
	var strategy2 int
	_, err = fmt.Scanf("%d", &strategy2)
	if err != nil {
		log.Fatal(err)
	}
	strategy2 -= 1


	fmt.Println()
	fmt.Println("Finally, how many trials you wan to run?")
	fmt.Print("Enter number of trials: ")
	var num int
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Print("Simulation in progress...")

	player1 := Player{"player1", strategies[strategy1]}
	player2 := Player{"player2", strategies[strategy2]}
	statistic := Statistic{num, player1,0, player2, 0}

	statistic = simulation(statistic)

	fmt.Println("Simulation done!")

	return statistic
}


/* Simulate all possibilities and compare the results */
func autoSimulation() Analysis{

	analysis := Analysis{}

	analysis.sentences = make(map[int][]Statistic)

	strategies := getStrategyList()

	analysis.strategies = strategies

	fmt.Println()
	fmt.Print("Simuation in progress...")

	for i := 0; i < len(strategies); i++  {
		for j := 0; j < len(strategies); j++  {
				player1 := Player{"player1", strategies[i]}
				player2 := Player{"player2", strategies[j]}
				statistic := Statistic{1000, player1,0, player2, 0}
				statistic = simulation(statistic)
				analysis.sentences[i] = append(analysis.sentences[i], statistic)
		}
	}

	fmt.Println("Simulation done!")

	return analysis
}