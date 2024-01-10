package jwtauth

import (
	
	"errors"
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	IAuth interface {
		SignToken() string
	}

	Claims struct {
		EmployeeId string `json:"emp_id"`
		RoleId     string `json:"role_id"`
	}

	AuthMapClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	authConcrete struct {
		Secret []byte
		Claims *AuthMapClaims `json:"claims"`
	}

	accessToken  struct{ *authConcrete }
	refreshtoken struct{ *authConcrete }
	apiKey       struct{ *authConcrete }
)

func (a *authConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	tokenString, _ := token.SignedString(a.Secret)
	return tokenString
}

func now() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func jwtTimeDurationCal(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(now().Add(time.Duration(t) * time.Second))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAccessToken(secret string, expiredAt int64, claims *Claims) IAuth {
	return &accessToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "fcc",
					Subject:   "access-token",
					Audience:  []string{"fcc"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expiredAt int64, claims *Claims) IAuth {
	return &refreshtoken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "fcc",
					Subject:   "refresh-token",
					Audience:  []string{"fcc"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expiredAt int64, claims *Claims) string {
	obj := &refreshtoken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "fcc",
					Subject:   "refresh-token",
					Audience:  []string{"fcc"},
					ExpiresAt: jwtTimeRepeatAdapter(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}

	return obj.SignToken()
}

func NewApiKey(secret string) IAuth {
	return &apiKey{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: &Claims{},
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "fcc",
					Subject:   "api-key",
					Audience:  []string{"fcc"},
					ExpiresAt: jwtTimeDurationCal(31560000),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil, errors.New("error: token is invalid")
		}
	}

	if claims, ok := token.Claims.(*AuthMapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("error: claims type is invalid")
	}
}

