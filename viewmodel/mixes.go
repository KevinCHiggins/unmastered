package viewmodel

import "intive/unmastered/model"

type Mixes struct {
	SiteTitle     string
	PageTitle     string
	ActiveNavLink string
	Mixes         []Mix
}

func NewMixes() Mixes {
	//var result Mixes
	var vmMixes []Mix
	modelData := model.GetMixes()
	for _, m := range modelData {
		vmMix := Mix{
			"Unmastered",
			"Mixes - " + m.Title,
			"mixes",
			m.Title,
			m.Location,
			m.Description,
		}
		vmMixes = append(vmMixes, vmMix)
	}
	return Mixes{"Unmastered", "Mixes", "mixes", vmMixes}
}
