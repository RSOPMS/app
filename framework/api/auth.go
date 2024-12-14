package api

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// JwtHandler handles JWT authorization.
type JwtHandler struct {
	secret        []byte
	cookieName    string
	expiration    time.Time         // Default: 12 hours after initialization
	validMethods  []string          // Default: HS256
	signingMethod jwt.SigningMethod // Default: HS256
}

// JwtHandlerOptionFunc sets a JwtHandler field.
type JwtHandlerOptionFunc func(j *JwtHandler)

// WithExpiration sets JWT expiration time.
func WithExpiration(expiration time.Time) JwtHandlerOptionFunc {
	return func(j *JwtHandler) {
		j.expiration = expiration
	}
}

// WithSigningMethod sets JWT signing method.
func WithSigningMethod(method jwt.SigningMethod) JwtHandlerOptionFunc {
	return func(j *JwtHandler) {
		j.signingMethod = method
	}
}

// WithValidMethods sets JWT valid methods for parsing.
func WithValidMethods(validMethods []string) JwtHandlerOptionFunc {
	return func(j *JwtHandler) {
		j.validMethods = validMethods
	}
}

// NewJwtHandler creates a new JWT handler.
func NewJwtHandler(cookieName string, secret []byte, options ...JwtHandlerOptionFunc) *JwtHandler {
	j := &JwtHandler{
		secret:        secret,
		cookieName:    cookieName,
		expiration:    time.Now().Add(time.Duration(12 * time.Hour)),
		validMethods:  []string{"HS256"},
		signingMethod: jwt.SigningMethodHS256,
	}

	for _, optionFunc := range options {
		optionFunc(j)
	}

	return j
}

// createJwtString creates a new JWT token string for the given subject.
func (j *JwtHandler) createJwtString(subject string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(j.expiration),
	}

	token := jwt.NewWithClaims(j.signingMethod, claims)

	return token.SignedString(j.secret)
}

// ParseJwt parses and validates the given JWT string and returns its token.
func (j *JwtHandler) ParseJwt(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	}, jwt.WithValidMethods(j.validMethods))
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ProcessLogin validates a login request and returns a cookie with a JWT
// authorization cookie.
func (j *JwtHandler) ProcessLogin(subject, password, passwordHash string) (*http.Cookie, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return nil, err
	}

	tokenString, err := j.createJwtString(subject)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     j.cookieName,
		Value:    tokenString,
		Path:     "/",
		Expires:  j.expiration,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie, nil
}

// ProcessLogout processes a logout request by returning a cookie to delete the
// JWT.
func (j *JwtHandler) ProcessLogout() *http.Cookie {
	return &http.Cookie{
		Name:     j.cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}
