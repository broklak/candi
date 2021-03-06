package middleware

import (
	"context"
	"errors"
	"strings"

	"pkg.agungdwiprasetyo.com/candi/candishared"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/config/env"
)

// Middleware impl
type Middleware struct {
	tokenValidator      interfaces.TokenValidator
	username, password  string
	authTypeCheckerFunc map[string]func(context.Context, string) (*candishared.TokenClaim, error)
}

// NewMiddleware create new middleware instance
func NewMiddleware(tokenValidator interfaces.TokenValidator) *Middleware {
	mw := &Middleware{
		tokenValidator: tokenValidator,
		username:       env.BaseEnv().BasicAuthUsername,
		password:       env.BaseEnv().BasicAuthPassword,
	}

	mw.authTypeCheckerFunc = map[string]func(context.Context, string) (*candishared.TokenClaim, error){
		Basic: func(ctx context.Context, key string) (*candishared.TokenClaim, error) {
			return nil, mw.Basic(ctx, key)
		},
		Bearer: func(ctx context.Context, token string) (*candishared.TokenClaim, error) {
			return mw.Bearer(ctx, token)
		},
	}

	return mw
}

func extractAuthType(prefix, authorization string) (string, error) {
	authValues := strings.Split(authorization, " ")
	if len(authValues) == 2 && strings.ToLower(authValues[0]) == prefix {
		return authValues[1], nil
	}

	return "", errors.New("Invalid authorization")
}
