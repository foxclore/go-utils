package sauth

import (
	scauth "github.com/foxclore/go-utils/simple-cookie-auth"
	"github.com/gin-gonic/gin"
)

func AuthFunc(c *gin.Context, cookieName string) (bool, string) {
	if cookie, err := c.Cookie(cookieName); err != nil {
		return false, ""
	} else {
		return scauth.GetID(cookie)
	}
}
