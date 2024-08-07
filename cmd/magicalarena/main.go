package main

import (
	// Go Internal Packages
	"fmt"

	// Local Packages
	match "magicalarena/internal/match"
	player "magicalarena/internal/player"
	x "magicalarena/internal/utils/constants"
	helpers "magicalarena/internal/utils/helpers"
)

func main() {
	helpers.PrintWelcomeMessage()

	player1, err1 := player.NewPlayer("Karna", 100, 5, 10)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	player2, err2 := player.NewPlayer("Arjuna", 50, 10, 5)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	helpers.PrintPlayerDetails(player1)
	helpers.PrintPlayerDetails(player2)

	match1 := match.CreateNewMatch(player1, player2)
	winner := match1.Fight()
	fmt.Printf(x.GREEN+"\n\n🎉 Hurray! %s Won 🎉\n"+x.RESET, winner)
}
