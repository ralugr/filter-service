package validators

import (
	"fmt"
	"testing"

	"github.com/ralugr/filter-service/internal/common"
	"github.com/ralugr/filter-service/internal/model"
)

var linkValidator = NewLinkValidator()

func TestHasExternalLink(t *testing.T) {
	input := []struct {
		m        model.Message
		expected bool
	}{
		{common.MockMessage1, false},
		{common.MockMessage2, true},
		{common.MockMessage3, false},
		{common.MockMessage4, false},
		{common.MockMessage5, false},
		{common.MockMessage6, false},
	}

	for tc, tt := range input {
		actual := linkValidator.hasExternalLink(&tt.m)
		fmt.Println(" Starting test case ", tc)

		if actual != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, actual)
		}
	}
}
