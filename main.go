package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/PaesslerAG/gval"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"image/color"
)

type (
	D = layout.Dimensions
	C = layout.Context
)
type Calc struct {
	Buttons     map[string]*widget.Clickable
	Theme       *material.Theme
	Calculation string
}

func NewCalc() Calc {
	buttons := map[string]*widget.Clickable{
		"0": new(widget.Clickable),
		"1": new(widget.Clickable),
		"2": new(widget.Clickable),
		"3": new(widget.Clickable),
		"4": new(widget.Clickable),
		"5": new(widget.Clickable),
		"6": new(widget.Clickable),
		"7": new(widget.Clickable),
		"8": new(widget.Clickable),
		"9": new(widget.Clickable),
		"+": new(widget.Clickable),
		"-": new(widget.Clickable),
		"/": new(widget.Clickable),
		"*": new(widget.Clickable),
		".": new(widget.Clickable),
		"=": new(widget.Clickable),
		"c": new(widget.Clickable),
		")": new(widget.Clickable),
		"(": new(widget.Clickable),
		"←": new(widget.Clickable),
	}
	return Calc{
		Buttons: buttons,
		Theme:   material.NewTheme(gofont.Collection()),
	}
}

func (c *Calc) Button(b string) func(gtx C) D {
	return func(gtx C) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx C) D {
			for c.Buttons[b].Clicked() {
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
			btn := material.Button(c.Theme, c.Buttons[b], b)
			btn.Background = HexARGB("ff888888")
			switch b {
			case "=":
				btn.Background = HexARGB("ff30cf30")
			case "c":
				btn.Background = HexARGB("ffcf3030")
			case "←":
				btn.Background = HexARGB("ffcfcf30")
			case "(", ")":
				btn.Background = HexARGB("ffcf30cf")
			case "/", "*", "-", "+":
				btn.Background = HexARGB("ff3030cf")
			default:
				btn.Background = HexARGB("ff30cfcf")
			}
			return btn.Layout(gtx)
		})
	}
}

func main() {
	calcApp := NewCalc()

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(500), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		var ops op.Ops
		for {
			select {
			case e := <-w.Events():
				switch e := e.(type) {
				case system.FrameEvent:
					gtx := layout.NewContext(&ops, e)
					Fill(gtx, HexARGB("ff303030"))
					lyt.Format(gtx, "vflexa(middle,f(0.2,_),f(0.8,_))",
						func(gtx C) D {
							return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx C) D {
								Fill(gtx, HexARGB("ffffffff"))
								return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx C) D {
									out := material.H3(calcApp.Theme, calcApp.Calculation)
									out.Alignment = text.End
									return out.Layout(gtx)
								})
							})
						},
						func(gtx C) D {
							return lyt.Format(gtx, "hflexa(middle,f(0.2,_),f(0.2,_),f(0.2,_),f(0.2,_))",
								calcApp.column("(", "7", "4", "1", "0"),
								calcApp.column(")", "8", "5", "2", "."),
								calcApp.column("c", "9", "6", "3", "+"),
								calcApp.column("←", "/", "*", "-", "="),
							)
						},
					)
					e.Frame(gtx.Ops)
				}
				w.Invalidate()
			}
		}
	}()
	app.Main()
}

func (c *Calc) column(ba, bb, bc, bd, be string) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexa(middle,f(0.2,_),f(0.2,_),f(0.2,_),f(0.2,_),f(0.2,_))",
			c.Button(ba),
			c.Button(bb),
			c.Button(bc),
			c.Button(bd),
			c.Button(be),
		)
	}
}

func HexARGB(s string) (c color.RGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}

func Fill(gtx layout.Context, col color.RGBA) layout.Dimensions {
	cs := gtx.Constraints
	d := cs.Min
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	return layout.Dimensions{Size: d}
}
