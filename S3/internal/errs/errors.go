package errs

import (
	"encoding/xml"
	"net/http"
)

// Error structure in order to save and retrive the error-messages
type Error struct {
	Code     int    `xml:"Code"`
	Message  string `xml:"Message"`
	Resource string `xml:"Resource"`
}

// Inner method for retriving the error's information
func (e *Error) WriteError(w http.ResponseWriter) {
	x, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(e.Code)
	w.Write(x)
}
