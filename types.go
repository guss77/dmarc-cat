package main

import (
	"net"
)

// DateRange time period
type DateRange struct {
	Begin int64 `xml:"begin"`
	End   int64 `xml:"end"`
}

// ReportMetadata for the report
type ReportMetadata struct {
	OrgName          string    `xml:"org_name"`
	Email            string    `xml:"email"`
	ExtraContactInfo string    `xml:"extra_contact_info"`
	ReportID         string    `xml:"report_id"`
	Date             DateRange `xml:"date_range"`
	Errors           []string  `xml:"error"`
}

// PolicyPublished found in DNS
type PolicyPublished struct {
	Domain string `xml:"domain"`
	ADKIM  string `xml:"adkim"`
	ASPF   string `xml:"aspf"`
	P      string `xml:"p"`
	SP     string `xml:"sp"`
	Pct    int    `xml:"pct"`
	Fo     string `xml:"fo"`
}

// PolicyEvaluated what was evaluated
type PolicyEvaluated struct {
	Disposition string                 `xml:"disposition"`
	DKIM        string                 `xml:"dkim"`
	SPF         string                 `xml:"spf"`
	Reasons     []PolicyOverrideReason `xml:"reason,omitempty"`
}

// PolicyOverrideReason are the reasons that may affect DMARC disposition
// or execution thereof
type PolicyOverrideReason struct {
	Type    string `xml:"type"`
	Comment string `xml:"comment"`
}

// Row for each IP address
type Row struct {
	SourceIP net.IP          `xml:"source_ip"`
	Count    int             `xml:"count"`
	Policy   PolicyEvaluated `xml:"policy_evaluated"`
}

// Identifiers headers checked
type Identifiers struct {
	HeaderFrom   string `xml:"header_from"`
	EnvelopeFrom string `xml:"envelope_from"`
	EnvelopeTo   string `xml:"envelope_to,omitempty"`
}

// Result for each IP
type Result struct {
	Domain      string `xml:"domain"`
	Selector    string `xml:"selector"`
	Result      string `xml:"result"`
	HumanResult string `xml:"human_result"`
}

// AuthResults for DKIM/SPF
type AuthResults struct {
	DKIM Result `xml:"dkim,omitempty"`
	SPF  Result `xml:"spf,omitempty"`
}

// Record for each IP
type Record struct {
	Row         Row         `xml:"row"`
	Identifiers Identifiers `xml:"identifiers"`
	AuthResults AuthResults `xml:"auth_results"`
}

// Report the report itself
type Report struct {
	Version  float32         `xml:"version"`
	Metadata ReportMetadata  `xml:"report_metadata"`
	Policy   PolicyPublished `xml:"policy_published"`
	Records  []Record        `xml:"record"`
}

// An actual file can have multiple reports
type Feedback []Report
