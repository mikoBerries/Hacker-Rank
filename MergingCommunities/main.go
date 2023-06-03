package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Hello World")

	/* getting my Go legs...
	// this is one empty slice:
	one := make([]int, 0)
	one = append(one, 1)
	one = append(one, 2)
	// this is big slice of slices:
	many := make([][]int, 3)
	many[0] = one
	// fmt.Println(many)
	... That all works */

	// I should stop reading all stdin to a buffer,
	// and learn to use the GO line-by-line idioms.
	// Irks me though, tending towards inefficient.

	sc := bufio.NewScanner(os.Stdin)
	needTopLine := true
	var n int
	var q int
	var person [][]int
	// idiomatic go: loop ends on eof or err:
	for sc.Scan() {
		line := strings.TrimRight(strings.TrimRight(sc.Text(), "\n"), "\r")
		if needTopLine {
			twain := strings.Split(line, " ")
			if len(twain) != 2 {
				panic("top line")
			}
			n, _ = strconv.Atoi(twain[0])
			q, _ = strconv.Atoi(twain[1])
			// fmt.Println(n, q)
			// Generate the slice of N slices.
			// HackerRank's people index 1-up.
			person = make([][]int, n+1)
			for i := 1; i <= n; i++ {
				person[i] = append(make([]int, 0), i)
			}
			// fmt.Println(person)
			needTopLine = false
			continue
		}
		// else not top line
		tokens := strings.Split(line, " ")
		if len(tokens) < 1 {
			panic("empty line")
		}
		switch tokens[0] {
		case "M":
			i, _ := strconv.Atoi(tokens[1])
			j, _ := strconv.Atoi(tokens[2])
			// fmt.Printf("Merge %v, %v\n", i, j)

			// check assumption
			if i == j {
				// THIS OCCURED: panic("i==j")
				// Expletives deleted!
				continue
			}

			iLeader := person[i][0]
			jLeader := person[j][0]

			// check for yet another stumble.
			// Adding test stopped my panics.
			if iLeader == jLeader {
				continue
			}

			// for sanity check
			leniL := len(person[iLeader])
			lenjL := len(person[jLeader])

			if len(person[iLeader]) < len(person[jLeader]) {
				// merge shorter iLeader into longer jLeader
				// append list of followers
				person[jLeader] = append(person[jLeader], person[iLeader]...)
				// tell everyone in iLeader, they now follow jLeader
				for _, k := range person[iLeader] {
					person[k][0] = jLeader
				}
				// in case iLeader had followers, drop them
				person[iLeader] = person[iLeader][:1]

				// sanity check
				if person[i][0] != person[j][0] {
					panic("1st i-j test")
				}
				if person[iLeader][0] != person[jLeader][0] {
					panic("1st iLeader-jLeader test")
				}
				if len(person[i]) != 1 {
					panic("len i not 1 test")
				}
				if len(person[iLeader]) != 1 {
					panic("len iLeader not 1 test")
				}
				if len(person[jLeader]) != leniL+lenjL {
					panic("len jLeader not sum test")
				}

			} else {
				// merge poss. shorter jLeader into poss. longer iLeader
				// append list of followers
				person[iLeader] = append(person[iLeader], person[jLeader]...)
				// tell everyone in jLeader, they now follow iLeader
				for _, k := range person[jLeader] {
					person[k][0] = iLeader
				}
				// in case jLeader had followers, drop them
				person[jLeader] = person[jLeader][:1]

				// sanity check
				if person[i][0] != person[j][0] {
					panic("2nd i-j test")
				}
				if person[iLeader][0] != person[jLeader][0] {
					panic("2nd iLeader-jLeader test")
				}
				if len(person[j]) != 1 {
					panic("len j not 1 test")
				}
				if len(person[jLeader]) != 1 {
					panic("len jLeader not 1 test")
				}
				if len(person[iLeader]) != leniL+lenjL {
					panic("len iLeader not sum test")
				}

			}
			break
		case "Q":
			i, _ := strconv.Atoi(tokens[1])
			// fmt.Printf("Report on %v\n", i)
			iLeader := person[i][0]
			fmt.Println(len(person[iLeader])) // desired answer
			break
		default:
			panic("not M, not Q")
			break
		}
		// way way TMI -- // fmt.Println(person)
		_ = q
	}
	if err := sc.Err(); err != nil && err != io.EOF {
		panic(err)
	}
	// fmt.Println("fini")
}
