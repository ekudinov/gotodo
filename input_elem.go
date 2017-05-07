package main

import (
	"honnef.co/go/js/dom"
	r "myitcv.io/react"
)

//go:generate reactGen

// InputBarDef component for data from user
type InputBarDef struct {
	r.ComponentDef
}

// InputBarState - state
type InputBarState struct {
	// input data
	Value string
}

// InputBarProps - props
type InputBarProps struct {
	ID     string
	Name   string
	Holder string
}

// InputBar - create input component
func InputBar(p InputBarProps) *InputBarDef {
	res := new(InputBarDef)
	r.BlessElement(res, p)
	return res
}

// OnChange - when change input
func (i *InputBarDef) OnChange(e *r.SyntheticEvent) {
	val := e.Target().(*dom.HTMLInputElement).Value
	st := i.State()
	st.Value = val
	i.SetState(st)
}

// OnClick - when click input
func (i *InputBarDef) OnClick(e *r.SyntheticMouseEvent) {
	i.clear()
}

// clear field
func (i *InputBarDef) clear() {
	st := i.State()
	st.Value = ""
	i.SetState(st)
}

// Render - render component
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

// getValue for input element and after clear it
func (i *InputBarDef) getValue() string {
	val := i.State().Value
	i.clear()
	return val
}

// setValue for input element
func (i *InputBarDef) setValue(data string) {
	st := i.State()
	st.Value = data
	i.SetState(st)
}
