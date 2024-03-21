package domain

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"PrimaryKey"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email" gorm:"index;unique;not null"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Code       int       `json:"code"`
	ExpiryTime time.Time `json:"expiry_time"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`
	UserRole   string    `json:"user_role"  gorm:"default:buyer"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
