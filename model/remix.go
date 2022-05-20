package model

type Remix struct {
	Title          string
	Producer       string
	Base           string
	Link           string
	Location       string
	ProducersNotes string
	CuratorsNotes  string
}

var hardcodedRemixes []Remix = []Remix{
	{Title: "Keyed Up",
		Producer:       "Thomas",
		Base:           "Gimme The Key",
		Link:           "https://audiojack.com/thomas/keyed-up",
		Location:       "keyed-up",
		ProducersNotes: "Slowed it down and added some bodhrán",
		CuratorsNotes:  ""},
	{Title: "Gimme",
		Producer:       "Suzanne Carla",
		Base:           "Gimme The Key",
		Link:           "https://instagram.com/soozie/349025743025",
		Location:       "gimme",
		ProducersNotes: "Nice groove!",
		CuratorsNotes:  "My longtime collaborator Soozie added some magnificent silky vocals here."},
}

func GetRemixes() []Remix {
	return hardcodedRemixes
}
