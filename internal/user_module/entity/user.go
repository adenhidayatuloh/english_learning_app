package entity

import (
	"english_app/pkg/errs"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = "SangatRahasia"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	Role      string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Println("Password doesn't match:", err)
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    u.ID,
			"email": u.Email,
			"role":  u.Role,
		})

	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		fmt.Print("Hallo ", err)
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mapClaims jwt.MapClaims

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {

		return errs.NewUnauthenticated("Token is not valid")
	}
	mapClaims = claims

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	id, okForId := claim["id"].(string)
	email, okForEmail := claim["email"].(string)
	role, okForRole := claim["role"].(string)

	if !okForId {
		return errs.NewUnauthenticated("Token doesn't contains id")
	}

	if !okForEmail {
		return errs.NewUnauthenticated("Token doesn't contains email")
	}

	if !okForRole {
		return errs.NewUnauthenticated("Token doesn't contains role")
	}
	uuid, err := uuid.Parse(id)

	if err != nil {
		return errs.NewBadRequest("ID invalid")
	}
	u.ID = uuid
	u.Email = email
	u.Role = role

	return nil
}
