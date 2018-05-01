package templateMethod

// AbstractDisplay printerの抽象メソッドを使って、Displayを実行するだけのTemplate Method
type AbstractDisplay struct {
}

type printer interface {
	open() string
	print() string
	close() string
}

// Display printerを付かて文字列を取得する
func (ad *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

// CharDisplay Hが渡されたら<<HHHHH>>と表示する
type CharDisplay struct {
	*AbstractDisplay
	char rune
}

// NewCharDisplay コンストラクタ
func NewCharDisplay(char rune) *CharDisplay {
	return &CharDisplay{
		char: char,
	}
}

func (cd *CharDisplay) open() string {
	return "<<"
}

func (cd *CharDisplay) print() string {
	return string(cd.char)
}

func (cd *CharDisplay) close() string {
	return ">>"
}

// StringDispay HelloWorldと渡されると、
type StringDispay struct {
	*AbstractDisplay
	str string
}

// NewStringDisplay コンストラクタ
func NewStringDisplay(str string) *StringDispay {
	return &StringDispay{
		str: str,
	}
}

func (sd *StringDispay) open() string {
	return sd.printLine()
}

func (sd *StringDispay) print() string {
	return "| " + sd.str + " |\n"
}

func (sd *StringDispay) close() string {
	return sd.printLine()
}

func (sd *StringDispay) printLine() string {
	line := "+-"
	for _ = range sd.str {
		line += "-"
	}
	line += "-+\n"
	return line
}
