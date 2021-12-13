package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type BodyPart struct {
	MaxValue uint8
	Quantity uint8
}

type Beatle map[string]*BodyPart

func NewBeatle() Beatle {
	ret := make(Beatle)
	// Max values
	ret["Body"] = &BodyPart{1, 0}
	ret["Head"] = &BodyPart{1, 0}
	ret["Antennae"] = &BodyPart{2, 0}
	ret["Eyes"] = &BodyPart{2, 0}
	ret["Mouth"] = &BodyPart{1, 0}
	ret["Legs"] = &BodyPart{6, 0}
	return ret
}

func dicetopart(diceroll int) string {
	switch diceroll {
	case 1:
		return "Body"
	case 2:
		return "Head"
	case 3:
		return "Antennae"
	case 4:
		return "Eyes"
	case 5:
		return "Mouth"
	case 6:
		return "Legs"
	default:
		return ""
	}
}

func (beatle *Beatle) complete() bool {
	for _, part := range *beatle {
		if part.MaxValue > part.Quantity {
			return false
		}
	}
	return true
}

func (part *BodyPart) increment() {
	if part.MaxValue > part.Quantity {
		part.Quantity++
	}
}

func (beatle *Beatle) report() string {

	var ret string
	for i := 1; i < 7; i++ {
		part := dicetopart(i)
		ret = fmt.Sprintf("%s%s: %d (%d needed)\n", ret, part, (*beatle)[part].Quantity, (*beatle)[part].MaxValue)
	}
	return ret

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("The number of players must be specified as a command line argument.")
		os.Exit(1)
	}

	players, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if players < 2 {
		fmt.Println("Number of players must be above 2")
		os.Exit(1)
	}

	// Slice of Beatles, one for each player
	beatles := make([]Beatle, players)
	for i := 0; i < players; i++ {
		beatles[i] = NewBeatle()
	}

	prng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rollDice := func() int {
		return prng.Int()%6 + 1
	}

	var gameOver bool
	// Round loop
	for gameOver == false {
		// Player turn loop
		for i := 0; i < players; i++ {
			log := func(msg string) {
				fmt.Printf("Player %d: %s\n", i+1, msg)
			}
			var beatle *Beatle = &beatles[i]
			log("Turn started")
			diceroll := rollDice()
			part := dicetopart(diceroll)
			log("Dice rolled a " + strconv.Itoa(diceroll) + " (" + part + ")")
			if (*beatle)[part].MaxValue > (*beatle)[part].Quantity {
				log("Adding " + part)
				(*beatle)[part].increment()
				if beatle.complete() {
					log("Beatle complete: Won")
					gameOver = true
					break
				}
				log(part + " now has " + strconv.Itoa(int((*beatle)[part].Quantity)) + " pieces")
			} else {
				log("All pieces already collected")
			}
			log("Beatle state:")
			fmt.Print(beatle.report())
			log("Turn ended\n")
		}
	}

}
