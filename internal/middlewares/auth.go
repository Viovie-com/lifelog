package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if len(h) == 0 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("access is denied (1)"))
			return
		}

		auth := strings.Fields(h)
		if len(auth) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("access is denied (2)"))
			return
		}

		if strings.ToLower(auth[0]) != "bearer" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("access is denied (3)"))
			return
		}

		member, err := db.GetMemberByToken(auth[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("access is denied (4)"))
			return
		}

		c.Set("member", member)

		return
	}
}

func GetMemberFromAuth(c *gin.Context) (member db.Member, err error) {
	m, hasAuth := c.Get("member")
	if !hasAuth {
		err = errors.New("access is denied (5)")
		return
	}
	member, ok := m.(db.Member)
	if !ok {
		err = errors.New("access is denied (6)")
		return
	}
	return
}
