package model

type SummaryReport struct {
	TotalComplaints int
	Resolved        int
	Pending         int
	SlaCompliance   string
}
