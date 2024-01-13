package Animals

type Pet struct {
	Type  string
	Name  string
	Age   int16
	Color string
	Sex   string
}

type Pets struct {
	Pets []Pet
}

func AddPet() Pets {
	var pets []Pet

	pets = append(pets, Pet{Type: "Chameleon", Name: "Habib", Age: 5, Color: "Different", Sex: "Male"})

	return Pets{Pets: pets}
}
