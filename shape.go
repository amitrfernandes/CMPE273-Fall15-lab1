func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

// interface which makes
// both Circle and Rectangle to behave like Shape
// since they implement the area() method
type Shape interface{
  area() float64
  perimeter() float64
}

type Circle struct{
  x, y, r float64
}

// method to circle struct
func (c *Circle) area() float64{
  return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64{
  return 2 * math.Pi * c.r
}

type Rectangle struct{
  x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64{
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64{
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return (2 * l) + (2 * w)
}


type MultiShape struct{
  shapes []Shape
}

func (m *MultiShape) area() float64{
  var area float64
  for _, s := range m.shapes{
    area += s.area()
  }
  return area
}




func main(){
  // c := Circle{0, 0, 5}
  c := Circle{x: 0, y: 0, r: 5}
  fmt.Println(c.area())
fmt.Println(c.perimeter())

  r := Rectangle{0,0,10,10}
  fmt.Println(r.area())
fmt.Println(r.perimeter())



}
