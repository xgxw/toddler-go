package middlewares

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

var jwtMid = JWTMiddleware{
	key:           []byte("aasdasfqw"),
	signingMethod: jwt.SigningMethodHS256,
}

func Test_JWTMiddleware_signToken(t *testing.T) {
	SkipConvey("Test", t, func() {
		expires := time.Now().Add(time.Hour)
		standClaims := new(jwt.StandardClaims)
		standClaims.ExpiresAt = expires.Unix()
		fmt.Println(jwtMid.signToken(&payloadClaims{
			StandardClaims: standClaims,
		}))
	})
}

func Test_JWTMiddleware_verifyToken(t *testing.T) {
	SkipConvey("Test", t, func() {
		fmt.Println(jwtMid.verifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoid3pzIn0.aICvPf2gQV7bNSVB5wBNax1keQCoi7iHev-5ak8Jlvs"))
	})
}
