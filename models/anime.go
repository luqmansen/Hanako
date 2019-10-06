package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
)

type Anime struct {
	gorm.Model
	//ID uint `json:"anime_id" gorm:"primary_key" `
	Name string `json:"name"`
	Genre string `json:"genre"`
	Kind string `json:"kind"`
	Episodes string `json:"episodes"`
	Rating float32 `json:"rating"`
	Members int `json:"members"`
}

func (anime *Anime) Validate() (map[string]interface{},bool){

	//if anime.ID <=0 {
	//	return u.Message(false, "Not Found"),false
	//}
	if anime.Name == "" {
		return u.Message(false, "Name should be on payload"),false
	}
	if anime.Genre == "" {
		return u.Message(false, "Genre should be on payload"),false
	}
	if anime.Kind == "" {
		return u.Message(false, "Kind should be on payload"),false
	}

	return u.Message(true, "success"),true
}

func (anime *Anime) AddEntry() (map[string]interface{}){

	if resp, ok := anime.Validate(); !ok{
		return resp
	}

	getDB().Create(anime)

	resp:= u.Message(true, "success")
	resp["anime"] = anime
	return resp
}

func GetByTitle(title string) ([]*Anime){

	animes := make([]*Anime, 0)
	err := getDB().Table("anime").Where("name ILIKE '%?%'",title).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}

func GetByID(ID string) ([]*Anime){

	animes := make([]*Anime, 0)
	err := getDB().Table("anime").Where("ID = ?",ID).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}

func GetAll() ([]*Anime){

	animes := make([]*Anime, 0)
	err := getDB().Limit(20).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}





