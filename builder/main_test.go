package builder

import (
	"fmt"
	"testing"
)

func TestTextBuilder(t *testing.T) {
	expect := `===========================
[Greeting]

■ 朝から昼にかけて

Items[0]: おはようございます。
Items[1]: こんにちは。

■ 夜に

Items[0]: こんばんは。
Items[1]: おやすみなさい。
Items[2]: さようなら。

===========================
`
	director := NewDirector(&TextBuilder{})
	result := director.Build()
	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}

func TestHTMLBuilder(t *testing.T) {
	expect := `<html><head><title>Greeting</title></head></html><body><h1>Greeting</h1><p>朝から昼にかけて</p><ul><li>Items[0]: おはようございます。</li><li>Items[1]: こんにちは。</li></ul><p>夜に</p><ul><li>Items[0]: こんばんは。</li><li>Items[1]: おやすみなさい。</li><li>Items[2]: さようなら。</li></ul></body></html>`
	director := NewDirector(&HTMLBuilder{})
	result := director.Build()
	if result != expect {
		fmt.Println(result)
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}
