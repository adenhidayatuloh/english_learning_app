package entity

import (
	"time"
)

type Otp struct {
	Email     string    `gorm:"type:varchar(255);not null;primaryKey"` // Email terkait
	Code      string    `gorm:"type:varchar(6);not null"`              // Kode OTP
	ExpiresAt time.Time `gorm:"not null"`                              // Waktu kadaluarsa OTP
	CreatedAt time.Time `gorm:"autoCreateTime"`                        // Waktu pembuatan OTP
}

func (Otp) TableName() string {
	return "auth.otps"
}
