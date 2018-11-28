package persist

import "testing"
import "../moder"
import "../persist"

func TestSaver(t *testing.T) {
	profile := moder.Profile{
		Name :"小红",
		Gender :"女",
		Age : "11",
		Height :"11",
		Weight :"11",
		Income :"11",
		Marrige :"11",
		Education :"11",
		Occupation :"11",
		Hukou :"11",
		Xingzuo :"11",
		House :"11",
		Car :"11",
	}
	persist.Save(profile)
}