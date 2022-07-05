package calculate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/rec-service/pkg/recommend"
)

func Calculate(c *gin.Context) {
	rec := recommend.Recommend(c)
	c.JSON(http.StatusOK, rec)
}
