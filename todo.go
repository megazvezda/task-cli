package main

import (
	"errors"
	"fmt"
	"time"
)

/* Todo struct contains all the needed lines to describe a task. */
type Todo struct {
	ID          int
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

/* This is basically an array or list of Todo, so tasks. */
type Todos []Todo

/* Add() is made for adding new tasks to the list.  */
func (todos *Todos) Add(description string) {

	var newID int
	if len(*todos) == 0 {
		newID = 1
	} else {
		last := (*todos)[len(*todos)-1]
		newID = last.ID + 1
	}

	todo := Todo{
		ID:          newID,
		Description: description,
		IsCompleted: false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

// ValidateIndex takes an index which is an int, and can return an error. It is used in other functions to check if the index is valid.
func (todos *Todos) ValidateIndex(index int) error {
	// we need to validate the index.
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

// delete takes an index, can return error. deletes by appending the objects of the array, excluding the needed one
func (todos *Todos) Delete(index int) error {
	idx, err := todos.FindByID(index)
	if err != nil {
		return err
	}
	// actualIndex := index - 1 // (deprecated) actualIndex prevents the deleting of index 2 task when calling index 1, because of the zero-based indexing
	t := *todos
	if err := t.ValidateIndex(idx); err != nil {
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...) // still using actualIndex here

	return nil
}

// toggle takes an index, can return error. it validates if index exists and has an isCompleted parameter, it has a switch basically, pretty smart
func (todos *Todos) Toggle(index int) error {
	idx, err := todos.FindByID(index)
	if err != nil {
		return err
	}
	t := *todos
	if err := t.ValidateIndex(idx); err != nil {
		return err
	}

	completed := t[idx].IsCompleted

	if !completed {
		completed := time.Now()
		t[idx].CompletedAt = &completed
	}
	t[idx].IsCompleted = !completed

	return nil
}

// FindByID takes an ID from the original struct, can return an actual ID or an error.
// goes through the range of all tasks, if ID matches it returns an index. Or an error.
func (todos *Todos) FindByID(id int) (int, error) {
	for i, todo := range *todos {
		if todo.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("ID not found")
}
