package viewmodel

import "intive/unmastered/model"

type Remixes struct {
	SiteTitle     string
	PageTitle     string
	ActiveNavLink string
	Remixes       []Remix
}

func NewRemixes() Remixes {
	//var result Mixes
	var vmRemixes []Remix
	modelData := model.GetRemixes()
	for _, rm := range modelData {
		vmRemix := Remix{
			"Unmastered",
			"Remixes - " + rm.Title,
			"remixes",
			rm.Title,
			rm.Producer,
			rm.Base,
			rm.Link,
			rm.Location,
			rm.ProducersNotes,
			rm.CuratorsNotes,
		}
		vmRemixes = append(vmRemixes, vmRemix)
	}
	return Remixes{"Unmastered", "Remixes", "remixes", vmRemixes}
}
