package main

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ2c2tmWUdxMkowSHhaR1FHZnJMZUxBZTZQVG1iSzdUWVBhTVNmN1dNUlRjIn0.eyJleHAiOjE3NzYzMDM5MjUsImlhdCI6MTc3NjMwMzYyNSwianRpIjoiOWY2YTYyNmEtZTk5Ni00NzYzLWJkNjctNTI0OTVkOGJhZmJmIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9leHBlbnNlLXRyYWNrZXIiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiOTA4NmU3YmItZDkyMi00MDY0LWIxYmQtOTg5ODQ4N2EwN2NjIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibXktYXBwIiwic2lkIjoiMzJjYjYxZGUtNjQ4My00OWUzLWI4ODgtMmZiYjhkOWQxODM3IiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLWV4cGVuc2UtdHJhY2tlciIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiRmVybmFuZG8gTWVuZGVzIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiZmVybmFuZG8iLCJnaXZlbl9uYW1lIjoiRmVybmFuZG8iLCJmYW1pbHlfbmFtZSI6Ik1lbmRlcyIsImVtYWlsIjoiZmVybmFuZG9AdGVzdGUuY29tIn0.g2ngNW2twqfH4dtfbnMVO3PI-qkld_nL_j-f0YZrnmsuww-KZ0CBbnNYtHBx82PrSkh-kgzwvMuS_imyLLIZFMpQ3qnNCPckiD9Q9JAByqmW0HUicXyN7jfdNWHY6vlpi2nCGaRA7TADANlnHbxiwgIL669y6qmZKny7hqGq-vL6Hq3f4ZvdF6m8C1fdzSWiVH0z0VQGkjDBN4tb5sKSX9vcuNAfLCw2siOHE1IfNUBffymL9x5uAHjBEYuvlwY5J1XMTS5lPMIaPqE8ET37k9iXdODLJpNd82uGmyhV_L603y6g2yeTYlU9vxMW5DrOvyqMq_Q8-LjoCf3o7LT3Bw"
	profile, err := ValidateToken(tokenString)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(profile.Name)
	fmt.Println(profile.Email)
}

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*Profile, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Profile{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		keyString := os.Getenv("PUBLIC_KEY")

		key, err := base64.StdEncoding.DecodeString(keyString)
		if err != nil {
			return nil, err
		}

		pub, err := x509.ParsePKIXPublicKey(key)
		if err != nil {
			return nil, fmt.Errorf("error parsing public key. %v", err)
		}

		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Profile); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error getting token claims")
	}
}