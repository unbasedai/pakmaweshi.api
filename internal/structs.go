package internal


type Product struct {
	Id				string 					`json:"id" bson:"id"`

	Name 			string					`json:"name" bson:"name"`										
	Price 			string					`json:"price" bson:"price"`
	Location 		string					`json:"location" bson:"location"`
	Contact 		string					`json:"contact" bson:"contact"`
	
	Avatar 			string					`json:"avatar" bson:"avatar"`
	Username 		string					`json:"username" bson:"username"`  // username
	UserId			string					`json:"user_id" bson:"user_id"`	 // user id

	Images 			[]string				`json:"images" bson:"images"` // url of the images
}

type User struct {
	Id 				string 					`json:"id" bson:"id"` // user id 

	Avatar			string 					`json:"avatar" bson:"avatar"` // url of the avatar image file

	Name 			string					`json:"name" bson:"name"` // full name
	Number 			string 					`json:"number" bson:"number"` // phone number only +92
	Username		string 					`json:"username" bson:"username"` // username

	Email			string 					`json:"email" bson:"email"` // email
	Password		string 					`json:"password" bson:"password"`	 // password
}