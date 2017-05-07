package main

import (
	"github.com/gu-io/gu/notifications"
	"github.com/satori/go.uuid"
	r "myitcv.io/react"
)

//go:generate reactGen

// List element is collection of todo elemens
// Message DataCollected - add item to map
// RemoveButtonClicked - remove item from elements
// EditButtonClicked - if key of element is in map
// it make EditDataSent
type ListElemDef struct {
	r.ComponentDef
}

type ListElemProps struct {
	ID   string
	Name string
}

type ListElemState struct {
	// id current edit element
	editId string
	// collection items
	todos map[string]Item
}

func ListElem(p ListElemProps) *ListElemDef {
	res := new(ListElemDef)
	r.BlessElement(res, p)
	notifications.Subscribe(func(e DataCollected) { res.addToList(e) })
	notifications.Subscribe(func(e RemoveButtonClicked) { res.removeFromList(e) })
	notifications.Subscribe(func(e EditButtonClicked) { res.editmode(e) })
	return res
}

func (le *ListElemDef) Render() r.Element {
	id := le.Props().ID
	name := le.Props().Name
	curId := le.State().editId
	show := true //flag to show buttons on element
	if curId != "" {
		show = false
	}

	var el r.Element
	var entries []*r.LiDef
	for id, item := range le.State().todos {
		//mark edited element as class "edited"
		class := ""
		if id == curId {
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

// needs for update state
func (le ListElemState) Equals(v ListElemState) bool {
	if len(le.todos) != len(v.todos) {
		return false
	}
	for i := range v.todos {
		if le.todos[i] != v.todos[i] {
			return false
		}
	}
	if le.editId != v.editId {
		return false
	}
	return true
}

func (le *ListElemDef) GetInitialState() ListElemState {
	todos := make(map[string]Item, 0)
	todos[uuid.NewV4().String()] = Item{Name: "Wake up", Value: "In 6-00"}
	todos[uuid.NewV4().String()] = Item{Name: "Go", Value: "Go in 7-30"}
	st := ListElemState{}
	st.todos = todos
	st.editId = "" //no element for edit
	return st
}

// remove item from map
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

// add item to map
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
	curId := st.editId
	// if new element generate uuid
	if curId == "" {
		curId = uuid.NewV4().String()
	}
	newMap[curId] = Item{Name: e.Name, Value: e.Value}
	st.todos = newMap
	st.editId = "" // reset id edit element
	le.SetState(st)
}

// get id of element to edit, save it and
// after create EditDataSent with data of element
func (le *ListElemDef) editmode(e EditButtonClicked) {
	st := le.State()
	todos := st.todos
	item, ok := todos[e.ButtonID]
	if ok {
		notifications.Dispatch(EditDataSent{Item: item})
	}
	//set id todo element is editing
	st.editId = e.ButtonID
	le.SetState(st)
}

// message with data to edit
type EditDataSent struct {
	Item
}
