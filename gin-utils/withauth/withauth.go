package withauth

import "github.com/gin-gonic/gin"

type AuthFunc func(c *gin.Context) (bool, string)

func OnAuth(auth AuthFunc, onSuccess, onFailure gin.HandlerFunc, c *gin.Context) {
	authenticated, id := auth(c)

	if authenticated {
		c.Set("internal_auth_id", id)
		onSuccess(c)
	} else {
		onFailure(c)
	}
}
