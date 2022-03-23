package middleware

import "github.com/gin-gonic/gin"

func TestASCII(ctx *gin.Context) {
	data := map[string]interface{}{
		"Name": "Go语言",
		"Age":  18,
	}
	ctx.AsciiJSON(200, data)
}
