package response

import "fmt"

func CreateErrorResponse(message string, err *error) map[string]string {
	ret := map[string]string{
		"message": message,
	}
	if err != nil {
		ret["error"] = fmt.Sprintf("%v", *err)
	}
	return ret
}
