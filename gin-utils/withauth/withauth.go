package withauth

import "github.com/gin-gonic/gin"

type AuthFunc func(c *gin.Context) (string, bool)

func OnAuth(auth AuthFunc, onSuccess, onFailure gin.HandlerFunc, c *gin.Context) {
	id, authenticated := auth(c)

	if authenticated {
		c.Set("internal_auth_id", id)
		onSuccess(c)
	} else {
		onFailure(c)
	}
}
