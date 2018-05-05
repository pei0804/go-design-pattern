package prototype

import "fmt"

//
// framwark
//

// Product インスタンスを複製する仕事を行うためのインターフェイス
// Use()をどうするか、どうやってクローンするかはサブクラスに任せられている
type Product interface {
	Use() string
	createClone() Product
}

// Manager Productインターフェイスを使って
// インスタンスを複製します（複製物がどういうものかは知らない）
type Manager struct {
	showCase map[string]Product
	product  Product
}

// Register 製品名とProductインターフェイスに準拠した何かを取得し
// showCaseへ登録する
func (m *Manager) Register(protoname string, product Product) {
	if m.showCase == nil {
		m.showCase = map[string]Product{}
	}
	m.showCase[protoname] = product
}

// Create 製品名から登録されているprototypeの複製を取得する
func (m *Manager) Create(protoname string) Product {
	// no err handling
	return m.showCase[protoname].createClone()
}

//
// Product
//

// MessageBox それっぽい何か
type MessageBox struct {
	name string
}

// Use Productインターフェイス準拠
func (mb *MessageBox) Use() string {
	return fmt.Sprintf("MessageBox: %s", mb.name)
}

// createClone インターフェイス準拠
// 本当は何か特殊な操作がある的な
func (mb *MessageBox) createClone() Product {
	return &MessageBox{mb.name}
}

// Underline それっぽい何か
type Underline struct {
	name string
}

// Use Productインターフェイス準拠
func (u *Underline) Use() string {
	return fmt.Sprintf("Underline: %s", u.name)
}

// createClone インターフェイス準拠
// 本当は何か特殊な操作がある的な
func (u *Underline) createClone() Product {
	return &Underline{u.name}
}
