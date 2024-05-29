package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/templates/pages"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/idtoken"
)

func Login(c *Context) error {
	sessionCookie, err := c.Cookie(logic.SessionCookieName)
	if err != nil || sessionCookie.Value == "" {
		blank := logic.NewSessionCookie("", time.Unix(0, 0))
		c.SetCookie(&blank)
		return c.RenderPage(pages.Login())
	}

	_, err = logic.GetAndVerifyUserSession(c.Request().Context(), c.DB(), sessionCookie.Value, logic.StringToIdentifier(c.RealIP()))
	if err != nil {
		return c.RenderPage(pages.Login())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/app")
}

func Logout(c *Context) error {
	blank := logic.NewSessionCookie("", time.Unix(0, 0))
	c.SetCookie(&blank)

	sessionCookie, err := c.Cookie(logic.SessionCookieName)
	if err != nil || sessionCookie.Value == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	session, err := logic.GetAndVerifyUserSession(c.Request().Context(), c.DB(), sessionCookie.Value, logic.StringToIdentifier(c.RealIP()))
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	err = logic.DeleteSession(c.Request().Context(), c.DB(), session.ID)
	if err != nil {
		log.Err(err).Str("sessionId", session.ID).Msg("failed to delete a session")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func GoogleAuthRedirect(c *Context) error {
	cookieToken, err := c.Cookie("g_csrf_token")
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context="+err.Error())
	}

	bodyToken := c.FormValue("g_csrf_token")
	if bodyToken == "" || cookieToken.Value == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context=missing token")
	}
	if bodyToken != cookieToken.Value {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context=missing token")
	}

	credential := c.FormValue("credential")
	if credential == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context=missing credential")
	}

	payload, err := idtoken.Validate(context.Background(), credential, logic.GoogleAuthClientID)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context="+err.Error())
	}

	googleUser, err := logic.GoogleTokenInfo(credential)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context="+err.Error())
	}
	if payload.Audience != googleUser.Aud || payload.Issuer != googleUser.Issuer || payload.Subject != googleUser.Subject {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context=bad user info received")
	}

	if googleUser.EmailVerified != "true" {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=You need to verify your Google Account before using it to log in")
	}
	if googleUser.Name == "" || googleUser.Email == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Your Google Account is incomplete&context=missing name or email")
	}

	var opts []db.UserSetParam
	if googleUser.Picture != "" {
		opts = append(opts, db.User.ProfilePicture.Set(googleUser.Picture))
	}

	user, err := c.DB().User.FindUnique(db.User.ExternalID.Equals(googleUser.Subject)).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		user, err = c.DB().User.CreateOne(db.User.Email.Set(googleUser.Email), db.User.ExternalID.Set(payload.Subject), db.User.Name.Set(googleUser.Name), opts...).Exec(c.Request().Context())
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context="+err.Error())
	}

	session, err := logic.NewUserSession(c.Request().Context(), c.DB(), user.ID, logic.StringToIdentifier(c.RealIP()), logic.SessionExpiration7Days())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to log in with Google&context="+err.Error())
	}

	sessionCookie := logic.NewSessionCookie(session.CookieValue, session.Expiration)
	c.SetCookie(&sessionCookie)
	c.SetUser(user)

	// Update the user record if needed
	go func() {
		var updates []db.UserSetParam
		if user.Email != googleUser.Email {
			updates = append(updates, db.User.Email.Set(googleUser.Email))
		}
		if user.Name != googleUser.Name {
			updates = append(updates, db.User.Name.Set(googleUser.Name))
		}
		if pic, _ := user.ProfilePicture(); googleUser.Picture != "" && googleUser.Picture != pic {
			updates = append(updates, db.User.ProfilePicture.Set(googleUser.Picture))
		}

		if len(updates) == 0 {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := c.DB().User.FindUnique(db.User.ExternalID.Equals(googleUser.Subject)).Update(updates...).Exec(ctx)
		if err != nil {
			log.Err(err).Str("externalId", googleUser.Subject).Str("userId", user.ID).Msg("failed to update user record")
		}
	}()

	return c.Redirect(http.StatusTemporaryRedirect, "/app")
}
