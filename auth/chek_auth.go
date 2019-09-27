package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"yinyushijing.cn/xnote-api/asession"
)

func IsAuthenticated() gin.HandlerFunc {
	fmt.Println("-------------auth called........")

	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		session, err := asession.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Println(session.Values["profile"])
		if _, ok := session.Values["profile"]; !ok {
			c.Redirect(http.StatusSeeOther, "/v/home.html")
		} else {
			c.Next()
		}

		return
	}
}
