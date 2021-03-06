// Template generated by reactGen

package main

import (
	r "myitcv.io/react"
)

// AppDef - application component
type AppDef struct {
	r.ComponentDef
	// collector element
	collector *CollectorDef
	// add button element
	addBtn *AddButtonDef
	// save button element
	saveBtn *SaveButtonDef
	// list elements
	list *ListElemDef
	// mode element
	mode *ModeDef
}

// App - create app
func App() *AppDef {
	res := new(AppDef)
	res.collector = Collector(CollectorProps{ID: "section-data"})
	res.mode = Mode(ModeProps{"id-mode"})
	res.list = ListElem(ListElemProps{ID: "el-list", Name: "Todo list:"})
	res.addBtn = AddButton(AddButtonProps{ID: "add-btn", Name: "Add to todo"})
	res.saveBtn = SaveButton(SaveButtonProps{ID: "save-btn", Name: "Save to todo"})
	r.BlessElement(res, nil)
	return res
}

// Render - render application component
func (a *AppDef) Render() r.Element {
	return r.Div(&r.DivProps{
		ID: "todo-input",
	},
		r.P(nil,
			r.S("Todo Application"),
		),
		a.mode,
		a.collector,
		a.addBtn, a.saveBtn,
		a.list,
	)
}
