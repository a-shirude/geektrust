package model

type Card struct {
	Name           string
	Balance        int
	CurrentStation string
	PreviosStation string
}

type Passenger struct { // ADULT, SENIOR_CITIZEN, KID
	Type  string
	Fare  int
	Count int
}

type Station struct { // CENTRAL, AIRPORT
	Name       string
	Collection int //
	Discount   int //
	Passenger  []Passenger
}
