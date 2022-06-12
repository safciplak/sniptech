package sawmill

import (
	"reflect"
	"testing"
)

func TestSawmill_SawTrunk(t *testing.T) {
	type testCase struct {
		name     string
		trunks   []int
		expected []int
		actual   []int
	}

	t.Run("test for some scenarios for equal", func(t *testing.T) {
		tests := []testCase{
			testCase{name: "2 4 1 saw trunks - response 2 1 3 1 scenario", trunks: []int{2, 4, 1}, expected: []int{2, 1, 3, 1}, actual: make([]int, 0)},
			testCase{name: "1 3 2 saw trunks - response 1 2 1 2 scenario", trunks: []int{1, 3, 2}, expected: []int{1, 2, 1, 2}, actual: make([]int, 0)},
			testCase{name: "3 2 1 saw trunks - response 3 2 1 scenario", trunks: []int{3, 2, 1}, expected: []int{3, 2, 1}, actual: make([]int, 0)},
			testCase{name: "4 4 saw trunks - response 3 1 2 2 scenario", trunks: []int{4, 4}, expected: []int{3, 1, 2, 2}, actual: make([]int, 0)},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				SawTrunk(test.trunks, SawBlocks, &test.actual)
				if !reflect.DeepEqual(test.expected, test.actual) {
					t.Errorf("Expected: %v\nActual: %v", test.expected, test.actual)
				}
			})
		}
	})
}

func TestSawmill_Equal(t *testing.T) {
	type testCase struct {
		name     string
		a        []int
		b        []int
		expected bool
		actual   bool
	}

	t.Run("test for some scenarios for equal", func(t *testing.T) {
		tests := []testCase{
			testCase{name: "should return true for equal case", a: []int{1, 2, 3}, b: []int{1, 2, 3}, expected: true},
			testCase{name: "should return true for equal case", a: []int{1, 2, 3}, b: []int{1, 2, 3, 4}, expected: false},
			testCase{name: "should return true for not equal case", a: []int{1, 2, 3}, b: []int{1, 2, 4}, expected: false},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := Equal(test.a, test.b)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}

func TestSawmill_Profit(t *testing.T) {
	type TestCase struct {
		name       string
		sawnTrunks []int
		expected   int
		actual     int
	}

	t.Run("test for some scenarios for profit calculation", func(t *testing.T) {
		tests := []TestCase{
			TestCase{name: "total profit : 3", sawnTrunks: []int{1, 3, 2}, expected: 3},
			TestCase{name: "total profit : 2", sawnTrunks: []int{1, 3, 2}, expected: 3},
			TestCase{name: "total profit : 1", sawnTrunks: []int{1, 1, 2}, expected: 1},
			TestCase{name: "total profit : 0", sawnTrunks: []int{1, 1, 3, 3}, expected: 0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := Profit(test.sawnTrunks)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}
