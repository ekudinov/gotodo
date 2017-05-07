package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
)

//go:generate reactGen

// SaveButtonDef - button element for save edited data
// When EditButtonClicked it show self
// When click on button create SaveButtonClicked
// DataCollected message hide button
type SaveButtonDef struct {
	r.ComponentDef
}

// SaveButtonProps - props
type SaveButtonProps struct {
	ID   string
	Name string
}

// SaveButtonState - state
type SaveButtonState struct {
	// how button?
	isShow bool
}

// SaveButton - create component
func SaveButton(p SaveButtonProps) *SaveButtonDef {
	res := new(SaveButtonDef)
	r.BlessElement(res, p)
	notifications.Subscribe(func(e EditButtonClicked) { res.show() })
	notifications.Subscribe(func(e DataCollected) { res.hide() })
	return res
}

// OnClick - when click button hide
func (sb *SaveButtonDef) OnClick(e *r.SyntheticMouseEvent) {
	notifications.Dispatch(SaveButtonClicked{})

}

// Render - render
func (sb *SaveButtonDef) Render() r.Element {
	id := sb.Props().ID
	name := sb.Props().Name
	if sb.State().isShow {
		return r.Div(&r.DivProps{
			ID: id,
		}, r.Button(&r.ButtonProps{
			OnClick: sb,
		}, r.S(name)),
		)
	}
	return r.Div(&r.DivProps{
		ID: id,
	})
}

// SaveButtonClicked message - save button clicked
type SaveButtonClicked struct{}

// show button
func (sb *SaveButtonDef) show() {
	st := sb.State()
	st.isShow = true
	sb.SetState(st)
}

// hide button if DataCollected is - ok
func (sb *SaveButtonDef) hide() {
	st := sb.State()
	st.isShow = false
	sb.SetState(st)
}
