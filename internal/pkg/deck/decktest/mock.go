// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package decktest

import (
	"github.com/google/uuid"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
	"sync"
)

// ManagerMock is a mock implementation of deck.Manager.
//
//	func TestSomethingThatUsesManager(t *testing.T) {
//
//		// make and configure a mocked deck.Manager
//		mockedManager := &ManagerMock{
//			NewDeckFunc: func(shuffled bool, cards ...string) (*deck.Deck, error) {
//				panic("mock out the NewDeck method")
//			},
//			OpenFunc: func(deckID uuid.UUID) (*deck.Deck, error) {
//				panic("mock out the Open method")
//			},
//		}
//
//		// use mockedManager in code that requires deck.Manager
//		// and then make assertions.
//
//	}
type ManagerMock struct {
	// NewDeckFunc mocks the NewDeck method.
	NewDeckFunc func(shuffled bool, cards ...string) (*deck.Deck, error)

	// OpenFunc mocks the Open method.
	OpenFunc func(deckID uuid.UUID) (*deck.Deck, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewDeck holds details about calls to the NewDeck method.
		NewDeck []struct {
			// Shuffled is the shuffled argument value.
			Shuffled bool
			// Cards is the cards argument value.
			Cards []string
		}
		// Open holds details about calls to the Open method.
		Open []struct {
			// DeckID is the deckID argument value.
			DeckID uuid.UUID
		}
	}
	lockNewDeck sync.RWMutex
	lockOpen    sync.RWMutex
}

// NewDeck calls NewDeckFunc.
func (mock *ManagerMock) NewDeck(shuffled bool, cards ...string) (*deck.Deck, error) {
	callInfo := struct {
		Shuffled bool
		Cards    []string
	}{
		Shuffled: shuffled,
		Cards:    cards,
	}
	mock.lockNewDeck.Lock()
	mock.calls.NewDeck = append(mock.calls.NewDeck, callInfo)
	mock.lockNewDeck.Unlock()
	if mock.NewDeckFunc == nil {
		var (
			deckOut *deck.Deck
			errOut  error
		)
		return deckOut, errOut
	}
	return mock.NewDeckFunc(shuffled, cards...)
}

// NewDeckCalls gets all the calls that were made to NewDeck.
// Check the length with:
//
//	len(mockedManager.NewDeckCalls())
func (mock *ManagerMock) NewDeckCalls() []struct {
	Shuffled bool
	Cards    []string
} {
	var calls []struct {
		Shuffled bool
		Cards    []string
	}
	mock.lockNewDeck.RLock()
	calls = mock.calls.NewDeck
	mock.lockNewDeck.RUnlock()
	return calls
}

// Open calls OpenFunc.
func (mock *ManagerMock) Open(deckID uuid.UUID) (*deck.Deck, error) {
	callInfo := struct {
		DeckID uuid.UUID
	}{
		DeckID: deckID,
	}
	mock.lockOpen.Lock()
	mock.calls.Open = append(mock.calls.Open, callInfo)
	mock.lockOpen.Unlock()
	if mock.OpenFunc == nil {
		var (
			deckOut *deck.Deck
			errOut  error
		)
		return deckOut, errOut
	}
	return mock.OpenFunc(deckID)
}

// OpenCalls gets all the calls that were made to Open.
// Check the length with:
//
//	len(mockedManager.OpenCalls())
func (mock *ManagerMock) OpenCalls() []struct {
	DeckID uuid.UUID
} {
	var calls []struct {
		DeckID uuid.UUID
	}
	mock.lockOpen.RLock()
	calls = mock.calls.Open
	mock.lockOpen.RUnlock()
	return calls
}
