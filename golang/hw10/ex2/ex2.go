package main

import ( 
	"math/rand" 
	"time" 
	"strconv"
	"fmt"
)

func shuffleDeck(cards []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(cards); i++ {
		var r = rand.Intn(i + 1)
		if i != r {
			cards[r], cards[i] = cards[i], cards[r]
		}
	}
}

func makeDeck() []int {
	var deck []int = make([]int, 52)
	var i = len(deck)
	for i > 0 {
		deck[i-1] = i
		i--
	}
	shuffleDeck(deck)
	return deck
}

func getCardVal(x int, isPlayer bool) int {
	var y int = (x % 13) + 1
	if y > 10 { 
		return 10
	} else if y == 1 {
		if isPlayer {
			var s string = ""
			var x int = 0
			for x != 11 || x != 1 { 
				fmt.Printf("Count your ace as 1 or 11? ")
				fmt.Scanln(&s)
				x, _ = strconv.Atoi(s)
				if x == 1 || x == 11 { return x }
			}
		} else {
			rand.Seed(time.Now().UnixNano())
			var x = rand.Intn(1)
			if x == 1 {
				fmt.Println("The dealer counts his ace as an 11!")
				return 11
			} else {	
				fmt.Println("The dealer counts his ace as a 1!")
				return 1
			}
		}
	} else {
		return y
	}
	return y
}

func getPrintVal(x int) string {
	var y int = (x % 13) + 1
	if y == 1 {
		return "Ace"
	} else if y == 11 {
		return "Jack"
	} else if y == 12 {
		return "Queen"
	} else if y == 13 {
		return "King"
	} else {
		var s string = strconv.Itoa(y)
		return s
	} 
	return ""
}

func getCard(cards []int, i int) int {
	return cards[i]
}

func getCardSuit(x int) string {
	var y int = x / 13	
	if y == 0 { 
		return "Hearts"
	} else if y == 1 {
		return "Spades"
	} else if y == 2 {
		return "Diamonds"
	} else {
		return "Clubs"
	}
}

func checkDone() bool {
	var s string = ""
	for true {
		fmt.Printf("The game is done, do you want to play again (y/n)? ")
		fmt.Scanln(&s)
	
		if s == "y" {
			return false
		} else if s == "n" {
			return true
		} else {
			fmt.Println("Please enter either y or n!")
		}
	}
	return true
}

func main() {
        var done bool = false
	for !done {
		var deck = makeDeck()	
		var dealersHand = 0
		var yourHand = 0
		var on = 0

		var yourCard = getCard(deck, on)
		on++
		var dealerCard = getCard(deck, on)
		on++
		fmt.Printf("The dealer puts a card down for you - it is a %s of %s\n", getPrintVal(yourCard), getCardSuit(yourCard))	
		fmt.Println("The dealer places a card down for himself - it is face down")
		dealersHand = dealersHand + getCardVal(dealerCard, false)
		yourHand = yourHand + getCardVal(yourCard, true)	
		
		dealerCard = getCard(deck, on) 
		on++
		yourCard = getCard(deck, on)
		on++
		fmt.Printf("The dealer puts a card for himself - it is a %s of %s\n", getPrintVal(dealerCard), getCardSuit(dealersHand))
		fmt.Printf("The dealer puts a card down for you - it is a %s of %s\n", getPrintVal(yourCard), getCardSuit(yourCard))	
		dealersHand = dealersHand + getCardVal(dealerCard, false)
		yourHand = yourHand + getCardVal(yourCard, true)	
		
		var inPlay bool = true
		var roundOn int = 0
		for inPlay {
			inPlay = false
			if (dealersHand > 21 && yourHand < 21) {
				fmt.Println("Dealer is bust, you win!")
				done = checkDone()
			} else if (dealersHand < 21 && yourHand > 21) {
				fmt.Println("Your bust, you lose!")
				done = checkDone()
			} else if (dealersHand > 21 && yourHand > 21) {
				fmt.Println("Your both bust, it's a tie!")
				done = checkDone()
			} else if (dealersHand == 21 && yourHand !=  21) {
				fmt.Println("Dealer gets 21, you lose!")
				done = checkDone()
			} else if (dealersHand != 21 && yourHand == 21) {
				fmt.Println("You get 21, you win!")
				done = checkDone()
			} else if (dealersHand == 21 && yourHand == 21) {
				fmt.Println("You both get 21, it's a tie!")
				done = checkDone()
			} else {
				if roundOn > 0 {
					if (yourHand > dealersHand) {
						fmt.Println("Your hand is higher, you win!")
						done = checkDone()
					} else if (dealersHand > yourHand) {
						fmt.Println("The dealers hand is higher, you lose!")
						done = checkDone()
					} else {
						fmt.Println("Both of your hands are the same, its a tie!")
						done = checkDone()
					}
				} else {
					var s string = ""
					fmt.Println("Since both your hands are less then 21, you can either hit or stand")
					fmt.Printf("The value of your hand is %d\n", yourHand)
					var isStanding bool = false
					for !isStanding {
						fmt.Printf("Hit or Stand? ") 
						fmt.Scanln(&s)
						if s == "hit" {
							yourCard = getCard(deck, on)
							on++ 
							fmt.Printf("The dealer puts a card down for you - it is a %s of %s\n", getPrintVal(yourCard), getCardSuit(yourCard))	
							yourHand = yourHand + getCardVal(yourCard, true)
							fmt.Printf("The value of your hand is %d\n", yourHand)
							if yourHand > 21 {
								fmt.Println("You bust!")
								isStanding = true
							}
						} else if s == "stand" { 
							isStanding = true 
						} else {
							fmt.Println("Please enter either hit or stand!")
						}
					}

					for dealersHand < 17 { 
						fmt.Println("The dealer is hitting!")
						dealerCard = getCard(deck, on)
						on++
						fmt.Printf("The dealer puts a card down for himself - it is a %s of %s\n", getPrintVal(dealerCard), getCardSuit(dealerCard))
						dealersHand = dealersHand + getCardVal(dealersHand, false)
						fmt.Printf("The dealers hands value is %d\n", dealersHand)
					}
					if (dealersHand > 17 && dealersHand < 21) {
						fmt.Println("The dealer stands")
					} else { 
						fmt.Println("Dealer goes bust!")
					}
					inPlay = true	
				} 
			}
			roundOn++
		}
	}
}
