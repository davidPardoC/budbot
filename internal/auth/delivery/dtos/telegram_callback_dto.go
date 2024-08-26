package dtos

type TelegramCallbackDto struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	AuthDate  string `json:"auth_date"`
	Hash      string `json:"hash"`
}
