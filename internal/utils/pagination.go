package utils

import "strconv"

type Pagination struct {
	Page  string
	Limit string
}

func LoadPaginationConfig(p Pagination, page, limit *int) {
	if p.Limit == "" {
		return
	}

	l, err := strconv.Atoi(p.Limit)
	if err == nil {
		*limit = l
		if p.Page != "" {
			p, err := strconv.Atoi(p.Page)
			if err == nil {
				*page = p
			}
		}
	}

}
