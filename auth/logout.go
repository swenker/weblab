package auth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"yinyushijing.cn/xnote-api/asession"
)

func LogoutHandler(c *gin.Context) {

	r := c.Request
	w := c.Writer

	domain := "dev-yinyushijing.auth0.com"

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic(err.Error())
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost/v/home.html")
	parameters.Add("client_id", "YA5vlXQk4u0vo1V2wGVCJsg80ry5DUW7")
	Url.RawQuery = parameters.Encode()

	session, err := asession.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(session.Values, "profile")
	fmt.Println(session.Values["profile"])
	c.Redirect(http.StatusTemporaryRedirect, Url.String())
}
