package models

type Body struct {
	MouseDiv string
	Search   string
	List     Result
}

type Result struct {
	List []string
}
