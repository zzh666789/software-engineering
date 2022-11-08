package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Countcode struct {
	gorm.Model
	Operate1 int    `json:"operate1" uri:"operate1" binding:"required"`
	Operate2 int    `json:"operate2" uri:"operate2" binding:"required"`
	Code     string `json:"code" uri:"code" binding:"required"`
	Answer   int    `json:"answer" uri:"answer" binding:"required"`
}

type Countcode1 struct {
	gorm.Model
	Operate1 int    `json:"operate1" uri:"operate1" binding:"required"`
	Operate2 int    `json:"operate2" uri:"operate2" binding:"required"`
	Operate3 int    `json:"operate3" uri:"operate3" binding:"required"`
	Code1    string `json:"code1" uri:"code1" binding:"required"`
	Code2    string `json:"code2" uri:"code2" binding:"required"`
	Answer   int    `json:"answer" uri:"answer" binding:"required"`
}

type Countcode2 struct {
	gorm.Model
	Operate1 int    `json:"operate1" uri:"operate1" binding:"required"`
	Operate2 int    `json:"operate2" uri:"operate2" binding:"required"`
	Operate3 int    `json:"operate3" uri:"operate3" binding:"required"`
	Operate4 int    `json:"operate4" uri:"operate4" binding:"required"`
	Code1    string `json:"code1" uri:"code1" binding:"required"`
	Code2    string `json:"code2" uri:"code2" binding:"required"`
	Code3    string `json:"code3" uri:"code3" binding:"required"`
	Answer   int    `json:"answer" uri:"answer" binding:"required"`
}
