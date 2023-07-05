package main

import(
	"fmt"
)

func main() {
	// board := [3][3]string{
	// 	[]string{" --- ", " --- ", " --- "},
	// 	[]string{" --- ", " --- ", " --- "},
	// 	[]string{" --- ", " --- ", " --- "},
	// }


	board := [][]string{
	{" O ", " - ", " X "},
	{" - ", " X ", " - "},
	{" X ", " - ", " - "}}
	


	printBoard(board)
	fmt.Println(winLose(board))



}


func printBoard(b [][]string){
	for i := 0 ; i < len(b); i++{
		fmt.Println(b[i])
	}
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