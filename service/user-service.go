package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

//AuthenticateUser ...
func AuthenticateUser(context *gin.Context) interface{} {
	db := *config.GetConnection()
	var (
		user   models.Users
		userDB models.Users
		result gin.H
	)

	if err := context.Bind(&user); err != nil {
		log.Println("Data tidak ada, error : ", err.Error())
	}

	db.Where("email = ?", user.Email).First(&userDB)

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		result = gin.H{
			"message": "email atau password salah",
		}
	} else {
		type authCustomClaims struct {
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			userDB.Email,
			userDB.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(utils.GetEnv("SECRET_TOKEN")))
		if err != nil {
			log.Println("Gagal create token, message ", err.Error())
			result = gin.H{
				"message": "Gagal create token",
				"token":   nil,
			}
		} else {
			log.Println("Email ", user.Email, " Berhasil login")
			result = gin.H{
				"message": "anda berhasil login",
				"token":   token,
			}
		}
	}

	return result
}

//GetListUsers func
func GetListUsers(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		users []models.Users
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	paginator := utils.Paging(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)

	return helper.ResponseSuccessList(paginator)
}

//GetUserByID func
func GetUserByID(id interface{}) interface{} {
	db := *config.GetConnection()

	var (
		user models.Users
	)

	result := db.First(&user, id)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&user)
}

//AddUser func
func AddUser(context *gin.Context) interface{} {
	db := *config.GetConnection()

	user := models.Users{
		Email: context.PostForm("name"),
		Password: context.PostForm("password"),
		Role: context.PostForm("role"),
	}

	result := db.Create(&user)

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessSingle(&user)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdateUser func
func UpdateUser(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		user models.Users
	)

	result := db.First(&user, context.Param("id"))

	if result.RowsAffected > 0 {
		user.Email = context.PostForm("email")
		user.Password = context.PostForm("password")
		user.Role = context.PostForm("role")

		result := db.Save(&user)
		if result.RowsAffected > 0 {
			return helper.ResponseSuccessSingle(&user)
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeleteUser func
func DeleteUser(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.Users{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data user")
	}

	return helper.ResponseFailed(result.Error)
}