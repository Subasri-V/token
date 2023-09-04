package models

type UserSignIn struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type DBResponse struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}