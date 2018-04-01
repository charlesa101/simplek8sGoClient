package core

import "github.com/gin-gonic/gin"

func HttpHeaderWithCorsSupport() gin.HandlerFunc {

	return func(httpHeader *gin.Context) {
		httpHeader.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		httpHeader.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		httpHeader.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		httpHeader.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		httpHeader.Next()
	}

}