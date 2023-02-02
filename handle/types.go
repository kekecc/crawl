package handle

type ParseRes struct {
	Requests []Request
	Contents []interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseRes
}
