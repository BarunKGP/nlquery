package cards

import "github.com/BarunKGP/nlquery/internal/database"

const (
	Active = iota + 1
	Dormant
	Deleted
)

type Cards struct {
	Id     int64
	Query  string
	UserId int64
}

type apiCards struct {
	userId int64
	query  string
}

func DestructureDbCard(dbCard database.Card) *Cards {
	card := Cards{
		Id:     dbCard.ID,
		Query:  dbCard.Query.String,
		UserId: dbCard.ID,
	}

	return &card
}
