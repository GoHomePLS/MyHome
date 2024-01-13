package Gadgets

type Gadgets struct {
	Type   string
	Age    int16
	Model  string
	Colour string
	Price  string
}

type Gadget struct {
	Gadget []Gadgets
}

func AddFurnitureSet() Gadget {
	var furn []Gadgets
	furn = append(furn, Gadgets{Type: "Smartphone", Age: 1, Colour: "Blue", Model: "Iphone", Price: "59867395829826 p"})
	furn = append(furn, Gadgets{Type: "Computer", Age: 3, Colour: "White", Model: "HP", Price: "54000 p"})
	furn = append(furn, Gadgets{Type: "TV", Age: 7, Colour: "Black", Model: "Sony", Price: "20000 p"})

	return Gadget{Gadget: furn}
}
