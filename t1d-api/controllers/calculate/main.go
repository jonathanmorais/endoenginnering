package calculate

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/rec-service/pkg/recommend"
)

func main(ctx gin.Context) {
	rec := recommend.Recommend()
}
