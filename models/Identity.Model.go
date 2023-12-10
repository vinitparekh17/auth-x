package models

type IdentityModel struct {
	UID      int64  `json:"id,omitempty" unique:"true"`
	Email    string `json:"email" unique:"true"`
	Password string `json:"password" min:"8"`
}
