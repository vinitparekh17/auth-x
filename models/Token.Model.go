package models

/* Token model that describes how our Token table looks like
 	Why is this sapearte from Identity Model?
	Because this token field will be often retrived or updated
	By not including this one column will make identity table's data retrival faster
	It may not be a big deal for small projects but for large projects it is
*/

type TokenModel struct {
	UID         int64  `json:"id,omitempty" unique:"true"`
	ForgotToken string `json:"forgot_token"`
}
