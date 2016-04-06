package irr

// ArinEntry constitutes an entry into ARIN's Route Registry
type ArinEntry struct {
	Email Email
	Entry RouteRegistryEntry
}

// ArinEntryFlat is the flat version of ArinEntry
type ArinEntryFlat struct {
	From         string
	To           string
	Subject      string
	SMTPServer   string
	Route        string
	Description  string
	ASN          int
	NotifyEmail  string
	MaintainedBy string
	ChangedEmail string
	Source       string
}

// RouteRegistryEntry is the Route Registry entry itself
type RouteRegistryEntry struct {
	Route        string
	Description  string
	ASN          int
	NotifyEmail  string
	MaintainedBy string
	ChangedEmail string
	Source       string
}

// Email is the data for sending / receiving the email.
type Email struct {
	From       string
	To         string
	Subject    string
	SMTPServer string
}

// Logger represents a logging object for use by IRR.
type Logger struct {
	Verbose bool
}
