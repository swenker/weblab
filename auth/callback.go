package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"yinyushijing.cn/xnote-api/asession"

	oidc "github.com/coreos/go-oidc"
)

// CallbackHandler is a callback handler that Auth0 will call once it redirects to your app.
// oidc implements OpenID Connect client logic for the golang.org/x/oauth2 package.
//
func CallbackHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
	session, err := asession.Store.Get(c.Request, "auth-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(r.URL)
	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(c.Writer, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := NewAuthenticator()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), c.Request.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(c.Writer, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "YA5vlXQk4u0vo1V2wGVCJsg80ry5DUW7",
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		http.Error(c.Writer, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(r, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to logged in page
	c.Redirect(http.StatusSeeOther, "/api/user")
}
