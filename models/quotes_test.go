// test code for the models package
package models

import (
	"testing"
)

// test GetBody
func TestGetBody(t *testing.T) {
	q:= &Quote{
		Body: "I think therefore I am",
		Author: "Rene Descartes",
		Submitter: "mesb",
	}

	if q.GetBody() != "I think therefore I am" {
		t.Errorf(`Expected body to be: %s but got %s`, "I think therefore I am.", q.GetBody())
	}

}
