package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

var (
	wins = [][]int{
		{1, 2, 3},
		{1, 4, 7},
		{1, 5, 9},
		{2, 5, 8},
		{3, 5, 7},
		{3, 6, 9},
		{4, 5, 6},
		{7, 8, 9}}
	turn, count, total int
	win                bool = false
	stat               string
)

type Game struct {
	board     [][]string
	protected []int
	size      int
}
type Piece struct {
	piece  map[int]string
	player map[string]Info
}
type Info struct {
	moveSet []int
	agent   string
}

func main() {
	//set game premise. size, player 1 and 2
	size := flag.Int("size", 3, "Board is an mxm matrix where m=3 by default")
	player1 := flag.String("player1", "X", "Player 1 is X by default")
	player2 := flag.String("player2", "O", "Player 2 is O by default")
	flag.Parse()
	//initialise and register players
	players := []string{*player1, *player2}
	p := initPiece(players)
	var g Game
	//startGame
	g.startGame(p, *size)
}
func initPiece(p []string) Piece {
	a := "human"
	var player Piece
	player.piece, player.player = make(map[int]string), make(map[string]Info)
	for index, value := range p {
		player.piece[index] = value
		player.player[value] = Info{agent: a}
	}
	return player
}
func (g Game) startGame(pl Piece, size int) {
	g.setBoard(size)
	count, turn = 0, 0
	//total number of moves is (mxm)
	total = int(math.Exp2(float64(size)))
	for {
		//display board markup and game board side by side
		displayBoard(g)
		//prompt user for input
		pl.play(turn, g)
		//check for win
		playerMoves := pl.player[pl.piece[turn]].moveSet
		for _, sub := range wins {
			win = CheckWin(playerMoves, sub)
			if win {
				stat = "W"
			}
		}
		//if no win and board is full call draw
		log.Println(count, ":", total)
		if count == total && !win {
			stat = "D"
		}
		//else continue for next player
		switch stat {
		case "W":
			displayBoard(g)
			msg(fmt.Sprintf("%s Won", pl.piece[turn]))
		case "D":
			displayBoard(g)
			msg("A draw")
		default:
			//update count and turn
			turn = findTurn(turn)
			count++
		}
	}
}
func (p *Piece) play(t int, g Game) {
	//update player set
	var move int
	pl := p.piece[t]
	fmt.Printf("Player %d:\n", turn+1)
	for {
		fmt.Printf("Select where to place your piece \"%s\" :- ", pl)
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("Wrong input, try again.")
		} else {
			if move < 1 || move > total+1 {
				fmt.Println("Position doesn't exist on board")
			} else {
				if g.protected[move-1] == move {
					fmt.Println("Position on the board is already occupied.")
				} else {
					break
				}
			}
		}
	}
	temp := p.player[pl].moveSet
	p.player[pl] = Info{moveSet: append(temp, move)}
	g.updateBoard(pl, move)
}

func (g *Game) updateBoard(player string, m int) {
	x, y := findCoordinates(m, g.size)
	g.board[x][y] = player
	g.protected[m-1] = m
}
func (g *Game) setBoard(size int) {
	g.board = make([][]string, size)
	g.protected = make([]int, int(math.Exp2(float64(size)))+1)
	g.size = size
	for in := range g.board {
		g.board[in] = make([]string, size)
		for index := range g.board[in] {
			g.board[in][index] = "*"
		}
	}
}
func findCoordinates(value, matrixSize int) (int, int) {
	x, y, initialValue := int(math.Ceil(float64(value)/float64(matrixSize))-1), matrixSize-1, value
	for {
		if value%matrixSize == 0 {
			y -= (value - initialValue)
			break
		} else {
			value++
		}
	}
	return x, y
}
func CheckWin(super, sub []int) bool {
	super, sub = sort.IntSlice(super), sort.IntSlice(sub)
	check := 0
	for _, value := range super {
		for _, val := range sub {
			if val == value {
				check += 1
				continue
			}
		}
		if check == len(sub) {
			return true
		}
	}
	return false
}
func msg(str string) {
	fmt.Println(str)
	os.Exit(1)
}
func findTurn(t int) int {
	if t == 0 {
		return 1
	} else {
		return 0
	}
}
func displayBoard(b Game) {
	board := b.board
	matrixEdge, boardUI, layout := len(board)-1, "", ""
	divider := fmt.Sprintf("\t\t-----------\n")
	l := 1
	for index := range board {
		for i, value := range board[index] {
			if i == matrixEdge {
				boardUI += fmt.Sprintf(" %s \n", value)
				layout += fmt.Sprintf(" %d \n", l)
			} else if i == 0 {
				boardUI += fmt.Sprintf("\t\t %s |", value)
				layout += fmt.Sprintf("\t\t %d |", l)
			} else {
				boardUI += fmt.Sprintf(" %s |", value)
				layout += fmt.Sprintf(" %d |", l)
			}
			l++
		}
		if index != matrixEdge {
			boardUI += divider
			layout += divider
		}
	}
	fmt.Printf("%s\n%s", layout, boardUI)
}
