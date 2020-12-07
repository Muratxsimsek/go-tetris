package shape

//import "github.com/gonutz/prototype/draw"

type Square struct {
	Shape
}

func NewSquare(X int, Y int) *Square {
	return &Square{
		Shape{X: X, Y: Y},
	}
}

//func (s *Square) Draw(window draw.Window) {
//	window.DrawRect(s.X, s.Y, 50, 50, draw.Blue)
//}
//
//func (s *Square) Move(window *draw.Window) bool {
//	//time.Sleep(1000 * time.Microsecond)
//	s.Y += 10
//	if s.Y >= 480 - 50 {
//		s.Y-=10
//		(*window).FillRect(s.X, s.Y, 50, 50, draw.Blue)
//		return false
//	}
//	(*window).FillRect(s.X, s.Y, 50, 50, draw.Blue)
//	return true
//}
