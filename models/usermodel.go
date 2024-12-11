package models

type User struct {
	User_name    string `json:"user_name"`
	RealName     string `json:"realname"`
	Greetings    string `json:"greetings"`
	LastLogin    string `json:"last_login"`
	Email        string `json:"email"`
	Base64qrcode string `json:"base64qrcode"`
	HP           string `json:"hp"`
	Bucketcoll   string `json:"bucketcoll"`
}

type UserLogin struct {
	User_name    string `json:"user_name"`
	UserPassword string `json:"password"`
	RealName     string `json:"realname"`
	Greetings    string `json:"greetings"`
	LastLogin    string `json:"last_login"`
	Email        string `json:"email"`
	Base64qrcode string `json:"base64qrcode"`
	HP           string `json:"hp"`
	Bucketcoll   string `json:"bucketcoll"`
	Token        string `json:"token"`
}

type Login struct {
	Username string `json:"username" example:"1879"`
	Password string `json:"password" example:"pass,123"`
}

type MyProfile struct {
	ID   string `json:"username"`
	Name string `json:"real_name"`
}
