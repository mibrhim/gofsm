/*
Package gofsm implements a simple library for Finit state machine .

Author: Mahmoud Ibrahim

This package could be used by Implementing the stateHandler for each
state or mode for a certain object
*/
package gofsm

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the Set
type Object generic.Type

// define new types
type Command int

// Transitions commands constants
const (
	CMD_START              Command = 0
	CMD_TERMINATE                  = 1
	CMD_MOVE_TO_NEXT_STATE         = 2
)

// StateMachine struct
type StateMachine struct {
	name string
}

// the Transition Event which contains the command
// and the data to be sent to the target state
type Event struct {
	command Command
	data    Object
}

type StateHandler interface {
	isAccepted(Command) bool
	handle(Event) StateHandler
	getName() string
}

// Global Variables to used in the gofsm package
var currentState StateHandler
var events chan Event = make(chan Event)

func FSM_handle(newEvent chan Event) {
	fmt.Printf("FSM go routine started \n")
	for true {
		event := <-newEvent
		fmt.Printf("\n**** New FSM Event Receieved **** \n")
		if !currentState.isAccepted(event.command) {
			fmt.Printf("[Invalid Event] The %d command is not allowed in the %s state \n",
				event.command,
				currentState.getName())
		}
		currentState = currentState.handle(event)
	}
}

func (s *StateMachine) init(initState StateHandler) {
	currentState = initState
	go FSM_handle(events)
	if events == nil {
		fmt.Println("null channel")
	}
}

func (s *StateMachine) trigger(event Event) {
	if events == nil {
		fmt.Println("null channel")
	}
	events <- event
}

func (s *StateMachine) start() {
	data := "START FSM"
	s.trigger(Event{CMD_START, data})
}

func (s *StateMachine) terminate() {
	data := "TERMINATE FSM"
	s.trigger(Event{CMD_TERMINATE, data})
}
