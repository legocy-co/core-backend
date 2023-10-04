package models

import auth "legocy-go/internal/domain/users/models"

type LegoCollection struct {
	User auth.User
	Sets []CollectionLegoSet
}

func (lc LegoCollection) TotalSets() int {
	return len(lc.Sets)
}
