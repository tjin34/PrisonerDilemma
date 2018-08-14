# Idea of this project:

This is a Prisoner's Dilemma Simulator that programmed with Go Programming language.

Instruction on how to compile and run is in section 2.

I came up with 5 strategies for this project. But this project is absolutly expandable to as more as strategies you want.

The 5 Strategies are:
I implemented this project with 5 strategies for each player.
1. Fully trust on opponent and remain silent! (100% co-operate, 0% defect)
2. Mostly trust opponent, but sometimes defect! (70% co-operate, 30% defect)
3. Half trust opponent, and randomly pick its decision! (50% co-operate, 50% defect)
4. Mostly defect opponent, but sometimes remain silent (30% co-operate, 70% defect)
5. Doesn't trust opponent and always back door defect! (0% co-operate, 100% defect)

And for outcome, if they both co-operate, they both receive 1 year of jail.
If one defect  and one co-operate, the one co-operate will receive 3 year of jail, and the one defect receive nothing.
If both defect, they both receive 2 year of jail.

There is choice for user to run one single simulation with their choice of strategies for player1 and player2.
And there is a AutoSimulation option that simulates all possible combinations of strategies 1000 times each.
It will give out the strategies that favor player1 the most, favor player2 the most, and favor both.

I hope you had fun simulating Prisoner's Dilemma with this.

Any question, please email me @ tonyjzy@gmail.com.


# HOW TO RUN THIS PROJECT:

1. Please make sure you have installed GO on your device, if not csa2 machine would work too.

2. Second, download the zip file and unzip everything

3. Set your GOROOT, PATH, GOPATH environment variables.

export GOROOT=/go installation path/
export PATH=$GOROOT/bin
export GOPATH=.../PrisonerDilemma   (... is the directory that you put this project in)


4. In your command line: cd $GOPATH, and then: go build main

5. run the project with: ./main

6. follow the instruction during execution

7. Please notice, every time you want to run the project, you need to start from step 3 and set environment variables first.
