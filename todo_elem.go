package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
)

//go:generate reactGen

// Item of todo element
// No have state and only generate messages
// Click remove button - create RemoveButtonClicked
// Click edit button - create EditButtonClicked
type Item struct {
	Name  string
	Value string
}

// TodoDef - component presents item
// when is edited buttons no show
// and red color
type TodoDef struct {
	r.ComponentDef
}

// TodoProps - props
type TodoProps struct {
	// todo id = button id
	todoID string
	// class for element
	class string
	// flag show buttons or no
	showBtn bool
	// item in todo element
	Item
}

// Todo - create component
func Todo(p TodoProps) *TodoDef {
	t := new(TodoDef)
	r.BlessElement(t, p)
	return t
}

// Render - render
func (t *TodoDef) Render() r.Element {
	todoProp := t.Props()
	id := todoProp.todoID
	item := todoProp.Item
	class := todoProp.class
	// show buttons in normal or hide in edit mode
	var btns r.Element
	if todoProp.showBtn {
		btns = r.Div(nil, r.Button(&r.ButtonProps{
			ID:      id,
			OnClick: remove{t},
		}, r.S("X")),
			r.Button(&r.ButtonProps{
				ID:      id,
				OnClick: edit{t},
			}, r.S("Edit")),
		)
	}
	return r.Div(&r.DivProps{
		ID:        id,
		ClassName: class,
	}, r.S("Name:"), r.S(item.Name), r.S(" Value:"), r.S(item.Value), btns,
	)
}

type remove struct{ t *TodoDef }
type edit struct{ t *TodoDef }

// OnClick - remove button click
func (t remove) OnClick(e *r.SyntheticMouseEvent) {
	id := e.Target().ID()
	notifications.Dispatch(RemoveButtonClicked{ButtonID: id})
	e.PreventDefault()
}

// OnClick - edit button click
func (t edit) OnClick(e *r.SyntheticMouseEvent) {
	id := e.Target().ID()
	notifications.Dispatch(EditButtonClicked{ButtonID: id})
	e.PreventDefault()
}

// RemoveButtonClicked - message for remove button
type RemoveButtonClicked struct {
	ButtonID string
}

// EditButtonClicked - message for edit button
type EditButtonClicked struct {
	ButtonID string
}
