package main

type MoveChooser struct {
	safeMoves []string
}

func NewMoveChooser(state GameState) MoveChooser {
	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height
	if myHead.X == 0 {
		isMoveSafe["left"] = false
	} else if myHead.X == boardWidth-1 {
		isMoveSafe["right"] = false
	} else if myHead.Y == 0 {
		isMoveSafe["down"] = false
	} else if myHead.Y == boardHeight-1 {
		isMoveSafe["up"] = false
	}

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	mybody := state.You.Body
	for i := 1; i < len(mybody); i++ {
		if mybody[i].X == myHead.X && mybody[i].Y == myHead.Y {
			switch mybody[i].X {
			case myHead.X - 1:
				isMoveSafe["left"] = false
			case myHead.X + 1:
				isMoveSafe["right"] = false
			case myHead.Y - 1:
				isMoveSafe["down"] = false
			case myHead.Y + 1:
				isMoveSafe["up"] = false
			}
		}
	}

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	opponents := state.Board.Snakes
	for _, opponent := range opponents {
		if opponent.ID == state.You.ID {
			continue
		}
		for i := 0; i < len(opponent.Body); i++ {
			if opponent.Body[i].X == myHead.X && opponent.Body[i].Y == myHead.Y {
				switch opponent.Body[i].X {
				case myHead.X - 1:
					isMoveSafe["left"] = false
				case myHead.X + 1:
					isMoveSafe["right"] = false
				case myHead.Y - 1:
					isMoveSafe["down"] = false
				case myHead.Y + 1:
					isMoveSafe["up"] = false
				}
			}
			// Check if the opponent is in the same row or column as the head
			if opponent.Body[i].X == myHead.X {
				if opponent.Body[i].Y < myHead.Y {
					// Opponent is above the head
					if myHead.Y-opponent.Body[i].Y < 2 {
						// If the opponent is too close, don't move up
						isMoveSafe["up"] = false
					}
				} else if opponent.Body[i].Y > myHead.Y {
					// Opponent is below the head
					if opponent.Body[i].Y-myHead.Y < 2 {
						// If the opponent is too close, don't move down
						isMoveSafe["down"] = false
					}
				}
			} else if opponent.Body[i].Y == myHead.Y {
				if opponent.Body[i].X < myHead.X {
					// Opponent is left of the head
					if myHead.X-opponent.Body[i].X < 2 {
						// If the opponent is too close, don't move left
						isMoveSafe["left"] = false
					}
				} else if opponent.Body[i].X > myHead.X {
					// Opponent is right of the head
					if opponent.Body[i].X-myHead.X < 2 {
						// If the opponent is too close, don't move right
						isMoveSafe["right"] = false
					}
				}
			}
		}
	}

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	return MoveChooser{safeMoves: safeMoves}
}
