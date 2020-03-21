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
				if b != "=" {
					c.Calculation = c.Calculation + b
				} else {
					cc, _ := gval.Evaluate(c.Calculation, nil)
					c.Calculation = fmt.Sprint(cc.(float64))
				}
			}
			btn := c.Theme.Button(b)
			btn.Layout(gtx, c.Buttons[b])
		})
	}
}

func main() {
	gofont.Register()
	calcApp := NewCalc()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
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
								calcApp.Theme.H1(calcApp.Calculation).Layout(gtx)
							})
					}),
					flexChild(gtx, 0.8, "ff303030", flex(gtx, layout.Horizontal,
						flexChild(gtx, 0.25, "ff303030", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "7")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "4")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "1")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "0")),
						)),
						flexChild(gtx, 0.25, "ff303030", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "8")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "5")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "2")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, ".")),
						)),
						flexChild(gtx, 0.25, "ff303030", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "9")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "6")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "3")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "=")),
						)),
						flexChild(gtx, 0.25, "ff303030", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "/")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "*")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "-")),
							flexChild(gtx, 0.25, "ff303030", calcApp.Button(gtx, "+")),
						)),
					)),
				)()
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
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
		//cs := gtx.Constraints
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
