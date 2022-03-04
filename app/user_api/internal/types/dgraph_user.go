package types

type User struct {
	Phone     int    `json:"phone"`
	Nickname  string `json:"nickname"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
	Avatar    string `json:"avatar"`
	Address   string `json:"address"`
	PublicKey string `json:"public_key"`
	Password  string `json:"password"`
}
