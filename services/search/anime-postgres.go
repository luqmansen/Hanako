package search

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/luqmansen/Hanako/services/user"
	u "github.com/luqmansen/hanako/api/utils"
	"net/http"
	"strconv"
)

type Anime2 struct {
	gorm.Model
	//ID uint `json:"anime_id" gorm:"primary_key" `
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Type     string  `json:"type"`
	Episodes string  `json:"episodes"`
	Rating   float32 `json:"rating"`
	Members  int     `json:"members"`
}

func (anime *Anime2) Validate() (map[string]interface{}, bool) {

	if anime.Name == "" {
		return u.Message(http.StatusBadRequest, "Name should be on payload"), false
	}
	if anime.Genre == "" {
		return u.Message(http.StatusBadRequest, "Genre should be on payload"), false
	}
	if anime.Type == "" {
		return u.Message(http.StatusBadRequest, "Type should be on payload"), false
	}

	return u.Message(http.StatusOK, "success"), true
}

func (anime *Anime2) AddEntry() map[string]interface{} {

	if resp, ok := anime.Validate(); !ok {
		return resp
	}

	user.getDB().Create(anime)

	resp := u.Message(http.StatusOK, "success")
	resp["anime"] = anime
	return resp
}

func GetAll(number string) []*Anime2 {

	if number == "" {
		number = "20"
	} else if _, err := strconv.Atoi(number); err != nil {
		number = "20"
	}
	animes := make([]*Anime2, 0)
	err := user.getDB().Table("anime2").Limit(number).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(animes) == 0 {
		fmt.Println("0 record found")
	}
	return animes
}

func GetByTitle(title string) []*Anime2 {

	animes := make([]*Anime2, 0)
	err := user.getDB().Table("animes").Where("name ILIKE '%' || ? || '%'", title).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}

func GetByID(ID uint) *Anime2 {

	animes := &Anime2{}
	err := user.getDB().Table("animes").Where("ID = ?", ID).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}
