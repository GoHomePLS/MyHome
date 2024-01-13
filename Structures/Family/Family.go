package Family

type Member struct {
	Name  string
	Title string
	Age   int16
	Sex   string
}

type Fam struct {
	Family []Member
}

func AddFamily() Fam {
	var family []Member

	family = append(family, Member{Name: "Julia", Title: "Mother", Age: 46, Sex: "Female"})
	family = append(family, Member{Name: "Vladimir", Title: "Father", Age: 47, Sex: "Male"})
	family = append(family, Member{Name: "Peter", Title: "Brother", Age: 16, Sex: "male"})
	family = append(family, Member{Name: "Potap", Title: "Brother", Age: 4, Sex: "male"})
	return Fam{Family: family}
}
