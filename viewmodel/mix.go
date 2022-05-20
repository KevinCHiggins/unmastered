package viewmodel

// needs to be refactored because
// when a slice of these is used in the
// Mixes viewmodel, the SiteTitle, PageTitle
// and ActiveNavLink fields are useless
type Mix struct {
	SiteTitle      string
	PageTitle      string
	ActiveNavLink  string
	MixTitle       string
	MixLocation    string
	MixDescription string
}
