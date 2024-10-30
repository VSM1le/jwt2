package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	"github.com/VSM1le/jwt2/repositorys"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	jwt.StandardClaims
}

func GenerateAllTokens(id int64, email string, firstName string, lastName string) (signedToken string, signedRefreshToken string, err error) {
	// Access the secret key from environment variables
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(secretKey) == 0 {
		return "", "", fmt.Errorf("SECRET_KEY is not set")
	}

	// Create claims for the access token
	claims := &SignedDetails{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	// Create claims for the refresh token
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(), // Refresh token expires in 7 days
		},
	}

	// Sign the access token with HS256
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		log.Panic("Error signing token: ", err)
		return "", "", err
	}

	// Sign the refresh token with HS256
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		log.Panic("Error signing refresh token: ", err)
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)

	if err != nil {
		msg = err.Error()
		return nil, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Sprintf("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Sprintf("the token has expired")
	}
	return claims, msg
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, user *models.User) error {
	// Establish database connection
	db, err := database.DBinstance()
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	// Create a new repository instance
	repositoryNew := repositorys.NewPostgreSQLRepository(db)

	// Get the current time
	updatedAt := time.Now()

	// Update tokens in the database
	err = repositoryNew.UpdateToken(user.ID, signedToken, signedRefreshToken, updatedAt)
	if err != nil {
		return fmt.Errorf("error updating tokens in the database: %w", err)
	}
	defer db.Close()
	// Update user struct fields with new token and refresh token
	user.Token = &signedToken               // Assuming Token is a *string in the User struct
	user.RefreshToken = &signedRefreshToken // Assuming RefreshToken is a *string in the User struct
	user.UpdatedAt = &updatedAt             // Assuming UpdatedAt is a *time.Time in the User struct

	return nil
}
