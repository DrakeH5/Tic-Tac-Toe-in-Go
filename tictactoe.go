package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	grid := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	var gameOver bool = false
	for gameOver == false {
		//draw board
		for i := 0; i < 3; i++ {
			fmt.Println(grid[i])
		}
		//ask for next move
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Column:")
		scanner.Scan()
		inputC, C := strconv.Atoi(scanner.Text())
		fmt.Println("Row:")
		scanner.Scan()
		inputR, R := strconv.Atoi(scanner.Text())
		//draw player move
		if C == nil && R == nil {
			if grid[inputC-1][inputR-1] == " " {
				grid[inputC-1][inputR-1] = "X"
			}
		}
		///check game over
		players := [2]string{"X", "Y"}
		for i := 0; i < 2; i++ {
			checkPlayer := players[i]
			//horizontal check
			for column := 0; column < 3; column++ {
				for i := 0; i < 3; i++ {
					if grid[column][i] != checkPlayer {
						i = 4
					}
					if i == 2 {
						gameOver = true
						fmt.Println(checkPlayer, " Won!")
					}
				}
			}
			//vertical check
			for row := 0; row < 3; row++ {
				for i := 0; i < 3; i++ {
					if grid[i][row] != checkPlayer {
						i = 4
					}
					if i == 2 {
						gameOver = true
						fmt.Println(checkPlayer, " Won!")
					}
				}
			}
			//diagonal check
			if grid[1][1] == checkPlayer {
				if grid[0][0] == checkPlayer {
					if grid[2][2] == checkPlayer {
						gameOver = true
						fmt.Println(checkPlayer, " Won!")
					}
				}
				if grid[0][2] == checkPlayer {
					if grid[2][0] == checkPlayer {
						gameOver = true
						fmt.Println(checkPlayer, " Won!")
					}
				}
			}
		}
		//tie check
		if gameOver == false {
			tie := true
			for row := 0; row < 2; row++ {
				for column := 0; column < 2; column++ {
					if grid[column][row] == " " {
						tie = false
					}
				}
			}
			if tie == true {
				gameOver = true
				fmt.Println("Tie!")
			} else {
				//Y moves
				grid = Ymoves(grid)
			}
		}
	}
	//draw board
	for i := 0; i < 3; i++ {
		fmt.Println(grid[i])
	}
}

func Ymoves(grid [3][3]string) [3][3]string {
	genratedMove := false
	for genratedMove == false {
		YmoveX := rand.Intn(3)
		YmoveY := rand.Intn(3)
		if grid[YmoveY][YmoveX] == " " {
			grid[YmoveY][YmoveX] = "Y"
			genratedMove = true
		}
	}
	return grid
}

