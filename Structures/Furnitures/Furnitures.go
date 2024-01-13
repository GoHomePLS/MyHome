package Furnitures

type Furniture struct {
	Type     string
	Age      int16
	Colour   string
	Material string
	Price    string
}

type FurnitureSet struct {
	FurnitureSet []Furniture
}

func AddFurnitureSet() FurnitureSet {
	var furn []Furniture
	furn = append(furn, Furniture{Type: "Chair", Age: 4, Colour: "Red", Material: "Wood", Price: "2000 p"})
	furn = append(furn, Furniture{Type: "Bed", Age: 6, Colour: "White", Material: "Wood", Price: "5000 p"})
	furn = append(furn, Furniture{Type: "TV", Age: 7, Colour: "Black", Material: "Plastic and glass", Price: "20000 p"})
	furn = append(furn, Furniture{Type: "Cupboard", Age: 3, Colour: "Gray", Material: "Plastic", Price: "6000 p"})
	furn = append(furn, Furniture{Type: "Microwave", Age: 2, Colour: "Black", Material: "Plastic", Price: "1000 p"})

	return FurnitureSet{FurnitureSet: furn}
}
