package main

import(
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe Game :)")
	var user1, user2 string
	fmt.Println("##############################")
	fmt.Println("Please enter your name: ")
	fmt.Scan(&user1)
	fmt.Println("##############################")
	fmt.Println("Please enter your name: ")
	fmt.Scan(&user2)
	fmt.Println("##############################")

	r := rand.Intn(2)
	var player1, player2 string
	if r == 0{
		player1 = user1
		player2 = user2
		fmt.Printf("The starter of the game is the %s\n", user1)
	}else {
		player1 = user2
		player2 = user1
		fmt.Println("The starter of the game is the %s\n", user2)
	}



	board := [][]string{
	{" - ", " - ", " - "},
	{" - ", " - ", " - "},
	{" - ", " - ", " - "}}



	i := 0
	for {
		if i > 8{
			fmt.Println("HAS NOT WINNER :(")
			var s int
			for{
				fmt.Println("1 - NEW GAME\n2 - EXIT GAME")
				fmt.Scan(&s)
				if s == 1 || s == 2{
					break
				}else{
					fmt.Println("PLEASE ENTER THE CORRECT NUMBER")
				}
			}
			if s == 1{
				i = 0
				board = [][]string{
					{" - ", " - ", " - "},
					{" - ", " - ", " - "},
					{" - ", " - ", " - "}}
				tmp := player1
				player1 = player2
				player2 = tmp
			}
			if s == 2{
				break
			}
		}else {
			printBoard(board)
			i++
			setVal(board, i, player1, player2)
			status := winLose(board)
			if status == "X"{
				fmt.Println("##################################")
				fmt.Printf("the %s Win the Game :)\n", player1)
				printBoard(board)
				break
			}else if status == "O"{
				fmt.Printf("the %s Win the Game :)\n", player2)
				printBoard(board)
				break
			}
		}
	}
}

var a, b int
func setVal(board [][]string, i int, pl1 string, pl2 string){
	if i % 2 == 1{
		fmt.Printf("it is your turn %s\n", pl1)
	}else{
		fmt.Printf("it is your turn %s\n", pl2)
	}
	
	for{
		fmt.Println("Please enter the X and Y: ")
		fmt.Scanf("%d %d", &a, &b)
		if a >= 1 && b >= 1 && a <= 3 && b <= 3 {
			break
		}
		fmt.Println("PLEASE ENTER THE RANGE OF GAME:")
	}
	a -= 1
	b -= 1
	if i % 2 == 1{
		board[a][b] = " X "
	}else{
		board[a][b] = " O "
	}
}

func printBoard(b [][]string){
	fmt.Println("******************************")
	fmt.Println("-----the board-----")
	fmt.Println("    1   2   3")
	for i := 0 ; i < len(b); i++{
		fmt.Printf("%d:", i + 1)
		fmt.Println(b[i])
	}
	fmt.Println("******************************")
}


func winLose(b [][]string)string{
	var flagX bool
	var flagY bool
	// for check row 
	for i := 0 ; i < len(b); i++{
		flagX = true
		flagY = true
		for j := 0 ; j < len(b); j++{
			if b[i][j] != " X "{
				flagX = false
			}
			if b[i][j] != " O "{
				flagY = false
			}
		}
		if flagX{
			return "X"
		}
		if flagY{
			return "O"
		}
	}

	// for check column
	for i := 0 ; i < len(b); i++{
		flagX = true
		flagY = true
		for j := 0; j < len(b); j++{
			if b[j][i] != " X "{
				flagX = false
			}
			if b[j][i] != " O "{
				flagY = false
			}
		}
		if flagX{
			return "X"
		}
		if flagY{
			return "O"
		}
	}

	flagX = true
	flagY = true

	//for check a cross
	for i := 0 ; i < len(b) ; i++{
		if b[i][i] != " X "{
			flagX = false
		}
		if b[i][i] != " O "{
			flagY = false
		}
	}
	if flagX{
		return "X"
	}
	if flagY{
		return "O"
	}



	flagX = true
	flagY = true




	for i, j:= 0, 2; i < 3 && j > -1; i, j = i + 1, j - 1{
		if b[i][j] != " X "{
			flagX = false
		}
		if b[i][j] != " O "{
			flagY = false
		}
	}
	if flagX{
		return "X"
	}
	if flagY{
		return "O"
	}

	// return nothing
	return "N"
}