package middleware

import (
	"github.com/ekkinox/fx-template/modules/fxauthenticationcontext"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AuthMiddleware is the authentication middleware.
type AuthMiddleware struct {
	config *fxconfig.Config
}

// NewAuthMiddleware creates a new AuthMiddleware.
func NewAuthMiddleware(config *fxconfig.Config) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

// Handle handles the middleware. 403 if no auth context.
func (m *AuthMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authContext := c.Get(fxauthenticationcontext.AuthenticationContextKey)
			if authContext == nil {
				// 403
				return c.JSON(
					http.StatusUnauthorized,
					map[string]any{"message": "Forbidden."},
				)
			}

			return next(c)
		}
	}
}
