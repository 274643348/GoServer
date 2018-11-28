package engine

type Request struct {
	Url string
	ParseFunc func([] byte) ParseRusult
}

type ParseRusult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url string
	Type string
	Id string
	Payload interface{}
}

func NilParser([]byte)ParseRusult{
	return  ParseRusult{}
}
