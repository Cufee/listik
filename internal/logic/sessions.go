package logic

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cufee/shopping-list/prisma/db"
	"github.com/rs/zerolog/log"
)

const SessionCookieName = "lk-session"

func NewSessionCookie(value string, expiration time.Time) http.Cookie {
	return http.Cookie{
		Name:     "lk-session",
		Value:    value,
		Expires:  expiration,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
}

func SessionExpiration7Days() time.Time {
	return time.Now().Add(time.Hour * 24 * 7)
}

func SessionExpiration30Days() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func StringToIdentifier(input string) Identifier {
	return Identifier(HashString(input))
}
func NewUserSession(ctx context.Context, client *db.PrismaClient, userID string, identifier Identifier, expiration time.Time) (*db.SessionModel, error) {
	cookieValue, err := RandomString(32)
	if err != nil {
		return nil, err
	}

	session, err := client.Session.CreateOne(db.Session.CookieValue.Set(cookieValue), db.Session.Identifier.Set(identifier.String()), db.Session.Expiration.Set(expiration), db.Session.User.Link(db.User.ID.Equals(userID))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetAndVerifyUserSession(ctx context.Context, client *db.PrismaClient, sessionValue string) (*db.SessionModel, error) {
	session, err := client.Session.FindFirst(db.Session.CookieValue.Equals(sessionValue)).With(db.Session.User.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	if session.Expiration.Before(time.Now()) {
		return nil, errors.New("session has expired")
	}
	if session.User() == nil {
		return nil, errors.New("session in invalid")
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.Session.FindUnique(db.Session.ID.Equals(session.ID)).Update(db.Session.LastUsed.Set(time.Now())).Exec(ctx)
		if err != nil {
			log.Err(err).Str("sessionId", session.ID).Msg("failed to update session")
		}
	}()

	return session, nil
}

func UpdateSessionExpiration(ctx context.Context, client *db.PrismaClient, sessionID string, newExpiration time.Time) (*db.SessionModel, error) {
	session, err := client.Session.FindUnique(db.Session.ID.Equals(sessionID)).Update(db.Session.Expiration.Set(newExpiration)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func DeleteSession(ctx context.Context, client *db.PrismaClient, sessionID string) error {
	_, err := client.Session.FindUnique(db.Session.ID.Equals(sessionID)).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return err
}
