package marketplace

func IsValidListingStatus(status string) bool {
	validStatuses := [3]string{
		ListingStatusCheckRequired,
		ListingStatusActive,
		ListingStatusSold,
	}

	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}

	return false

}
