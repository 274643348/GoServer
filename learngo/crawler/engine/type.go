package engine

type Request struct {
	Url string
	ParseFunc func([] byte) ParseRusult
}

type ParseRusult struct {
	Requests []Request
	Items []interface{}
}

func NilParser([]byte)ParseRusult{
	return  ParseRusult{}
}
