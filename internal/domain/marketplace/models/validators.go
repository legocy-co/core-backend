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

func IsValidSetState(setState string) bool {
	validStates := [6]string{
		SetStateBrandNew,
		SetStateBoxOpened,
		SetStateBagsOpened,
		SetStateBuiltWithBox,
		SetStateBuiltWithoutBox,
		SetStateBuiltPiecesLost,
	}

	for _, validState := range validStates {
		if setState == validState {
			return true
		}
	}

	return false
}
