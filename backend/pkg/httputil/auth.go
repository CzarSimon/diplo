package httputil

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	UserHeader = "X-UserID"
)

var (
	ErrNoAuthHeader        = NewError(http.StatusBadRequest, "No Authorization header provided")
	ErrMalformedAuthHeader = NewError(http.StatusBadRequest, "Badly formed Authorization header")
)

// AuthMiddleware creates gin middleware for authentication via bearer tokens.
func AuthMiddleware(opts *AuthOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if opts.routeIsExempt(c.Request) {
			c.Next()
			return
		}

		tokenString, err := GetTokenString(c)
		if err != nil {
			JSONError(c, err)
			return
		}

		userID, err := ValidateToken(tokenString, opts)
		if err != nil {
			JSONError(c, err)
			return
		}

		c.Set(UserHeader, userID)
		c.Next()
	}
}

// ValidateToken checks if token is valid and returns the Subject if so.
func ValidateToken(tokenString string, opts *AuthOptions) (string, error) {
	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{}
	err = token.Claims(opts.JwtSecret, &cl)
	if err != nil {
		return "", err
	}

	err = cl.ValidateWithLeeway(jwt.Expected{
		Issuer: opts.TokenIssuer,
	}, opts.ValidationLeeway)
	if err != nil {
		return "", ErrAccessDenied
	}
	return cl.Subject, nil
}

// GetTokenString gets authorization token from the request context.
func GetTokenString(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", ErrMalformedAuthHeader
	}
	return tokenParts[1], nil
}

// AuthOptions options for checking if a client authenticated.
type AuthOptions struct {
	JwtSecret        []byte
	TokenIssuer      string
	ValidationLeeway time.Duration
	ExemptedRoutes   map[string]bool
}

// NewAuthOptions creates new authentication options
func NewAuthOptions(jwtSecret, tokenIssuer string, exemptedRoutes ...string) *AuthOptions {
	routesSet := make(map[string]bool)
	for _, exemptedRoute := range exemptedRoutes {
		routesSet[exemptedRoute] = true
	}

	return &AuthOptions{
		JwtSecret:        []byte(jwtSecret),
		TokenIssuer:      tokenIssuer,
		ValidationLeeway: jwt.DefaultLeeway,
		ExemptedRoutes:   routesSet,
	}
}

// routeIsExempt checks if a route should be exempt from authentication checks.
func (opts AuthOptions) routeIsExempt(r *http.Request) bool {
	log.Printf("Route: %s\n", r.URL.String())
	_, ok := opts.ExemptedRoutes[r.URL.String()]
	return ok
}
