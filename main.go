package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Player ...
type Player struct {
	*tl.Entity
	prevX int
	prevY int
}

// Tick ...
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

// Collide ...
func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(30)
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})
	type rect struct {
		x     int
		y     int
		w     int
		h     int
		color tl.Attr
	}
	for _, rect := range []rect{
		{10, 10, 14, 17, tl.ColorWhite},
		{11, 11, 12, 10, tl.ColorBlack},
		{12, 12, 10, 8, tl.ColorWhite},
		{13, 13, 8, 6, tl.ColorBlack},
		{13, 14, 1, 1, tl.ColorGreen},
		{15, 14, 1, 1, tl.ColorGreen},
		{17, 14, 1, 1, tl.ColorGreen},
		{14, 16, 2, 1, tl.ColorGreen},
		{17, 16, 2, 1, tl.ColorGreen},
		{15, 21, 4, 1, tl.ColorBlack},
		{11, 22, 12, 4, tl.ColorBlack},
		{11, 22, 1, 1, tl.ColorWhite},
		{22, 22, 1, 1, tl.ColorWhite},
		{14, 23, 1, 1, tl.ColorWhite},
		{16, 23, 1, 1, tl.ColorWhite},
		{18, 23, 1, 1, tl.ColorWhite},
		{20, 23, 1, 1, tl.ColorWhite},
		{13, 24, 1, 1, tl.ColorWhite},
		{15, 24, 1, 1, tl.ColorWhite},
		{17, 24, 1, 1, tl.ColorWhite},
		{19, 24, 1, 1, tl.ColorWhite},
	} {
		level.AddEntity(tl.NewRectangle(
			rect.x*2,
			rect.y,
			rect.w*2,
			rect.h,
			rect.color,
		))
	}
	player := Player{tl.NewEntity(1, 1, 2, 1), 0, 0}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	g.Screen().SetLevel(level)
	g.Start()
}
