package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	tree := NewSegmentTree()
	tree.Push(1, 10, "hello, world")
	tree.Push(5, 6, "how are you today?")
	tree.Push(6, 45, "test")
	err := tree.BuildTree()
	assert.NoError(t, err)

	results, err := tree.QueryIndex(1)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	result := <-results
	assert.Equal(t, "hello, world", result)

	// tree.Print()
}

func TestFindingSingleElement(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"
	tree.Push(1, 10, test)
	tree.BuildTree()

	results, err := tree.QueryIndex(4)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results

	if result != test {
		fmt.Println("wrongElement", result, test)
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

}

func TestFindingElementSizeZeroRange(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"
	tree.Push(1, 1, test)
	tree.BuildTree()

	results, err := tree.QueryIndex(1)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results

	if result != test {
		fmt.Println("wrongElement", result, test)
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

}

func TestFindingElementPseudoEndlessRange(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"
	tree.Push(1, Inf, test)
	tree.BuildTree()

	results, err := tree.QueryIndex(9999)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results

	if result != test {
		fmt.Println("wrongElement", result, test)
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

}

func TestFindingMultipleElements(t *testing.T) {

	tree := NewSegmentTree()
	tests := make([]interface{}, 5)
	tests[0] = "one"
	tests[1] = "two"
	tests[2] = 3.14
	tests[3] = 4
	tests[4] = "stuff"

	for i, test := range tests {
		tree.Push(10+i, 100+i, test)
	}

	tree.BuildTree()
	results, err := tree.QueryIndex(15)

	if err != nil {
		fmt.Println("queryFailed")
	}

	for result := range results {
		found, index := find(result, tests)
		if !found {
			fmt.Println("receivedUnexpected", result)
		}
		// Remove element from tests
		tests[index], tests = tests[len(tests)-1], tests[:len(tests)-1]
	}

	if len(tests) > 0 {
		for _, test := range tests {
			fmt.Println(fmt.Errorf("Did not find %v\n", test))
		}
	}

	if _, ok := <-results; ok != false {

	}

}

func TestFindingOverlappingElements(t *testing.T) {

	tree := NewSegmentTree()
	tests := make([]interface{}, 2)
	tests[0] = "one"
	tests[1] = "two"
	tree.Push(1, 10, tests[0])
	tree.Push(5, 15, tests[1])
	tree.BuildTree()

	// Index only in first range
	results, err := tree.QueryIndex(4)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results
	if result != tests[0] {
		fmt.Println("wrongElement", result, tests[0])
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

	// Index only in second range
	results, err = tree.QueryIndex(11)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result = <-results
	if result != tests[1] {
		fmt.Println("wrongElement", result, tests[1])
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

	// Index in both ranges
	results, err = tree.QueryIndex(6)
	if err != nil {
		fmt.Println("queryFailed")
	}

	for result := range results {
		found, index := find(result, tests)
		if !found {
			fmt.Println("receivedUnexpected", result)
		}
		// Remove element from tests
		tests[index], tests = tests[len(tests)-1], tests[:len(tests)-1]

	}

	if len(tests) > 0 {
		for _, test := range tests {
			fmt.Println(fmt.Errorf("Did not find %v\n", test))
		}
	}

	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

}

func TestOutOfRangeNotFound(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"

	tree.Push(1, 10, test)

	tree.BuildTree()
	results, err := tree.QueryIndex(20)

	if err != nil {
		fmt.Println("queryFailed")
	}

	if result, ok := <-results; ok != false {

		fmt.Println("receivedUnexpected", result)
	}

}

func TestAddingReverseDirection(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"
	tree.Push(10, 1, test)
	tree.BuildTree()

	results, err := tree.QueryIndex(4)
	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results

	if result != test {
		fmt.Println("wrongElement", result, test)
	}

}

func TestFindingSameElementTwice(t *testing.T) {

	tree := NewSegmentTree()
	test := "hello, world"

	tree.Push(1, 10, test)
	tree.Push(5, 15, test)

	tree.BuildTree()
	results, err := tree.QueryIndex(10)

	if err != nil {
		fmt.Println("queryFailed")
	}

	result := <-results

	if result != test {
		fmt.Println("wrongElement", result, test)
	}

	// `test` should only be sent once
	if _, ok := <-results; ok != false {
		fmt.Println("toManyElements")
	}

}
