package server

import "github.com/gin-gonic/gin"

type RequestBody struct {
	List []string `json:"list"`
}

func successResp(result []string) map[string]any {
	return gin.H{
		"success": true,
		"result":  result,
	}
}

func failResp(msg string) map[string]any {
	return gin.H{
		"success": false,
		//"result":  result,
		"msg": msg,
	}
}
