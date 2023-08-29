package config

// Passenger Type
const (
	Adult        = "ADULT"
	SenioCitizen = "SENIOR_CITIZEN"
	Kid          = "KID"
)

// Station Type
const (
	Central = "CENTRAL"
	Airport = "AIRPORT"
)

// Fare
const (
	AdultFare              = 200
	SenioCitizenFare       = 100
	KidFare                = 50
	InsufficientBalFeePerc = 2
)

// Command line args
const (
	PrintSummary = "PRINT_SUMMARY"
	CheckIn      = "CHECK_IN"
	Balance      = "BALANCE"
)

// Print Summary  args
const (
	TotalCollection = "TOTAL_COLLECTION"
	PassTypeSummary = "PASSENGER_TYPE_SUMMARY"
)
