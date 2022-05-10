package validators

import (
	"fmt"
	"github.com/ralugr/filter-service/internal/common"
	"github.com/ralugr/filter-service/internal/model"
	"reflect"
	"testing"
)

var validator = NewImageValidator()

func TestGetTags(t *testing.T) {
	input := []struct {
		m        model.Message
		state    string
		expected []string
	}{
		{common.MockMessage1, model.Accepted, []string{"<!--state: Accepted   -->", "<!--state: accepted-->"}},
		{common.MockMessage1, model.Rejected, []string{"<!--state: Rejected-->", "<!--    state:    rejected-->"}},
		{common.MockMessage2, model.Accepted, nil},
		{common.MockMessage3, model.Accepted, []string{"<!--state: Accepted   -->"}},
		{common.MockMessage3, model.Rejected, nil},
		{common.MockMessage4, model.Accepted, nil},
		{common.MockMessage4, model.Rejected, []string{"<!--state: rejected   -->"}},
	}

	for tc, tt := range input {
		actual := validator.getTags(&tt.m, tt.state)
		fmt.Println(" Starting test case ", tc)

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}

func TestHasTag(t *testing.T) {
	input := []struct {
		m        model.Message
		state    string
		expected bool
	}{
		{common.MockMessage1, model.Accepted, true},
		{common.MockMessage1, model.Rejected, true},
		{common.MockMessage2, model.Accepted, false},
		{common.MockMessage2, model.Rejected, false},
		{common.MockMessage3, model.Accepted, true},
		{common.MockMessage3, model.Rejected, false},
		{common.MockMessage4, model.Accepted, false},
		{common.MockMessage4, model.Rejected, true},
	}

	for tc, tt := range input {
		actual := validator.hasTag(&tt.m, tt.state)
		fmt.Println("Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}

func TestHasImages(t *testing.T) {
	input := []struct {
		m        model.Message
		expected bool
	}{
		{common.MockMessage1, true},
		{common.MockMessage2, false},
		{common.MockMessage3, true},
		{common.MockMessage4, true},
		{common.MockMessage5, true},
		{common.MockMessage6, false},
	}

	for tc, tt := range input {
		actual := validator.hasImages(&tt.m)
		fmt.Println("Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}

func TestValidate(t *testing.T) {
	input := []struct {
		m        model.Message
		expected error
		newState string
	}{
		{common.MockMessage1, nil, model.Rejected},
		{common.MockMessage2, nil, model.Accepted},
		{common.MockMessage3, nil, model.Queued},
		{common.MockMessage4, nil, model.Queued},
		{common.MockMessage5, nil, model.Rejected},
		{common.MockMessage6, nil, model.Accepted},
	}

	for tc, tt := range input {
		actual := validator.Validate(&tt.m)
		fmt.Println("Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}

		if tt.m.State != tt.newState {
			t.Errorf("\nExpected %v \nActual   %v", tt.newState, tt.m.State)

		}
	}
}
