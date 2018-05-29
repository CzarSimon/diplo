package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(opts AuthOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if opts.routeIsExempt(c.Request) {
			c.Next()
			return
		}
		c.Next()
	}
}

type AuthOptions struct {
	JwtSecret      []byte
	ExemptedRoutes map[string]bool
}

func NewAuthOptions(jwtSecret string, exemptedRoutes ...string) AuthOptions {
	routesSet := make(map[string]bool)
	for _, exemptedRoute := range exemptedRoutes {
		routesSet[exemptedRoute] = true
	}

	return AuthOptions{
		JwtSecret:      []byte(jwtSecret),
		ExemptedRoutes: routesSet,
	}
}

func (opts AuthOptions) routeIsExempt(r *http.Request) bool {
	_, ok := opts.ExemptedRoutes[r.URL.String()]
	return ok
}
