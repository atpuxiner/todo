package common


const (
	SUCCESS = 0
	ERROR = 1
	UNAUTHORIZED = 2
)

var codeMsg = map[int]string{
	SUCCESS: "succeed.",
	ERROR: "failed!",
	UNAUTHORIZED: "unauthorized!",
}

func CodeMsg(code int) string {
	return codeMsg[code]
}
