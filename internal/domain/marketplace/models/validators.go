package marketplace

import lego "legocy-go/internal/domain/lego/models"

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
		lego.SetStateBrandNew,
		lego.SetStateBoxOpened,
		lego.SetStateBagsOpened,
		lego.SetStateBuiltWithBox,
		lego.SetStateBuiltWithoutBox,
		lego.SetStateBuiltPiecesLost,
	}

	for _, validState := range validStates {
		if setState == validState {
			return true
		}
	}

	return false
}
