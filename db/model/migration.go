package model

var Migrations []Model // slice of model

type Model interface {
	TableName() string
}
