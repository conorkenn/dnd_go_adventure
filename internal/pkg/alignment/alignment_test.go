package alignment

import (
	"reflect"
	"testing"
)

func TestAlignment(t *testing.T) {

	t.Run("Test Getters", func(t *testing.T) {
		testAlignment := "Dummy"
		testDescription := "Dummy Desc"
		dummy := Alignment{testAlignment, testDescription}

		checkResult(t, dummy.name, testAlignment)
		checkResult(t, dummy.description, testDescription)
	})

	t.Run("Test GetAlignments, Alignment Found", func(t *testing.T) {
		expectedName := "Lawful Good"
		expectedDesc := "Characters of this alignment believe in order, justice, and morality. They will typically adhere to laws and rules but are not necessarily rigid or inflexible. They strive to do good and uphold the greater good, often sacrificing their own interests for the sake of others."
		lg := "LG"
		result, _ := GetAlignments(lg)

		checkResult(t, result.name, expectedName)
		checkResult(t, result.description, expectedDesc)

	})

	t.Run("Test GetAlignments, Alignment not found", func(t *testing.T) {
		testAlignment := "ABC"
		_, ok := GetAlignments(testAlignment)
		if ok {
			t.Errorf("found %v expected to fail", testAlignment)
		}
	})
}

func checkResult(t testing.TB, got, want string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
