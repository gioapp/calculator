package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/PaesslerAG/gval"
	"github.com/p9c/learngio/helpers"
	"image"
	"image/color"
)

type Calc struct {
	Buttons     map[string]*widget.Button
	Theme       *material.Theme
	Calculation string
}

func NewCalc() Calc {
	buttons := map[string]*widget.Button{
		"0": new(widget.Button),
		"1": new(widget.Button),
		"2": new(widget.Button),
		"3": new(widget.Button),
		"4": new(widget.Button),
		"5": new(widget.Button),
		"6": new(widget.Button),
		"7": new(widget.Button),
		"8": new(widget.Button),
		"9": new(widget.Button),
		"+": new(widget.Button),
		"-": new(widget.Button),
		"/": new(widget.Button),
		"*": new(widget.Button),
		".": new(widget.Button),
		"=": new(widget.Button),
		"c": new(widget.Button),
		")": new(widget.Button),
		"(": new(widget.Button),
		"←": new(widget.Button),
	}
	return Calc{
		Buttons: buttons,
		Theme:   material.NewTheme(),
	}
}

func (c *Calc) Button(gtx *layout.Context, b string) func() {
	return func() {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, func() {
			for c.Buttons[b].Clicked(gtx) {
				switch b {
				case "=":
					cc, _ := gval.Evaluate(c.Calculation, nil)
					c.Calculation = fmt.Sprint(cc.(float64))
				case "c":
					c.Calculation = ""
				case "←":
					if last := len(c.Calculation) - 1; last >= 0 {
						c.Calculation = c.Calculation[:last]
					}
				default:
					c.Calculation = c.Calculation + b
				}
			}
			btn := c.Theme.Button(b)
			btn.Background = helpers.HexARGB("ff888888")
			switch b {
			case "=":
				btn.Background = helpers.HexARGB("ff30cf30")
			case "c":
				btn.Background = helpers.HexARGB("ffcf3030")
			case "←":
				btn.Background = helpers.HexARGB("ffcfcf30")
			case "(", ")":
				btn.Background = helpers.HexARGB("ffcf30cf")
			case "/", "*", "-", "+":
				btn.Background = helpers.HexARGB("ff3030cf")
			default:
				btn.Background = helpers.HexARGB("ff30cfcf")
			}
			btn.Layout(gtx, c.Buttons[b])
		})
	}
}

func main() {
	gofont.Register()
	calcApp := NewCalc()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(500), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				flex(gtx, layout.Vertical,
					flexChild(gtx, 0.2, "ffcfcfcf", func() {
						layout.E.Layout(gtx,
							func() {
								calcApp.Theme.H3(calcApp.Calculation).Layout(gtx)
							})
					}),
					flexChild(gtx, 0.8, "ff303030", flex(gtx, layout.Horizontal,
						calcApp.column(gtx, "(", "7", "4", "1", "0"),
						calcApp.column(gtx, ")", "8", "5", "2", "."),
						calcApp.column(gtx, "c", "9", "6", "3", "="),
						calcApp.column(gtx, "←", "/", "*", "-", "+"),
					)),
				)()
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func (c *Calc) column(gtx *layout.Context, ba, bb, bc, bd, be string) layout.FlexChild {
	return flexChild(gtx, 0.25, "ff303030", flex(gtx, layout.Vertical,
		flexChild(gtx, 0.2, "ff303030", c.Button(gtx, ba)),
		flexChild(gtx, 0.2, "ff303030", c.Button(gtx, bb)),
		flexChild(gtx, 0.2, "ff303030", c.Button(gtx, bc)),
		flexChild(gtx, 0.2, "ff303030", c.Button(gtx, bd)),
		flexChild(gtx, 0.2, "ff303030", c.Button(gtx, be)),
	))
}

func flex(gtx *layout.Context, axis layout.Axis, content ...layout.FlexChild) func() {
	return func() {
		layout.Flex{
			Axis: axis,
		}.Layout(gtx, content...)
	}
}

func flexChild(gtx *layout.Context, size float32, bg string, content func()) layout.FlexChild {
	return layout.Flexed(size, func() {
		cs := gtx.Constraints
		drawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB(bg), [4]float32{0, 0, 0, 0}, unit.Dp(0))
		content()
	})
}

func drawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)
		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}
