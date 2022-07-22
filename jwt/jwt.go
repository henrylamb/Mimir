package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(setClaimValue, setClaimName string, secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims[setClaimName] = setClaimValue
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in its claims.
//If the claim that has been found matches the va;ue will be returned and that identifaction value
//will be passed onto the next http handler func
func ParseToken(tokenStr string, findClaim string, secretKey []byte) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		foundClaim := claims[findClaim].(string)
		return foundClaim, nil
	} else {
		return "", err
	}
}

/*
**EXAMPLE AKIN TO ABOVE**
but for the above to work more const values would be needed along with hard coded values

type contextKey struct {
	name string
}
type users struct {
	id       string
	username string
	password string
}

func (u users) GetUsername() string {
	return u.username
}

var userCtxKey = &contextKey{"user"}

// ForContext finds the user from the context. REQUIRES Auth to have run.
func ForContext(ctx context.Context) *users {
	raw, _ := ctx.Value(userCtxKey).(*users)
	return raw
}

func Auth() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			id := users2.GetUsernameByID(username)
			if id != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), &contextKey{"user"}, &username)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			// not sure about what the next step within this function will actually do?
		})
	}
}



*/
