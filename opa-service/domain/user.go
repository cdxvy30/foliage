package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type RequestBody struct {
	UID string `json:"uid"`
}

type Claims struct {
	UID  string `json:"uid"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
