package db

type DataPuller interface {
	PullUsers() Users
}

type Storage struct {
}

// TODO: Change this because we are using a global variable "UsersCollection" in the app.
func InitStorage() {
	JsonUnmarshal()
}

// TODO: Delete this because I can't index in the user field from Storage.
func NewStorage() *Storage {
	return &Storage{}
}

func PullUsers() *Users {
	return &UsersCollection
}
