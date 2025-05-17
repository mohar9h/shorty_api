package link

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}
