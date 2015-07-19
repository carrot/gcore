package response

// Custom error codes
const (
	ErrorNoError           = 0
	ErrorInvalidParameters = 1
)

var errorText = map[int]string{
	ErrorNoError:           "No Error",
	ErrorInvalidParameters: "Invalid Parameters",
}

// Returns the associated error text
func GetErrorText(code int) string {
	return errorText[code]
}
