package usercontroller

import (
	"net/http"
	"reflect"
	"time"

	"github.com/AhmadHafidz1316/goAPI/config"
	"github.com/AhmadHafidz1316/goAPI/helpers"
	"github.com/AhmadHafidz1316/goAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ShowAll(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
	}
	user.Password = hashedPassword

	v := reflect.ValueOf(user)
	typeOfUser := v.Type()
	for i := 0 ; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.Interface() == "" {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message" : typeOfUser.Field(i).Name + " Kosong"})
			return
		}
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"User": user})
}

func Login(c *gin.Context) {

	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Username atau Password Salah"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Username atau Password Salah"})
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "GO REST API",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})

}
