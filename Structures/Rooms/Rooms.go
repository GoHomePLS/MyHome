package Rooms

type Room struct {
	Name   string
	Length float32
	Width  float32
	Square float32
	Style  string
}

type Rooms struct {
	kitchen  Room
	bathroom Room
	bedroom  Room
	hall     Room
}

func AddRoom() Rooms {
	return Rooms{
		kitchen:  Room{Name: "LivingRoom", Length: 6.0, Width: 4.0, Square: 24.0, Style: "Feng Shui"},
		bathroom: Room{Name: "Bathroom", Length: 4.0, Width: 4.0, Square: 16.0, Style: "Feng Shui"},
		bedroom:  Room{Name: "Bedroom", Length: 5.0, Width: 3.0, Square: 15.0, Style: "Feng Shui"},
		hall:     Room{Name: "Hallway", Length: 8.0, Width: 2.0, Square: 16.0, Style: "Feng Shui"},
	}
}
