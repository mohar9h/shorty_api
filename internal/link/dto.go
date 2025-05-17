package link

type shortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}
