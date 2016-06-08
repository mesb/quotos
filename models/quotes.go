// implementation of quotos models
package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	ERROR_TYPE   = "Error"
	SUCCESS_TYPE = "Success"
)

//  object that stores a message.
// a message could be an error message or a success message
type Message struct {
	Typ  string
	Body map[string]string
}

// construct a new message object
func MakeMessage() *Message {
	m := new(Message)
	m.Body = make(map[string]string)

	return m
}

// set message type
func (m *Message) SetType(t string) {
	m.Typ = t
}

// add message text
func (m *Message) AddText(key, val string) {
	m.Body[key] = val
}

func (m *Message) AddTextMap(msg map[string]string) {
	m.Body = msg
}

// get the Message type
func (m Message) GetType() string {
	return m.Typ
}

func (m Message) String() string {
	return fmt.Sprintf("{type: %s, messages: %v}", m.GetType(), m.Body)
}

// helper struct for creating quote objects
type QuoteAux struct {
	Body      string `json:"body"`
	Submitter string `json:"submitter"`
	Author    string `json:"author"`
}

// object that stores a full and invdividual quotes object
type Quote struct {
	Body          string `json:"body"`
	Submitter     string `json:"submitter"`
	Author        string `json:"author"`
	SubmittedTime time.Time
	Flagged       bool
}


func NewQuote(q QuoteAux) *Quote{
	return &Quote{
		Body: q.GetBody(),
		Submitter: q.GetSubmitter(),
		Author: q.GetAuthor(),
		SubmittedTime: time.Now(),
		Flagged : false,
	}

}

// returns a string representation of a quote aux  object
func (q QuoteAux) String() string {
	return fmt.Sprintf("{body: %s, author: %s, submitter: %s}", q.GetBody(),q.GetAuthor(), q.GetSubmitter())
}

// getters
// return the body of a quote
func (q QuoteAux) GetBody() string {
	return q.Body
}

// returns the submitter of a quote
func (q QuoteAux) GetSubmitter() string {
	return q.Submitter
}

// returns the author of a quote
func (q QuoteAux) GetAuthor() string {
	return q.Author
}

// check if string is up to required length
func IsValidLength(s string, length int) bool {
	return length < len(strings.TrimSpace(s))
}

// check if string is empty
func IsEmpty(s string) bool {
	return "" == strings.TrimSpace(s)
}

func (q QuoteAux) Validate() map[string]string {

	result := make(map[string]string)

	/*
		if IsValidLength(q.GetBody(), 5) {
			result["Invalid Quote Length"]=  "Your quote should be made of at least 5 characters"
		}
	*/

	if IsEmpty(q.GetBody()) {

		result["Empty Text"] = "Quote text needed!"
	}

	if IsEmpty(q.GetAuthor()) {

		result["Empty Author field"] = "Provide an author for the quote"
	}

	if IsEmpty(q.GetSubmitter()) {
		result["Empty Submitter field"] = "Provide your name as submitter for the quote"
	}

	return result
}
