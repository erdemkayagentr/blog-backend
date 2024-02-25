package entities

import "github.com/google/uuid"

type Category struct {
	Id    uuid.UUID `column:"id";json:"id";primaryKey:"true";default:"uuid_generate_v4()"`
	Name  string    `column:"name";json:"name"`
	Posts []Post    `json:"posts"`
}
