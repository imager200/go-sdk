package client

const (
	ApplicationOctetStreamContentType = "application/octet-stream"
	ApplicationJsonContentType        = "application/json"
)

func ToPtr[T any](t T) *T {
	return &t
}
