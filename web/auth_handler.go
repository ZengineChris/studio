package web

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/core/auth"
)

type writeCloserWrapper struct {
	io.Writer
}

func (w writeCloserWrapper) Close() error {
	return nil // No-op close for bytes.Buffer
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func LoginHandler(c echo.Context) error {
	auth, err := auth.New()
	if err != nil {
		return err
	}
	state, err := generateRandomState()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["state"] = state
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
}

func LogoutHandler(c echo.Context) error {
	return nil
}


// TODO: store a reference for the user in the database 
func AuthCallbackHandler(c echo.Context) error {
	auth, err := auth.New()
	if err != nil {
		return err
	}
	sess, err := session.Get("session", c)
	if c.QueryParam("state") != sess.Values["state"] {
		return c.String(http.StatusBadRequest, "Invalid state parameter.")
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(c.Request().Context(), c.QueryParam("code"))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
	}

	idToken, err := auth.VerifyIDToken(c.Request().Context(), token)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to verify ID Token.")
	}

	// TODO find the data we need 
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess.Values["access_token"] = token.AccessToken
	sess.Values["profile"] = profile
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Redirect to logged in page.
	return c.Redirect(http.StatusTemporaryRedirect, "/user")
}
