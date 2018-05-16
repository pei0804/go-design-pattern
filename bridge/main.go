package bridge

// DisplayImpl implementor 実装者 APIを実装するためのメソッドを規定する役
type DisplayImpl interface {
	RawOpen() string
	RawPrint() string
	RawClose() string
}

// Display Abstraction 抽象化 機能クラス階層の最上位クラス
// implementorのメソッドを使って
// 基本的な機能だけが記述されているクラス
type Display struct {
	impl DisplayImpl
}

// NewDisplay コンストラクタ
func NewDisplay(impl DisplayImpl) *Display {
	return &Display{
		impl: impl,
	}
}

// Open .
func (d *Display) Open() string {
	return d.impl.RawOpen()
}

// Print .
func (d *Display) Print() string {
	return d.impl.RawPrint()
}

// Close .
func (d *Display) Close() string {
	return d.impl.RawClose()
}

// Display .
func (d *Display) Display() string {
	return d.Open() + d.Print() + d.Close()
}

// CountDisplay Abstraction に対して機能を追加した
type CountDisplay struct {
	display *Display
}

// NewCountDisplay コンストラクタ
func NewCountDisplay(display *Display) *CountDisplay {
	return &CountDisplay{
		display: display,
	}
}

// MultiDisplay 追加した機能
func (c *CountDisplay) MultiDisplay(count int) string {
	result := c.display.Open()
	for i := 0; i < count; i++ {
		result += c.display.Print()
	}
	result += c.display.Close()
	return result
}

// StringDisplayImpl implementorのAPIを実装する役です
type StringDisplayImpl struct {
	str string
}

// NewStringDisplayImpl コンストラクタ
func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{
		str: str,
	}
}

// RawOpen .
func (s *StringDisplayImpl) RawOpen() string {
	return s.printLine()
}

// RawPrint .
func (s *StringDisplayImpl) RawPrint() string {
	return "|" + s.str + "|\n"
}

// RawClose .
func (s *StringDisplayImpl) RawClose() string {
	return s.printLine()
}

func (s *StringDisplayImpl) printLine() string {
	result := "*"
	for i := 0; i < len(s.str); i++ {
		result += "-"
	}
	result += "*\n"
	return result
}
