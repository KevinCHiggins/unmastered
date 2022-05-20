package model

type Mix struct {
	Title       string
	Location    string
	Description string
}

var hardcodedMixes []Mix = []Mix{
	{Title: "Gimme The Key",
		Location:    "key",
		Description: "Electro house banger with 90s PC gaming theme"},
	{Title: "Tenebres",
		Location:    "tenebres",
		Description: "Some atmospheric and moody tech-house"},
}

func GetMixes() []Mix {
	return hardcodedMixes
}
