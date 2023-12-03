package models

// Creating user model which suits postgres database
type UserModel struct {
	ID                  int64  `json:"id,omitempty" unique:"true"`
	Name                string `json:"name"`
	Email               string `json:"email" unique:"true"`
	Password            string `json:"password" min:"8"`
	Mobile              string `json:"mobile" length:"10" unique:"true"`
	Role                string `omitempty:"true" json:"role"`
	OTP                 int64  `json:"otp" length:"6"`
	IsVerified          bool   `json:"is_verified"`
	ForgetPasswordToken string `json:"forget_password_token"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}
