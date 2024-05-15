package middlewares

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type CookieData struct {
	Value   string
	Expires time.Time
}

// var secretKey = "a3b762f871cdb3bae0044c649622fc1396eda3e3"
var secretKey = `{
	"e": "AQAB",
	"kty": "RSA",
	"alg": "RS256",
	"use": "sig",
	"kid": "a3b762f871cdb3bae0044c649622fc1396eda3e3",
	"n": "uBHF-esPKiNlFaAvpdpejD4vpONW9FL0rgLDg1z8Q-x_CiHCvJCpiSehD41zmDOhzXP_fbMMSGpGL7R3duiz01nK5r_YmRw3RXeB0kcS7Z9H8MN6IJcde9MWbqkMabCDduFgdr6gvH0QbTipLB1qJK_oI_IBfRgjk6G0bGrKz3PniQw5TZ92r0u1LM-1XdBIb3aTYTGDW9KlOsrTTuKq0nj-anW5TXhecuxqSveFM4Hwlw7pw34ydBunFjFWDx4VVJqGNSqWCfcERxOulizIFruZIHJGkgunZnB4DF7mCZOttx2dwT9j7s3GfLJf0xoGumqpOMvecuipfTPeIdAzcQ"
}`

func CheckCookie(c echo.Context, jwtToken string) uint {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return 0
	}

	token, _, err := new(jwt.Parser).ParseUnverified(cookie.Value, jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return 0
	}

	// Extract header data from the token
	headerData := token.Claims.(jwt.MapClaims)

	// Print the header data
	fmt.Println("Header data:")
	fmt.Println(headerData["email"])
	return 1
}

func GetExp(jwtToken string) (time.Time, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return time.Now(), err
	}

	// Extract header data from the token
	bodyData := token.Claims.(jwt.MapClaims)
	fmt.Println("exp data: ", bodyData["exp"])
	exp := time.Unix(int64(bodyData["exp"].(float64)), 0)
	fmt.Println("exp time data: ", exp)
	fmt.Println(time.Now().Unix())

	return exp, nil
}
