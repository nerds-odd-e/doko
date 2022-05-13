package final_day_poker_hands

type Game struct {
	MyHand       Hand
	OpponentHand Hand
}

func NewGame(cards []string) *Game {
	return &Game{
		MyHand:       cards[:5],
		OpponentHand: cards[5:],
	}
}

func (game Game) isP1Winner() bool {
	return game.MyHand.Wins(game.OpponentHand)
}