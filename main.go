package main

//https://github.com/gonutz/prototype
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"log"
	"math/rand"
	"runtime"
	"tetris/shape"
	"time"
)

var ishapes []shape.IShape
var ishapeCount int
var gameOver bool

//var dWindow *draw.Window

func longRunningTask() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// Simulate a workload.
		for {
			time.Sleep(time.Second * 3)
		}
		//time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

//func main() {
//	aCh, bCh, cCh := longRunningTask(), longRunningTask(), longRunningTask()
//	a, b, c := <-aCh, <-bCh, <-cCh
//
//	fmt.Println(a, b, c)
//}

func main() {

	pixelgl.Run(run)

	//ishapes = append(ishapes, randomShape())
	//
	////ishape = shape.NewSquare(10, 10) //&Square{Shape{X: 10, Y: 10}}
	//draw.RunWindow("Murat's Tetris", 400, 480, update)
	//ishapes = append(ishapes, randomShape())
	//for _, ishape := range ishapes {
	//	ishape.Move(dWindow)
	//}

}

func run() {
	runtime.GOMAXPROCS(1)
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(200, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(800, 100))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(500, 700))
	imd.Polygon(0)

	go Move(imd, win)
	//go Move(imd, win)
	//go Move(imd, win)

	for !win.Closed() {
		imd.Draw(win)
		win.Update()
	}
}

//for !win.Closed() {
//if win.JustPressed(pixelgl.MouseButtonLeft) {
//tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
//trees = append(trees, tree)
//matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(win.MousePosition()))
//}
func Move(imd *imdraw.IMDraw, win *pixelgl.Window) {

	var i float64
	for ii := 0; ii < 200; ii++ {
		imd.Push(pixel.V(10-i, 100))
		imd.Push(pixel.V(70-i, 100))
		imd.Push(pixel.V(150-i, 700))
		imd.Polygon(0)
		i -= 15
		time.Sleep(100 * time.Millisecond)
		if win.Pressed(pixelgl.KeySpace) {
			log.Print("KeySpace Pressed!")
		}
	}
}
func randomShape() shape.IShape {

	//rand.Seed(time.Now().UTC().UnixNano())
	//shapeNumber := rand.Intn(3)
	//if shapeNumber == 0 {
	//	return shape.NewSquare(175, 0)
	//}
	//if shapeNumber == 1 {
	//	return shape.NewLine(195, 0)
	//}
	//if shapeNumber == 2 {
	//	return shape.NewLShape(190, 0)
	//}
	return nil
}

//func update(window draw.Window) {
//
//	dWindow = &window
//
//	//if window.WasKeyPressed(draw.KeyEscape) {
//	//	window.Close()
//	//}
//	//if window.WasKeyPressed(draw.KeyN) {
//	//	//resetGame()
//	//}
//	//log.Print("Logging in Go!")
//	time.Sleep(100 * time.Millisecond)
//	//
//	//if !gameOver {
//	//
//	//	var hasMoving = true
//	//	for _, ishape := range ishapes {
//	//		hasMoving = ishape.Move(&window)
//	//	}
//	//
//	//	if !hasMoving {
//	//		ishapeCount++
//	//		ishapes = append(ishapes, randomShape())
//	//	}
//	//}
//
//}
