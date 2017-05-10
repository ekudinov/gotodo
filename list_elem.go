package main

import (
	"github.com/gu-io/gu/notifications"
	"github.com/satori/go.uuid"
	r "myitcv.io/react"
	"github.com/BurntSushi/xgb/res"
)

//go:generate reactGen

// ListElemDef - collection of todo elemens
// Message DataCollected - add item to map
// RemoveButtonClicked - remove item from elements
// EditButtonClicked - if key of element is in map
// it make EditDataSent
type ListElemDef struct {
	r.ComponentDef
}

// ListElemProps - props
type ListElemProps struct {
	ID   string
	Name string
}

// ListElemState - state
type ListElemState struct {
	// id current edit element
	editID string
	// collection items
	todos map[string]Item
}

// ListElem - create list component
func ListElem(p ListElemProps) *ListElemDef {
	res := new(ListElemDef)
	r.BlessElement(res, p)
	return res
}

// Render - render
func (le *ListElemDef) Render() r.Element {
	id := le.Props().ID
	name := le.Props().Name
	curID := le.State().editID
	show := true //flag to show buttons on element
	if curID != "" {
		show = false
	}

	var el r.Element
	var entries []*r.LiDef
	for id, item := range le.State().todos {
		//mark edited element as class "edited"
		class := ""
		if id == curID {
			class = "edited"
		}
		todoDef := Todo(TodoProps{todoID: id, class: class, Item: item, showBtn: show})
		entry := r.Li(nil, todoDef)
		entries = append(entries, entry)
	}
	if len(entries) > 0 {
		el = r.Ul(nil, entries...)
	} else {
		el = r.S("No items")
	}
	return r.Div(&r.DivProps{ID: id}, r.S(name), el)
}

// Equals needs for update state
func (le ListElemState) Equals(v ListElemState) bool {
	if len(le.todos) != len(v.todos) {
		return false
	}
	for i := range v.todos {
		if le.todos[i] != v.todos[i] {
			return false
		}
	}
	if le.editID != v.editID {
		return false
	}
	return true
}

// GetInitialState - init state
func (le *ListElemDef) GetInitialState() ListElemState {
	todos := make(map[string]Item, 0)
	todos[uuid.NewV4().String()] = Item{Name: "Wake up", Value: "In 6-00"}
	todos[uuid.NewV4().String()] = Item{Name: "Go", Value: "Go in 7-30"}
	st := ListElemState{}
	st.todos = todos
	st.editID = "" //no element for edit
	return st
}

// removeFromList - removes item from map
func (le *ListElemDef) removeFromList(e RemoveButtonClicked) {
	st := le.State()
	oldl := st.todos
	newl := make(map[string]Item, len(oldl))
	for k, v := range oldl {
		if k != e.ButtonID {
			newl[k] = v
		}
	}
	st.todos = newl
	le.SetState(st)
}

// addToList - add item to map
// if id edited is empty, generate new id
// else update value for edited id
func (le *ListElemDef) addToList(e DataCollected) {
	// work with map
	st := le.State()
	oldMap := st.todos
	newMap := make(map[string]Item, 0)
	for k, v := range oldMap {
		newMap[k] = v
	}
	curID := st.editID
	// if new element generate uuid
	if curID == "" {
		curID = uuid.NewV4().String()
	}
	newMap[curID] = Item{Name: e.Name, Value: e.Value}
	st.todos = newMap
	st.editID = "" // reset id edit element
	le.SetState(st)
}

// editmode - get id of element to edit, save it and
// after create EditDataSent with data of element
func (le *ListElemDef) editmode(e EditButtonClicked) {
	st := le.State()
	todos := st.todos
	item, ok := todos[e.ButtonID]
	if ok {
		notifications.Dispatch(EditDataSent{Item: item})
	}
	//set id todo element is editing
	st.editID = e.ButtonID
	le.SetState(st)
}

// EditDataSent - message with data to edit
type EditDataSent struct {
	Item
}

// subscribe for notifications
func (le *ListElemDef) ComponentDidMount() {
	notifications.Subscribe(func(e DataCollected) { le.addToList(e) })
	notifications.Subscribe(func(e RemoveButtonClicked) { le.removeFromList(e) })
	notifications.Subscribe(func(e EditButtonClicked) { le.editmode(e) })
}