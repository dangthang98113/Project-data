package subscribe

type RecipientResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
