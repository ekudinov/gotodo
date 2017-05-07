package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
)

//go:generate reactGen

// Element for mode status show
// Show changemode status - DataCollected
// Show edit status - EditButtonClicked
// Error - ErrorValidation
type ModeDef struct {
	r.ComponentDef
}

type ModeState struct {
	// current status
	Status interface{}
}

type ModeProps struct {
	ID string
}

func Mode(p ModeProps) *ModeDef {
	res := new(ModeDef)
	r.BlessElement(res, p)
	notifications.Subscribe(func(e DataCollected) { res.setStatus(NormalStatus{"Normal"}) })
	notifications.Subscribe(func(e EditButtonClicked) { res.setStatus(EditStatus{"Edit"}) })
	notifications.Subscribe(func(e ErrorValidation) { res.setStatus(ErrorStatus{e.Msg}) })
	return res
}

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

func (s *ModeDef) GetInitialState() ModeState {
	st := ModeState{NormalStatus{"Normal"}}
	return st
}

func (s *ModeDef) setStatus(state Status) {
	st := s.State()
	st.Status = state
	s.SetState(st)
}

type ModeStateChanged struct {
	Current Status
}

type Status interface {
	String() string
}

type ErrorStatus struct {
	Text string
}
type NormalStatus struct {
	Text string
}

type EditStatus struct {
	Text string
}

func (e ErrorStatus) String() string {
	return e.Text
}

func (n NormalStatus) String() string {
	return n.Text
}

func (e EditStatus) String() string {
	return e.Text
}
