package interpol

type Issue struct {
	Locale  string
	Message string
}

type Result struct {
	Errors []Issue
}
