package main

import (
	"bufio"
	"fmt"
	"os"
	"sniptech/sawmill"
	"strconv"
	"strings"
)

func userCli() []int {
	scanner := bufio.NewScanner(os.Stdin)
	userInputs := make([][]int, 0)
	trunks := make([]int, 0)

	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		text = strings.TrimSpace(text)

		if text == "" {
			fmt.Println("Wrong value")
			return nil
			// exit if user entered an empty string
		}
		if text != "0" {
			if len(text) > 1 {
				var input = strings.Split(text, " ")

				_, err := strconv.Atoi(input[0])
				if err != nil {
					fmt.Println(err)
					return nil
				}
				for i := 1; i <= len(input[1:]); i++ {
					j, err := strconv.Atoi(input[i])
					if err == nil {
						if j >= 0 {
							fmt.Println("after first number, each number should be positive")
							return nil
						}
					}
					if err != nil {
						fmt.Println(err)
					}
					trunks = append(trunks, j)
				}
				userInputs = append(userInputs, trunks)
			}
		} else if text == "0" {
			return userInputs[0]
		}
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
		return nil
	}

	return nil
}

func main() {
	trunkList := make([][]int, 0)
	maxProfit := 0

	trunks := userCli()

	Perm(trunks, func(r []int) {
		t := make([]int, len(r))
		copy(t, r)

		output := make([]int, 0)

		sawmill.SawTrunk(t, sawmill.SawBlocks, &output)
		prf := sawmill.Profit(output)

		if prf >= maxProfit {
			if prf > maxProfit {
				trunkList = make([][]int, 0)
			}

			maxProfit = prf

			for _, t := range trunkList {
				if sawmill.Equal(t, r) {
					return
				}
			}

			trunkList = append(trunkList, r)
		}
	})

	fmt.Printf("Case: %d\n", 1)
	fmt.Printf("Max profit: %d\n", maxProfit)
	fmt.Printf("Order: %v\n", trunkList)
}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		aa := make([]int, len(a))
		copy(aa, a)
		f(aa)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
