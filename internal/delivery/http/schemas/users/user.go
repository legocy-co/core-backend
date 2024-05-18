package users

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	log "github.com/sirupsen/logrus"
)

type UserRegistrationResponse struct {
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Username string `json:"username"`
}

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ur *UserRegistrationRequest) ToUser() *models.User {
	return &models.User{
		Email:    ur.Email,
		Username: ur.Username,
		Role:     models.USER,
	}
}

func (ur *UserRegistrationRequest) ToAdmin() *models.User {
	return &models.User{
		Email:    ur.Email,
		Username: ur.Username,
		Role:     models.ADMIN,
	}
}

type UserDetailResponse struct {
	ID           int                       `json:"id"`
	Username     string                    `json:"username"`
	Email        string                    `json:"email"`
	Role         int                       `json:"role"`
	Images       []UserImageInfoResponse   `json:"images"`
	ReviewTotals *UserReviewTotalsResponse `json:"reviewTotals"`
}

func (r *UserDetailResponse) WithReviewTotals(totals *UserReviewTotalsResponse) *UserDetailResponse {
	r.ReviewTotals = totals
	return r
}

func GetUserDetailResponse(u *models.User) UserDetailResponse {

	images := make([]UserImageInfoResponse, 0, len(u.Images))
	log.Debugf("User images length: %d", len(u.Images))
	for _, img := range u.Images {
		images = append(images, GetUserImageResponse(img))
	}

	return UserDetailResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
		Images:   images,
	}
}

type UserUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *UserUpdateRequest) ToUserValueObject() *models.UserValueObject {
	return &models.UserValueObject{
		Email:    r.Email,
		Username: r.Username,
	}
}

type GoogleSignInUpRequest struct {
	Token string `json:"token"`
}
