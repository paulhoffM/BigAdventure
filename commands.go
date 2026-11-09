package main

func cmdLook(state *GameState) {
	text := "You look up and around. Different shades of black."
	textPacer(text)
}
func cmdPickPockets(state *GameState) {
	text := "You rummage around in your pockets and you feel something that seems to be a little lighter"
	textPacer(text)
	state.Knowledge["lighterInPockets"] = true
}

func cmdLightLighter(state *GameState) {
	if state.Knowledge["lighterInPockets"] && !state.Inventory["lidLighter"] {
		text := "There will be light!"
		textPacer(text)
		lighterASCII()
		state.Inventory["lidLighter"] = true
	} else if state.Inventory["lidLighter"] {
		text := "The flame is already burning and the darkness only lingers in the corners."
		textPacer(text)
	} else {
		text := "Do you know if you have a lighter?"
		textPacer(text)
	}
}
