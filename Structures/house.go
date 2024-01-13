package house

import (
	"HouseP/Structures/Animals"
	"HouseP/Structures/Family"
	"HouseP/Structures/Furnitures"
	"HouseP/Structures/Gadgets"
	"HouseP/Structures/Rooms"
)

type Dom struct {
	Animals Animals.Pets
	Family Family.Fam
	Furnitures Furnitures.FurnitureSet
	Gadgets Gadgets.Gadget
	Rooms Rooms.Rooms
}

func AddHouse() Dom {
	return Dom{
		Animals: Animals.AddPet(),
		Family: Family.AddFamily(),
		Furnitures: Furnitures.AddFurnitureSet(),
		Gadgets: Gadgets.AddGadgets(),
		Rooms: Rooms.AddRoom(),
	}
}
