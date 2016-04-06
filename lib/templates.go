package irr

const (
	// DefaultToEmail is the e-mail address to use when emailing ARIN about the RR
	DefaultToEmail = "rr@arin.net"
	// DefaultToEmail = "kevin.kirsche@one.verizon.com"

	// DefaultFromEmail is the email address to use when responding / sending from CRS
	DefaultFromEmail = "crs-crew@verizon.com"
	// DefaultNotifyEmail is who should be notified once the update goes through
	DefaultNotifyEmail = "crs-crew@verizon.com"
	// DefaultChangedByEmail is who should be notified once the update goes through
	DefaultChangedByEmail = "crs-crew@verizon.com"
	// DefaultSMTPServer is the SMTP server to use when sending the email
	DefaultSMTPServer = "smtp.vzbi.com:smtp"
	// DefaultMntBy is the group maintaining the entry
	DefaultMntBy = "MNT-MCICS"
	// DefaultSubject is the subject line for the ARIN email
	DefaultSubject = "route"
	// DefaultSource The source line for the ARIN email
	DefaultSource = "ARIN"
	// DefaultDescription the description for use in the email description
	DefaultDescription = "Verizon Internet Services"

	// ArinEntryTemplate is the template to execute when sending a new route
	// to ARIN
	ArinEntryTemplate = "To: {{ .To }}\r\n" +
		"From: {{ .From }}\r\n" +
		"Subject: {{ .Subject }}\r\n" +
		"\r\n" +
		"route: {{ .Route }}\r\n" +
		"descr: {{ .Description }}\r\n" +
		"origin: AS{{ .ASN }}\r\n" +
		"notify: {{ .NotifyEmail }}\r\n" +
		"mnt-by: {{ .MaintainedBy }}\r\n" +
		"changed: {{ .ChangedEmail }}\r\n" +
		"source: {{ .Source }}\r\n"
)
