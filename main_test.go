package main 
import( 
	 "testing"
)

func TestHelloWorld(T *testing.T){
	expected:= "Hello, World\n"
	if got:= getMessage(); got !=expected{
		T.Errorf("getMessage %v is not eqaul to expected %v", got, expected)
	}
}

func getMessage() string{
	return "Hello, World\n" }
