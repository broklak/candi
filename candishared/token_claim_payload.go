package candishared

import "github.com/dgrijalva/jwt-go"

// TokenClaim for token claim data
type TokenClaim struct {
	jwt.StandardClaims
	Additional map[string]interface{}
}
