package helpers

func GetOffsetByPageLimit(page, limit int) int {
	return (page - 1) * limit
}
