package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
)

func GenerateJWT(hostname, role string) (string, error) {

	secretkey := os.Getenv("L3AFD_SEC_KEY")
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

	secretkey := os.Getenv("L3AFD_SEC_KEY")

	fmt.Println("Secret key", secretkey)

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey1), nil
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
	//machineHostname, err := os.Hostname()
	//if err != nil {
	//	fmt.Println("hostname failed")
	//}
	//role := "admin"
	//tknStr, err := GenerateJWT(machineHostname, role)
	//if err != nil {
	//	fmt.Println("token failed")
	//}

	tknStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJob3N0bmFtZSI6ImwzYWYtbG9jYWwtdGVzdCIsInJvbGUiOiJhZG1pbiJ9.GzC1bH5fQoPLDS1HxoT5i2PqHnuKcerxeAaSd0Dy7rc"
	fmt.Println("token: ", tknStr)

	fmt.Println("valid token : ", ValidateJWT(tknStr))
}
