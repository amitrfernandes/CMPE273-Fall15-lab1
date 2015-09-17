package main

import
(
"fmt"
"Math"
"github.com/amitrfernandes/shape"
)
func shapeTest(t *testing.T)
{
 c := Circle{x: 0, y: 0, r: 5}
  fmt.Println(c.area())
if{
c.perimeter!=31.4
t.Error("Error!")
}

  r := Rectangle{0,0,10,10}
  fmt.Println(r.area())
if{
r.perimeter!=40
t.Error("Error!")
}
