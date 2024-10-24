package Handler

import (
	"net/http"
)

// function to execute the errors
func Error_handle(w http.ResponseWriter, num int) {
	message := http.StatusText(num)
	w.WriteHeader(num)
	error_template.Execute(w, message)
}

// function to check if there's a problem in parsing error and index html files
func CheckParseFile(w http.ResponseWriter, r *http.Request, err1 error, err2 error, err3 error) bool {
	if err2 != nil {
		http.Error(w, http.StatusText(500), 500)
		return false
	}
	if err1 != nil || err3 != nil {
		Error_handle(w, 500)
		return false
	}
	return true
}
