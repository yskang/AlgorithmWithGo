package discoveryGo

type Address struct {
	City	string
	State 	string
}

type Telephone struct {
	Mobile	string
	Direct 	string
}

type Contact struct {
	Address
	Telephone
}
