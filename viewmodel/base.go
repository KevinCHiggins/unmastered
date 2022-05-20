package viewmodel

type Base struct {
	SiteTitle     string
	PageTitle     string
	ActiveNavLink string
}

func NewBase(pageTitle, path string) Base {
	return Base{"Unmastered", pageTitle, path}
}
