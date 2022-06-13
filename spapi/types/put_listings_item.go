package types

type PutListingsItemResponse struct {
	Sku          string   `json:"sku"`
	Status       string   `json:"status"`
	SubmissionID string   `json:"submissionId"`
	Issues       []Issues `json:"issues"`
}
type Issues struct {
	Code           string   `json:"code"`
	Message        string   `json:"message"`
	Severity       string   `json:"severity"`
	AttributeNames []string `json:"attributeNames"`
}
