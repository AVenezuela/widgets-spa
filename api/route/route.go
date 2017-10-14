package route

import (
	"encoding/json"
)

//JSONError type for API errors
type JSONError struct {
    Message string
}

//newJSONError creates a JSONError instance
func newJSONError(message string) *JSONError {
    return &JSONError{
        Message: message,
    }
}

func GetJSONError(errorMessage string) ([]byte){
	jerr := newJSONError(errorMessage)
	if json, err := json.Marshal(jerr); err == nil{
		return json
	}
	panic("Can't create json error")	
}