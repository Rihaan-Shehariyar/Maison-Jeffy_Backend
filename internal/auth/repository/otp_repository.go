package repository

import (
	"backend/internal/auth/entity"

	"gorm.io/gorm"
)

type OtpRepository interface {
	Create(otp *entity.OTP) error
	FindLatestByEmail(email string) (*entity.OTP, error)
	MarkUsed(id uint) error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OtpRepository {
	return &otpRepository{db}
}

func (r *otpRepository) Create(otp *entity.OTP) error {
	return r.db.Create(otp).Error
}

func (r *otpRepository) FindLatestByEmail(email string) (*entity.OTP, error) {

	var otp entity.OTP

	err := r.db.Where("email = ? AND used=false", email).
		Order("created_at DESC").
		First(&otp).Error
	return &otp, err

}

func (r *otpRepository) MarkUsed(id uint) error {

	return r.db.Model(&entity.OTP{}).
		Where("id = ?", id).
		Update("used", true).Error

}
