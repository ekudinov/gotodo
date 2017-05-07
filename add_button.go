package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
)

//go:generate reactGen

// AddButtonDef is button element for add data to list
// When get EditButtonClicked hide self
// DataCollected - show self
// When click on button create AddButtonClicked
type AddButtonDef struct {
	r.ComponentDef
}

// AddButtonProps is property
type AddButtonProps struct {
	ID   string
	Name string
}

// AddButtonState is state
type AddButtonState struct {
	// is button hidden?
	isHide bool
}

// AddButton create button element
func AddButton(p AddButtonProps) *AddButtonDef {
	res := new(AddButtonDef)
	r.BlessElement(res, p)
	notifications.Subscribe(func(e EditButtonClicked) { res.hide() })
	notifications.Subscribe(func(e DataCollected) { res.show() })
	return res
}

// OnClick create AddButtonClicked message
func (rb *AddButtonDef) OnClick(e *r.SyntheticMouseEvent) {
	notifications.Dispatch(AddButtonClicked{})
}

// Render - render component
func (rb *AddButtonDef) Render() r.Element {
	id := rb.Props().ID
	name := rb.Props().Name
	var btn r.Element
	if rb.State().isHide {
		btn = r.Div(&r.DivProps{
			ID: id,
		})
	} else {
		btn = r.Div(&r.DivProps{
			ID: id,
		}, r.Button(&r.ButtonProps{
			OnClick: rb,
		}, r.S(name)),
		)
	}
	return btn
}

// AddButtonClicked - add button message
type AddButtonClicked struct{}

// hide add button
func (rb *AddButtonDef) hide() {
	st := rb.State()
	st.isHide = true
	rb.SetState(st)
}

// show add button
func (rb *AddButtonDef) show() {
	st := rb.State()
	st.isHide = false
	rb.SetState(st)
}
