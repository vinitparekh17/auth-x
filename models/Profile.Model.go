package models

type ProfileModel struct {
	UID       int64  `json:"id,omitempty" unique:"true"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile" length:"10" unique:"true"`
	AdminID   int64  `json:"admin_id ,omitempty" unique:"true"`
}
