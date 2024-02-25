package constant

type ResponseStatus int
type Headers int
type General int

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnkonwnError
	InvalidRequest
	Unauthorized
)

func (rs ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[rs-1]
}

func (rs ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data not found", "Unknown error", "Invalid request", "Unauthorized"}[rs-1]
}
