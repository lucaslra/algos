package elems

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestWhenCalledWithValidDataReturnsValidResult(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1>")
	data, _ := html.Parse(node)

	result, err := ElementsTagByName(data, "h1")

	if result == nil || len(result) != 1 {
		t.Errorf("Result should contain exactly 1 result: %v", err)
	} else {
		t.Logf("Function returns valid result: %v", result)
	}
}

func TestWhenCalledWithoutElementsReturnError(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1>")
	data, _ := html.Parse(node)
	_, err := ElementsTagByName(data)

	if err == nil {
		t.Errorf("Function should fail if no element names are provided")
	} else {
		t.Logf("Function call without element names failed, as expected: %v", err)
	}
}

func TestWhenCalledWithValidContentResultIsAccessible(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1>")
	data, _ := html.Parse(node)
	result, err := ElementsTagByName(data, "h1")

	if err != nil {
		t.Error(err)
	}
	actual := result[0].FirstChild.Data
	expected := "Test"
	if actual != expected {
		t.Errorf("Result did not match expected content: %v", actual)
	}
}

func TestWhenCalledWithAtLeastOneValidElementReturnsValidResult(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1>")
	data, _ := html.Parse(node)
	result, err := ElementsTagByName(data, "h1", "h2", "h3")

	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Errorf("Result length should be exactly 1: %v", result)
	}
}

func TestWhenCalledWithNoExistingElementsShouldReturnEmptySlice(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1>")
	data, _ := html.Parse(node)
	result, err := ElementsTagByName(data, "span")

	if err != nil {
		t.Error(err)
	}

	if len(result) != 0 {
		t.Errorf("Result length should be exactly 0: %v", result)
	}
}

func TestWhenCalledWithMultipleExistingElementsShouldReturnEqualNumberOfResults(t *testing.T) {
	node := strings.NewReader("<h1>Test</h1><h2>Test</h2><h3>Test</h3>")
	data, _ := html.Parse(node)
	result, err := ElementsTagByName(data, "h1", "h2", "h3")

	if err != nil {
		t.Error(err)
	}

	if len(result) != 3 {
		t.Errorf("Result length should be exactly 3: %v", result)
	}
}

func TestWhenCalledWithNestedValidElementsShouldReturnMultipleResults(t *testing.T) {
	node := strings.NewReader("<h1>Test<h2>Test<h3>Test</h3></h2></h1>")
	data, _ := html.Parse(node)
	result, err := ElementsTagByName(data, "h1", "h2", "h3")

	if err != nil {
		t.Error(err)
	}

	if len(result) != 3 {
		t.Errorf("Result length should be exactly 3: %v", result)
	}
}