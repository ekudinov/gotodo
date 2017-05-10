package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
	"github.com/BurntSushi/xgb/res"
)

//go:generate reactGen

// ModeDef - element for mode status show
// Show normal mode status - DataCollected
// Show edit status - EditButtonClicked
// Error - ErrorValidation
type ModeDef struct {
	r.ComponentDef
}

// ModeState - state
type ModeState struct {
	// current status
	Status interface{}
}

// ModeProps - props
type ModeProps struct {
	ID string
}

// Mode - create mode component
func Mode(p ModeProps) *ModeDef {
	res := new(ModeDef)
	r.BlessElement(res, p)
	return res
}

// Render - render
func (s *ModeDef) Render() r.Element {
	st := s.State().Status
	var status r.Element
	switch e := st.(type) {
	case ErrorStatus:
		status = r.Span(&r.SpanProps{
			ID: "err",
		}, r.S(e.String()))
	case NormalStatus:
		status = r.Span(&r.SpanProps{
			ID: "norm",
		}, r.S(e.String()))
	case EditStatus:
		status = r.Span(&r.SpanProps{
			ID: "edit",
		}, r.S(e.String()))
	}
	return r.Div(&r.DivProps{
		ID: s.Props().ID,
	}, r.S("mode:"), status)
}

// GetInitialState - init state
func (s *ModeDef) GetInitialState() ModeState {
	st := ModeState{NormalStatus{"Normal"}}
	return st
}

// setStatus - set mode
func (s *ModeDef) setStatus(state Status) {
	st := s.State()
	st.Status = state
	s.SetState(st)
}

// ModeStateChanged - status changed
type ModeStateChanged struct {
	Current Status
}

// Status - status
type Status interface {
	String() string
}

// ErrorStatus - error
type ErrorStatus struct {
	Text string
}

// NormalStatus - normal
type NormalStatus struct {
	Text string
}

// EditStatus - edit
type EditStatus struct {
	Text string
}

// String - impl
func (e ErrorStatus) String() string {
	return e.Text
}

// String - impl
func (n NormalStatus) String() string {
	return n.Text
}

// String - impl
func (e EditStatus) String() string {
	return e.Text
}

// subscribe for notifications
func (s *ModeDef) ComponentDidMount() {
	notifications.Subscribe(func(e DataCollected) { s.setStatus(NormalStatus{"Normal"}) })
	notifications.Subscribe(func(e EditButtonClicked) { s.setStatus(EditStatus{"Edit"}) })
	notifications.Subscribe(func(e ErrorValidation) { s.setStatus(ErrorStatus{e.Msg}) })
}