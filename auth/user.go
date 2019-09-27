package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yinyushijing.cn/xnote-api/asession"
)

func UserHandler(c *gin.Context) {

	r := c.Request
	w := c.Writer

	session, err := asession.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//templates.RenderTemplate(w, "user", session.Values["profile"])
	c.JSON(http.StatusOK, session.Values["profile"])
}
