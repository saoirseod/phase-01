package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//When a secret key is used with distributed systems it is best to have a dedicated service to generate and verify tokens
//Here, I'll be configuring the secret key into all services where it is used as an example,
//however, having the secret key stated in various locations
//increases the risk of it being compromised in real world applications
var KeyForSigning = []byte(os.Getenv("SECRET_KEY"))

//HS256 involves a secret key that is shared between two parties (not really used here but could be later implemented)
//It's used to encrypt the data and when its used, on the recieving end, it decrypts the data
func createJWT() (string, error) {
	//Could change to SigningMethodEd25519
	//HS256 algorithm, which is short for HMAC-SHA256
	token := jwt.New(jwt.SigningMethodHS256)

	//JWT claims are pieces of information asserted about a subject
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	//name: The name of the user authenticating
	claims["name"] = "Saoirse"
	//aud (audience): Recipient for which the JWT is intended
	claims["aud"] = "testing"
	//iss (issuer): Issuer of the JWT
	claims["iss"] = "Saoirseo"
	//exp (expiration): Expiration of token, this one is valid for 5 minutes
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(KeyForSigning)

	if err != nil {
		fmt.Errorf("ERROR! There was an issue generating your JWT token: %s", err.Error())
		return "", err
	}
	fmt.Println("Generated token: ")
	fmt.Println(tokenString)
	return tokenString, nil
}

func main() {
	createJWT()
}
