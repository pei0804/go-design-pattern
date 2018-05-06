package builder

import "fmt"

// Builder 文章作成に必要なメソッドをまとめたインターフェイス
type Builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

// Director Builderで宣言されていメソッドを使って文章を作成する
type Director struct {
	builder Builder
}

// NewDirector コンストラクタ
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Build 文章作成
func (d *Director) Build() string {
	result := d.builder.makeTitle("Greeting")
	result += d.builder.makeString("朝から昼にかけて")
	result += d.builder.makeItems([]string{
		"おはようございます。",
		"こんにちは。",
	})
	result += d.builder.makeString("夜に")
	result += d.builder.makeItems([]string{
		"こんばんは。",
		"おやすみなさい。",
		"さようなら。",
	})
	result += d.builder.close()
	return result
}

// TextBuilder Builderクラスのサブクラス
type TextBuilder struct {
}

func (tb *TextBuilder) makeTitle(title string) string {
	result := "===========================\n"
	result += fmt.Sprintf("[%s]\n", title)
	result += "\n"
	return result
}

func (tb *TextBuilder) makeString(str string) string {
	result := fmt.Sprintf("■ %s\n", str)
	result += "\n"
	return result
}

func (tb *TextBuilder) makeItems(items []string) string {
	result := ""
	for k, v := range items {
		result += fmt.Sprintf("Items[%d]: %s\n", k, v)
	}
	result += "\n"
	return result
}

func (tb *TextBuilder) close() string {
	return "===========================\n"
}

// HTMLBuilder Builderクラスのサブクラス
type HTMLBuilder struct {
}

func (hb *HTMLBuilder) makeTitle(title string) string {
	result := fmt.Sprintf("<html><head><title>%s</title></head></html><body>", title)
	result += fmt.Sprintf("<h1>%s</h1>", title)
	return result
}

func (hb *HTMLBuilder) makeString(str string) string {
	result := fmt.Sprintf("<p>%s</p>", str)
	return result
}

func (hb *HTMLBuilder) makeItems(items []string) string {
	result := "<ul>"
	for k, v := range items {
		result += fmt.Sprintf("<li>Items[%d]: %s</li>", k, v)
	}
	result += "</ul>"
	return result
}

func (hb *HTMLBuilder) close() string {
	return "</body></html>"
}
