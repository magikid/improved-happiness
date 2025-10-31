package main

import "math/rand/v2"

func (mc MoveChooser) getNextMove(state GameState) BattlesnakeMoveResponse {
	if len(mc.safeMoves) == 0 {
		return BattlesnakeMoveResponse{Move: "down", Shout: "Avenge me brothers!"}
	}

	// TODO: Move towards food instead of random, to regain health and survive longer

	return getRandomMove(mc)
}

func getRandomMove(state MoveChooser) BattlesnakeMoveResponse {
	nextMove := state.safeMoves[rand.IntN(len(state.safeMoves))]
	return BattlesnakeMoveResponse{Move: nextMove}
}
