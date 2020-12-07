package shape

type Line struct {
	Shape
}

//func NewLine(X int, Y int) *Line {
//	return &Line{
//		Shape{X: X, Y: Y},
//	}
//}
//
//func (s *Line) Draw(window draw.Window) {
//	window.DrawRect(s.X, s.Y, 10, 50, draw.Blue)
//}
//
//func (s *Line) Move(window *draw.Window) bool {
//	//time.Sleep(1000 * time.Microsecond)
//	s.Y += 10
//	if s.Y >= 480 - 50 {
//		s.Y-=10
//		(*window).FillRect(s.X, s.Y, 10, 50, draw.Green)
//		return false
//	}
//	(*window).FillRect(s.X, s.Y, 10, 50, draw.Green)
//	return true
//
//}
