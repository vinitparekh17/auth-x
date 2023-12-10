package models

// Identity struct that describes how our Identity table looks like

type UserModel struct {
	ID        int64  `json:"id,omitempty" unique:"true"`
	Name      string `json:"name"`
	Email     string `json:"email" unique:"true"`
	Password  string `json:"password" min:"8"`
	Mobile    string `json:"mobile" length:"10" unique:"true"`
	Role      string `omitempty:"true" json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
