package utilty 


import (
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"os"
	"time"
	"errors"
)


func HashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifiedPassword(hashedPassword string, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}



var secretKey=os.Getenv("secretKey")
var jwtKey = []byte(secretKey)

func GenerateToken(email string,id int64) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = email
	claims["id"]=id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}
func VerifyToken(tokenString string) (int64, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,nil
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if token == nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Check if the "id" claim exists and extract the user ID
	idClaim, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("user ID claim not found or invalid")
	}
	id := int64(idClaim)

	return id, nil
}
