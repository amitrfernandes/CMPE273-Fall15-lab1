package main
import (
"fmt"
"github.com/amitrfernandes/hello"
)

func TestFib (t *testing.T){

var v float 64
v = Fib(3)
if v != 2{
t.Error ("Expected 2, got ",v)

}
