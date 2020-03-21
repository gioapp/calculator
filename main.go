package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/gioapp/calculator/calc"
	"github.com/p9c/learngio/helpers"
	"image"
	"image/color"
)

func main() {
	gofont.Register()
	calcApp := calc.NewCalc()
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
						flexChild(gtx, 0.25, "ff3030cf", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "7")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "4")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "1")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "0")),
						)),
						flexChild(gtx, 0.25, "ff3030cf", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "8")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "5")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "2")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, ".")),
						)),
						flexChild(gtx, 0.25, "ff3030cf", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "9")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "6")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "3")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "=")),
						)),
						flexChild(gtx, 0.25, "ff3030cf", flex(gtx, layout.Vertical,
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "/")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "*")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "-")),
							flexChild(gtx, 0.25, "ff3030cf", calcApp.Button(gtx, "+")),
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
