package models

type Student struct {
	Name string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student
