package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload               //subject
	claims["exp"] = now.Add(ttl).Unix()   //expiration time
	claims["iat"] = now.Unix()            //issued at
	claims["nbf"] = now.Unix()            //notbefore

	// Create the token with RS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the RSA private key
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return signedToken, nil
}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	// Step 1: Decode the public key from a base64 string to its original form
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key: %w", err)
	}

	// Step 2: Parse the decoded public key so we can use it to verify the token
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: could not parse public key: %w", err)
	}

	// Step 3: Parse the token and verify its signature using the parsed public key
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check if the signing method is RSA
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: could not parse token: %w", err)
	}

	// Step 4: Extract claims from the token if it is valid
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: token is invalid")
	}

	// Step 5: Return the "sub" claim from the token
	return claims["sub"], nil
}

