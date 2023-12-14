package models

type Body struct {
	List        Result
	ContactList []Contact
}

type Result struct {
	List []string
}
