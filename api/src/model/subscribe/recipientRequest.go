package subscribe

type RecipientRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
