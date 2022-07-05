package calculate

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/rec-service/pkg/recommend"
)

func main(c *gin.Context) {
	rec := recommend.Recommend()
	c.JSON(http.StatusOK, rec)
}
