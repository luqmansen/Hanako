package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}


func (account *Account) Validate() (map[string] interface{}, bool){

	if !strings.Contains(account.Email, "@"){
		return u.Message(false, "Email Address is Required"),false
	}

	if len(account.Password) <6{
		return u.Message(false, "Password is Required"),false
	}

	temp := &Account{}

//	Check for duplicate email
	err:= getDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return u.Message(false, "Connection error, Please Retry"),false
	}

	if temp.Email != ""{
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true

}

func (account *Account) Create() (map[string] interface{})  {

	if resp, ok := account.Validate(); !ok{
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	getDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed To Create Account, connection Error.")
	}

//	Crete new jwt Token for newly created account
	tk := &Token{UserId:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //Delete pass from memory

	response := u.Message(true , "Account Has Been Created")
	response["account"] = account
	return response

}

func Login(email, password string)  (map[string] interface{}){

	account := &Account{}
	err := getDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return u.Message(false, "Email Address Not Fond")
		}
		return u.Message(false, "Connection Error, Please Retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil  && err == bcrypt.ErrMismatchedHashAndPassword{
		return u.Message(false, "Invalid Login Credential")
	}

	//password work, and delete password immediately
	account.Password = ""

	//Create Jwt Token
	tk:= &Token{UserId:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS356"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint)  *Account{

	acc := &Account{}
	getDB().Table("account").Where("id = ?", u).First(acc)
	if acc.Email == ""{
		return nil
	}
	acc.Password =""
	return acc
}




























