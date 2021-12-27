package sauth

import (
	scauth "github.com/foxclore/go-utils/simple-cookie-auth"
	"github.com/gin-gonic/gin"
)

var cookieName = ""

func SetCookieName(cookie string) {
	cookieName = cookie
}

func AuthFunc(c *gin.Context) (bool, string) {
	if cookie, err := c.Cookie(cookieName); err != nil {
		return false, ""
	} else {
		return scauth.GetID(cookie)
	}
}
