package main

func cmdLook(state *GameState) {
	if state.Inventory["lidLighter"] {
		text := "You look around. You can make out shelf and a table in one corner, a bed in the other and straight in front of you"
		textPacer(text)
		state.Knowlege["room"]
	} else {
		text := "You look up and around. Different shades of black."
		textPacer(text)
	}
}

func cmdLookShelf(state *GameState) {
	if state.Knowldege["room"] {
		text := "The shelf is dusty and looks like it will collapse on it self at any time. Besides books and other kind of junk you can make out a little box. There seems to be some kind of electronic lock and a little display asks: 'What is the current temperature in Eberswalde?"
		textPacer(text)
		state.Knowledge["box"]
	} else {
		text := "Its dark you cant see anything"
		textPacer(text)
	}
}

func cmdLookTable(state *GameState) {
	if state.Knowldege["room"] {
		text := "Looks like little TV or computer. There does not seem to be a powersource. Looking closer you can make out that batteries are missing"
		textPacer(text)
		state.Knowledge["app"]
	} else {
		text := "Its dark you cant see anything"
		textPacer(text)
	}
}

func cmdLookBed(state *GameState) {
	if state.Knowldege["room"] {
		text := "The bed is dirty and empty bottles are scattered around it. However, you seem to see something blinking under the bed. Seems to be a key"
		textPacer(text)
		state.Knowledge["key"]
	} else {
		text := "Its dark you cant see anything"
		textPacer(text)
	}
}

func cmdPickPockets(state *GameState) {
	if state.Knowledge["lighterInPockets"] {
		text := "Nothing more than dust in your pockets"
		textPacer(text)
	} else if !state.Knowledge["lighterInPockets"] {
		text := "You rummage around in your pockets and you feel something that seems to be a little lighter"
		textPacer(text)
		state.Knowledge["lighterInPockets"] = true
	}
}
func cmdLookBed(state *GameState) {
	if state.Knowldege["room"] {
		text := ""
		textPacer(text)
		state.Knowledge["app"]
	} else {
		text := "Its dark you cant see anything"
		textPacer(text)
	}
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
