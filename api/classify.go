package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func Classify(c *gin.Context) {
	var classify home.HomeClassify
	res := classify.GetClassify()
	c.IndentedJSON(200, &res)
}
