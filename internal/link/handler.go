package link

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "shorty_api/internal/utils"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = response.ValidationError(c.Writer, c.Request, err)
		return
	}

	ip := c.ClientIP()

	code, err := h.usecase.Shorten(c.Request.Context(), req.URL, ip)
	if err != nil {
		_ = response.Error(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	shortURL := c.Request.Host + "/link/" + code
	_ = response.Success(c.Writer, shortURL, "Short URL created successfully")
}

func (h *Handler) Redirect(c *gin.Context) {
	code := c.Param("code")

	original, err := h.usecase.GetOriginalURL(c.Request.Context(), code)
	if err != nil {
		_ = response.NotFound(c.Writer, "Link not found or expired")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, original)
}
