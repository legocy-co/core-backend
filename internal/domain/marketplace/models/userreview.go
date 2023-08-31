package marketplace

import (
	"errors"
	"legocy-go/internal/domain/users/models"
)

type UserReview struct {
	ID       int
	Seller   auth.User
	Reviewer auth.User
	Rating   int
	Message  string
	Date     string
}

type UserReviewValueObject struct {
	SellerID   int
	ReviewerID int
	Rating     int
	Message    string
	Date       string
}

func NewUserReview(ID int, Seller auth.User, Reviewer auth.User, Rating int, Message string, Date string) (*UserReview, error) {
	if Rating > 5 || Rating < 1 {
		err := errors.New("wrong Rating value (1-5)")
		return nil, err
	}

	if len(Message) == 0 {
		err := errors.New("empty message")
		return nil, err
	}

	if Seller == Reviewer {
		err := errors.New("wrong Seller (not Reviewer)")
		return nil, err
	}

	return &UserReview{
		ID:       ID,
		Seller:   Seller,
		Reviewer: Reviewer,
		Rating:   Rating,
		Message:  Message,
		Date:     Date,
	}, nil
}

func NewUserReviewValueObject(SellerID int, ReviewerID int, Rating int, Message string, Date string) (*UserReviewValueObject, error) {
	if Rating > 5 || Rating < 1 {
		err := errors.New("wrong Rating value (1-5)")
		return nil, err
	}

	if len(Message) == 0 {
		err := errors.New("empty message")
		return nil, err
	}

	if SellerID == ReviewerID {
		err := errors.New("wrong Seller (not Reviewer)")
		return nil, err
	}

	return &UserReviewValueObject{
		SellerID:   SellerID,
		ReviewerID: ReviewerID,
		Rating:     Rating,
		Message:    Message,
		Date:       Date,
	}, nil
}

/*type UserReviewValueObject struct {
	SellerID   int
	ReviewerID int
	Rating     int
	Message    string
	Date       string
}*/
