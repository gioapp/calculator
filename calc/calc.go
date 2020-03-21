package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
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
				c.Calculation = c.Calculation + b
			}
			btn := c.Theme.Button(b)
			btn.Layout(gtx, c.Buttons[b])
		})
	}
}
