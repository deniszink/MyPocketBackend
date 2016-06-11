package services

import (
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"backend/models"
	"backend/core/authentication"

	"backend/core/store"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type LoginResponse struct {
	*models.User `json:"user"`
	*models.Token
	IncomeCategories []models.Category `json:"incomeCategories"`
	ExpenseCategories []models.Category `json:"expenseCategories"`
}

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	mongo := store.ConnectMongo()
	user := &models.User{}
	//check is user exists
	err := mongo.GetOne(store.TableUsers, bson.M{"email":requestUser.Email, "password":requestUser.Password}, user);

	if err != nil {
		response, _ := json.Marshal(models.Error{
			Error: "Password and/or email were incorrect, please try again",
		})
		return http.StatusBadRequest, []byte(response)
	}

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(string(requestUser.Id))
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			var incomeCategories []models.Category
			var expenseCategories []models.Category
			var all []models.Category

			mongo.FindAll(store.TableCategories,bson.M{},&all)

			for _, value := range all{
				if value.Type == "income"{
					incomeCategories = append(incomeCategories,value)
				}else{
					expenseCategories = append(expenseCategories,value)
				}
			}
			response, _ := json.Marshal(LoginResponse{
				user,
				&models.Token{token},
				incomeCategories,
				expenseCategories,
			})
			fmt.Println(string(response))
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(string(requestUser.Id))
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(models.Token{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}