package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func (mc MoveChooser) getNextMove(state GameState) BattlesnakeMoveResponse {
	if len(mc.safeMoves) == 0 {
		log.Printf("Ran out of safe moves %v", mc.safeMoves)
		return BattlesnakeMoveResponse{Move: "down", Shout: "Avenge me brothers!"}
	}

	// TODO: Move towards food instead of random, to regain health and survive longer

	return getRandomMove(mc)
}

func getRandomMove(state MoveChooser) BattlesnakeMoveResponse {
	nextMove := state.safeMoves[rand.IntN(len(state.safeMoves))]
	log.Printf("Moving %s")
	return BattlesnakeMoveResponse{Move: nextMove, Shout: fmt.Sprintf("I guess I'll go %s then", nextMove)}
}
