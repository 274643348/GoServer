package engine

type Parse interface {
	Parser(contents []byte,url string)ParseRusult
	Serialize()(name string,args interface{})
}


type Request struct {
	Url   string
	Parse Parse
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

type NilParser struct {

}

func ( NilParser) Parser(_ []byte, _ string) ParseRusult {
	return ParseRusult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}




//////////////////////////////////////////////包装一下parserFunc//////

type ParserFunc  func(contents []byte,url string) ParseRusult

type FuncParse struct{
	parser ParserFunc
	name string
}

func (f *FuncParse) Parser(contents []byte, url string) ParseRusult {
	return f.parser(contents,url)
}

func (f *FuncParse) Serialize() (name string, args interface{}) {
	return f.name,nil
}

func NewFuncParser(p ParserFunc,name string)*FuncParse{
	return &FuncParse{parser:p,name:name}
}


