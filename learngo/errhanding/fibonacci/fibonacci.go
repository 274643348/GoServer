package fibonacci

type intGen func()int
func Fibonacci() intGen{
	a,b:=0,1
	return func() int{
		a,b=b,a+b
		return  a
	}
}
