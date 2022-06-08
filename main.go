package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
)

var secretkey = "l3afd"

func GenerateJWT(hostname, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["hostname"] = hostname
	claims["role"] = role

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tknStr string) bool {
	machineHostname, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname failed")
	}
	role := "admin"
	// Initialize a new instance of `Claims`
	claims := &jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("ErrSignatureInvalid")
			return false
		}
		fmt.Println("Err : ", err)
		return false
	}
	if !tkn.Valid {
		fmt.Println("StatusUnauthorized     ")
		return false
	}

	if machineHostname == (*claims)["hostname"] && role == (*claims)["role"] {
		fmt.Println("valid token, authorized")
	}

	fmt.Println("tkn :", tkn)
	fmt.Println(" claims : ", claims)
	return true
}

func main() {
	machineHostname, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname failed")
	}
	role := "admin"
	tknStr, err := GenerateJWT(machineHostname, role)
	if err != nil {
		fmt.Println("token failed")
	}

	fmt.Println("token: ", tknStr)

	fmt.Println("valid token : ", ValidateJWT(tknStr))
}
