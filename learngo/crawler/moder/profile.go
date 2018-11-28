package moder

import "encoding/json"

type Profile struct {
	Url string
	Id string
	Name string
	Gender string
	Age string
	Height string
	Weight string
	Income string
	Marrige string
	Education string
	Occupation string
	Hukou string
	Xingzuo string
	House string
	Car string
}

func FromJsonObj(o interface{})(Profile,error){
	var profile Profile
	s,err := json.Marshal(o)
	if err != nil {
		return profile,err
	}

	err = json.Unmarshal(s,&profile)
	return profile,err
}