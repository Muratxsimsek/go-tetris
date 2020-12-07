package shape

type LShape struct {
	Shape
}

func NewLShape(X int, Y int) *LShape {
	return &LShape{
		Shape{X: X, Y: Y},
	}
}

//func (s *LShape) Draw(window draw.Window) {
//	window.DrawRect(s.X, s.Y, 10, 50, draw.Blue)
//}
//
//func (s *LShape) Move(window *draw.Window) bool {
//	//time.Sleep(1000 * time.Microsecond)
//	s.Y += 10
//	if s.Y >= 480 - 50 {
//		s.Y-=10
//		(*window).FillRect(s.X, s.Y, 10, 40, draw.Yellow)
//		(*window).FillRect(s.X, s.Y + 40, 30, 10, draw.Yellow)
//		return false
//	}
//	(*window).FillRect(s.X, s.Y, 10, 40, draw.Yellow)
//	(*window).FillRect(s.X, s.Y + 40, 30, 10, draw.Yellow)
//
//	return true
//
//}
