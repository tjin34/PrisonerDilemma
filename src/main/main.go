package main

import (
	"fmt"
	"log"
	"math"
)

type Statistic struct {
	numberRuns int
	player1 Player
	sentence1 int
	player2 Player
	sentence2 int
}

type Analysis struct {
	sentences map[int][]Statistic
	strategies []Strategy
	player1Favored Strategy
	player2Favored Strategy
	bothFavored [2]Strategy
}


func main() {

	fmt.Println("Welcome to the Simulation of Prisoner's Dilemma!")
	fmt.Println()
	fmt.Print("Do you want to do simulation manually? y/n ")
	var reply string
	_, err := fmt.Scanf("%s", &reply)
	if err != nil {
		log.Fatal(err)
	}

	if reply == "y" || reply == "yes" {

		statistic := manualSimulation()
		statsDisplay(statistic)

	} else {

		fmt.Println()
		fmt.Println("Simuilation will run between those strategies below for two players: ")
		analysis := autoSimulation()

		fmt.Println()
		fmt.Print("Do you want to do see detail data? y/n ")
		var reply string
		_, err := fmt.Scanf("%s", &reply)
		if err != nil {
			log.Fatal(err)
		}

		if reply == "y" || reply == "yes" {
			analysisResults(analysis, true)
		}

		analysisResults(analysis, false)

	}


}

func statsDisplay(statistic Statistic) {
	fmt.Println("----------------------------")
	fmt.Println("|Simulation Result Analysis|")
	fmt.Println("----------------------------")

	fmt.Println("Player1's strategy: ", statistic.player1.strategy.display)
	fmt.Println("Player1's sentence: ", statistic.sentence1)
	fmt.Println("Player1's sentence percentage: ", float32(statistic.sentence1)/float32(statistic.numberRuns*3))
	fmt.Println("Player2's strategy: ", statistic.player2.strategy.display)
	fmt.Println("Player2's sentence: ", statistic.sentence2)
	fmt.Println("Player2's sentence percentage: ", float32(statistic.sentence2)/float32(statistic.numberRuns*3))
	fmt.Println("Total sentence percentage: ", float32(statistic.sentence1+statistic.sentence2) / float32(statistic.numberRuns*4))
}

func analysisResults(analysis Analysis, show bool){

	var p1Average []int
	var p2Average []int
	var bothAverage []int

	for i :=0; i < len(analysis.sentences); i++  {
		p1total := 0
		p2total := 0
		for j := 0; j < len(analysis.sentences[i]); j++  {
			p1total += analysis.sentences[i][j].sentence1
			p2total += analysis.sentences[j][i].sentence2
			bothAverage = append(bothAverage, analysis.sentences[i][j].sentence1 + analysis.sentences[i][j].sentence2)
		}
		p1Average = append(p1Average, int(float32(p1total) / 5))
		p2Average = append(p2Average, int(float32(p2total) / 5))
	}

	var p1MinIndex = 0
	var p2MinIndex = 0
	for i := 0; i < len(p1Average); i++ {
		if show {
			fmt.Println()
			fmt.Println("With Strategy Number " , i+1)
			fmt.Println("Player1 has a average sentence of: ", p1Average[i])
			fmt.Println("Player2 has a average sentence: ", p2Average[i])
		}

		if p1Average[i] < p1Average[p1MinIndex] {
			p1MinIndex = i
		}
		if p2Average[i] < p2Average[p2MinIndex] {
			p2MinIndex = i
		}
	}

	analysis.player1Favored = analysis.strategies[p1MinIndex]
	analysis.player2Favored = analysis.strategies[p2MinIndex]


	var bothMinIndex = 0
	for i := 0; i < len(bothAverage); i++  {
		if show {
			fmt.Println()
			str := fmt.Sprintf("With player on Strategy %d and playey2 on Strategy %d",
				(i / 5) + 1, (int(math.Mod(float64(i), 5)))+1)
			fmt.Println(str)
			fmt.Println("The total sentence they receive is: ", bothAverage[i])
		}
		if bothAverage[i] < bothAverage[bothMinIndex] {
			bothMinIndex = i
		}
	}

	analysis.bothFavored[0] = analysis.strategies[bothMinIndex / 5]
	analysis.bothFavored[1] = analysis.strategies[int(math.Mod(float64(bothMinIndex), 5))]


	fmt.Println()
	fmt.Println("-----------------------")
	fmt.Println("|Simulation Statistics|")
	fmt.Println("-----------------------")

	fmt.Println("Player1's best strategy is : " , analysis.player1Favored.display)
	fmt.Println("Player2's best strategy is : " , analysis.player2Favored.display)

	fmt.Println()
	fmt.Println("For mutual-best result: ")
	fmt.Println("Player1's strategy should be: " , analysis.bothFavored[0].display)
	fmt.Println("Player2's strategy should be: " , analysis.bothFavored[1].display)

}
