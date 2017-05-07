package main

import (
	"honnef.co/go/js/dom"
	r "myitcv.io/react"
)

//go:generate reactGen

// Input component for data from user
type InputBarDef struct {
	r.ComponentDef
}

type InputBarState struct {
	// input data
	Value string
}

type InputBarProps struct {
	ID     string
	Name   string
	Holder string
}

func InputBar(p InputBarProps) *InputBarDef {
	res := new(InputBarDef)
	r.BlessElement(res, p)
	return res
}

func (i *InputBarDef) OnChange(e *r.SyntheticEvent) {
	val := e.Target().(*dom.HTMLInputElement).Value
	st := i.State()
	st.Value = val
	i.SetState(st)
}

func (i *InputBarDef) OnClick(e *r.SyntheticMouseEvent) {
	i.clear()
}

// clear field
func (i *InputBarDef) clear() {
	st := i.State()
	st.Value = ""
	i.SetState(st)
}

func (i *InputBarDef) Render() r.Element {
	id := i.Props().ID
	name := i.Props().Name
	holder := i.Props().Holder
	//render
	return r.Div(&r.DivProps{
		ID: id,
	},
		r.Label(nil, r.S(name)),
		r.Input(&r.InputProps{
			ClassName:   id,
			Type:        "text",
			Placeholder: holder,
			Value:       i.State().Value,
			OnChange:    i,
			OnClick:     i,

		}),
	)
}

// get value for input element and after clear it
func (i *InputBarDef) getValue() string {
	val := i.State().Value
	i.clear()
	return val
}

// set value for input element
func (i *InputBarDef) setValue(data string) {
	st := i.State()
	st.Value = data
	i.SetState(st)
}
