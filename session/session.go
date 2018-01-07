package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetEmailFromSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	return session.Get("email")
}
