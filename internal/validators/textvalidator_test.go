package validators

import (
	"fmt"
	"github.com/ralugr/filter-service/internal/common"
	"github.com/ralugr/filter-service/internal/model"
	"testing"
)

var textValidator = NewTextValidator()

func TestStartsWith(t *testing.T) {
	input := []struct {
		m        model.Message
		expected bool
	}{
		{common.MockMessage1, false},
		{common.MockMessage2, true},
		{common.MockMessage3, true},
		{common.MockMessage4, false},
		{common.MockMessage5, false},
		{common.MockMessage6, true},
	}

	for tc, tt := range input {
		actual := textValidator.startsWithHeading1(&tt.m)
		fmt.Println(" Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}

func TestHasParagraph(t *testing.T) {
	input := []struct {
		m        model.Message
		expected bool
	}{
		{common.MockMessage1, true},
		{common.MockMessage2, true},
		{common.MockMessage3, false},
		{common.MockMessage4, false},
		{common.MockMessage5, true},
		{common.MockMessage6, true},
	}

	for tc, tt := range input {
		actual := textValidator.hasParagraph(&tt.m)
		fmt.Println(" Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}

func TestValidateText(t *testing.T) {
	input := []struct {
		m        model.Message
		expected error
		newState model.MsgState
	}{
		{common.MockMessage1, nil, model.Rejected},
		{common.MockMessage2, nil, model.Accepted},
		{common.MockMessage3, nil, model.Rejected},
		{common.MockMessage4, nil, model.Rejected},
		{common.MockMessage5, nil, model.Rejected},
		{common.MockMessage6, nil, model.Accepted},
	}

	for tc, tt := range input {
		actual := textValidator.Validate(&tt.m)
		fmt.Println("Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}

		if tt.m.State != tt.newState {
			t.Errorf("\nExpected %v \nActual   %v", tt.newState, tt.m.State)

		}
	}
}
