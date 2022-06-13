package types

type DeleteListingsItemResponse struct {
	Sku          string        `json:"sku"`
	Status       string        `json:"status"`
	SubmissionID string        `json:"submissionId"`
	Issues       []interface{} `json:"issues"`
}
