package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var joysticks [16]*sdl.Joystick

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 400, 400}
	surface.FillRect(&rect, 0x000000ff)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			case *sdl.JoyAxisEvent:
				fmt.Printf("[%d ms] JoyAxis\ttype:%d\twhich:%c\taxis:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Axis, t.Value)
			case *sdl.JoyBallEvent:
				fmt.Printf("[%d ms] JoyBall\ttype:%d\twhich:%d\tball:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.Ball, t.XRel, t.YRel)
			case *sdl.JoyButtonEvent:
				fmt.Printf("[%d ms] JoyButton\ttype:%d\twhich:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.Button, t.State)
			case *sdl.JoyHatEvent:
				fmt.Printf("[%d ms] JoyHat\ttype:%d\twhich:%d\that:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Hat, t.Value)
			case *sdl.JoyDeviceAddedEvent:
				joysticks[int(t.Which)] = sdl.JoystickOpen(t.Which)
				if joysticks[int(t.Which)] != nil {
					fmt.Printf("Joystick %d connected\n", t.Which)
				}
			case *sdl.JoyDeviceRemovedEvent:
				if joystick := joysticks[int(t.Which)]; joystick != nil {
					joystick.Close()
				}
				fmt.Printf("Joystick %d disconnected\n", t.Which)
			default:
				fmt.Printf("Some event\n")

			}
		}
	}
}
