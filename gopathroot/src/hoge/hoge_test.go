package hoge_test

import (
	"fmt"
	"hoge"
	"testing"
)

// Test〜(t *testing.T)
func TestHoge(t *testing.T) {
	if actual := hoge.Hoge(); actual != "hoge" {
		t.Errorf("hoge.Hoge must return hoge but %s", actual)
	}
}

func ExampleHoge() {
	fmt.Println(hoge.Hoge())
	// Output: hoge
}
