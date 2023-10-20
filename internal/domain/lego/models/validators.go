package lego

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
