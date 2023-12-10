package models

type TokenModel struct {
	UID         int64  `json:"id,omitempty" unique:"true"`
	ForgotToken string `json:"forgot_token"`
}
