package users

import (
	"os"
	"time"
	"hotel_management/database"
	"hotel_management/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func prepareToken(user *User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, _ := jwtToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return token
}

func prepareResponse(user *User, account ResponseAccount, withToken bool) map[string]interface{} {
	responseUser := &ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Account: account,
	}
	var response = map[string]interface{}{"message": "all is fine", "status_code": 200}
	if withToken {
		var token = prepareToken(user);
		response["jwt"] = token
	}
	response["data"] = responseUser
	return response
}

func AuthenticateUser(username, password string) map[string]interface{} {
	user := &User{}
	query := database.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return utils.HandleResponse("Wrong username", 400)
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return utils.HandleResponse("Wrong password", 400)
	}

	account := ResponseAccount{}
	database.DB.Table("accounts").Select("id, first_name, last_name, birth_date").Where("user_id = ?", user.ID).First(&account)

	response := prepareResponse(user, account, true)
	return response
}

func UserRegistration(data *SignupUserSerializer) map[string]interface{} {
	generatePassword := utils.HashAndSalt([]byte(data.Password))
	user := &User{
		Username: data.Username, 
		Password: generatePassword, 
		Email: data.Email}
	database.DB.Create(&user)

	account := &Account{
		FirstName: data.Account.FirstName, 
		LastName: data.Account.LastName, 
		BirthDate: data.Account.BirthDate, 
		UserID: user.ID}
	database.DB.Create(&account)

	responseAccount := ResponseAccount{
		ID: account.ID, FirstName: 
		account.FirstName, LastName: 
		account.LastName, 
		BirthDate: account.BirthDate}
	response := prepareResponse(user, responseAccount, true)
	
	return response
}

func GetUser(jwtToken string) map[string]interface{} {
	isValid, tokenData := utils.ValidateToken(jwtToken)
	if isValid {
		user := &User{}
		query := database.DB.Where("id = ?", tokenData["user_id"]).First(&user)
		if query.Error != nil {
			return utils.HandleResponse("Not Found", 404)
		}

		account := ResponseAccount{}
		database.DB.Table("accounts").Select("id, first_name, last_name, birth_date").Where("user_id = ?", user.ID).First(&account)

		response := prepareResponse(user, account, false)
		return response
	} else {
		return utils.HandleResponse("Invalid Token", 400)
	}
} 