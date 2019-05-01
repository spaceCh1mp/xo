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
	turn int
	win  bool = false
	stat string
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
	size := flag.Int("Size", 3, "Board is an mxm matrix where m=3 by default")
	player1 := flag.String("Player 1", "X", "Player 1 is X by default")
	player2 := flag.String("Player 2", "O", "Player 2 is O by default")
	flag.Parse()
	//initialise and register players
	players := []string{*player1, *player2}
	p := initPiece(players)
	var g Game
	//startGame
	g.startGame(p, *size)
}
func initPiece(p []string) Piece {
	var a string = "human"
	var player Piece
	player.piece = make(map[int]string)
	player.player = make(map[string]Info)
	for index, value := range p {
		player.piece[index] = value
		player.player[value] = Info{agent: a}
	}
	return player
}
func (g Game) startGame(pl Piece, size int) {
	g.setBoard(size)
	var count int = 0
	turn = 0
	//total number of moves is (mxm)
	total := int(math.Exp2(float64(size)))
	for {
		fmt.Println(count)
		//display board markup and game board side by side
		displayBoard(g)
		//prompt user for input
		pl.play(turn, g)
		//update player moveset
		//check for win
		for _, sub := range wins {
			win = CheckWin(pl.player[pl.piece[turn]].moveSet, sub)
			if win {
				//announce winner
				stat = "W"
			}
		}
		//if no win and board is full call draw
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
	fmt.Printf("Player %d:\nSelect where to place your piece \"%s\" :-", turn+1, pl)
	_, err := fmt.Scanf("%d", &move)
	if err != nil {

	}
	p.player[pl] = Info{moveSet: append(p.player[pl].moveSet, move)}
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
	for i := range g.board {
		g.board[i] = make([]string, size)
		for index := range g.board[i] {
			g.board[i][index] = "*"
		}
	}
}
func findCoordinates(j, matrixSize int) (int, int) {
	x, y := int(math.Ceil(float64(j)/float64(matrixSize))-1), 2
	c := 0
	for {
		if j%matrixSize == 0 {
			y -= c
			break
		} else {
			j++
			c++
		}
	}
	log.Println(x, ":", y)
	return x, y
}
func CheckWin(super, sub []int) bool {
	super = sort.IntSlice(super)
	sub = sort.IntSlice(sub)
	var check int = 0
	for _, value := range super {
		for _, val := range sub {
			if val == value {
				check += 1
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
func displayBoard(board Game) {
	b := board.board
	s := len(b) - 1
	st := ""
	divider := fmt.Sprintf("\t\t-----------\n")
	str := `
		 1 | 2 | 3	
		 ---------
		 4 | 5 | 6
		 ---------
		 7 | 8 | 9
	`
	for index := range b {
		for i, value := range b[index] {
			if i == s {
				st += fmt.Sprintf(" %s \n", value)
			} else if i == 0 {
				st += fmt.Sprintf("\t\t %s |", value)
			} else {
				st += fmt.Sprintf(" %s |", value)
			}
		}
		if index != 2 {
			st += divider
		}
	}
	fmt.Printf("%s\n%s", str, st)
}
