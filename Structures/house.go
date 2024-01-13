package house

type House struct {
	Family    Fam.Family
	Animals   Pets.Animals
	Rooms     rooms.Rooms
	Gadgets   gadgets.Gadgets
	Furniture furniture.FurnitureSet
}

func AddHouse() House {
	return House{

		Family:    family.AddFamily(),
		Pets:      pets.AddPet(),
		Rooms:     rooms.AddRoom(),
		Gadgets:   gadgets.AddGadget(),
		Furniture: furniture.AddFurnitureSet(),
	}
}
