package main

import (
	"github.com/gu-io/gu/notifications"
	r "myitcv.io/react"
)


//go:generate reactGen

// Collector component must work with input components
// When add button is clicked it collect data from input
// and after validation send DataCollected with data
// When edit button is clicked it get edited data and fill input elements
// When save button is clicked it collect data from input
// and send DataCollected with data
// if validation false create ErrorValidation
type CollectorDef struct {
	r.ComponentDef
	InputName  *InputBarDef
	InputValue *InputBarDef
}

type CollectorProps struct {
	ID string
}

func Data(p CollectorProps) *CollectorDef {
	res := new(CollectorDef)
	res.InputName = InputBar(InputBarProps{ID: "todo-name", Name: "Name ", Holder: "Todo Name"})
	res.InputValue = InputBar(InputBarProps{ID: "todo-value", Name: "Value ", Holder: "Todo Value"})
	r.BlessElement(res, p)
	notifications.Subscribe(func(e AddButtonClicked) { res.collect() })
	notifications.Subscribe(func(e EditDataSent) { res.fill(e) })
	notifications.Subscribe(func(e SaveButtonClicked) { res.collect() })
	return res
}

func (d *CollectorDef) Render() r.Element {
	return r.Div(&r.DivProps{
		ID: d.Props().ID,
	},
		d.InputName,
		r.BR(nil),
		d.InputValue)
}

// data collected message
// when data is ok it must be generated
type DataCollected struct {
	Item
}

// validation error message
// when data is false it must be generated
type ErrorValidation struct {
	Msg string
}

// collect data from input elements
// validate and create proper message
func (d *CollectorDef) collect() {
	name := d.InputName.getValue()
	val := d.InputValue.getValue()
	if validate(name) && validate(val) {
		notifications.Dispatch(DataCollected{Item: Item{Name: name, Value: val}})
	} else {
		notifications.Dispatch(ErrorValidation{Msg: "Input error"})
	}
}

// data validation
func validate(str string) bool {
	if str == "" {
		return false
	}
	return true
}

// fill input elements with data
func (d *CollectorDef) fill(e EditDataSent) {
	name := e.Name
	d.InputName.setValue(name)
	value := e.Value
	d.InputValue.setValue(value)
}
