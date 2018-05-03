package factoryMethod

// goには抽象クラスのような概念がない
// type Product interface {
// 	use()
// }

// framework

// User
type User interface {
	Use() string
}

// Factory template methodパターン
// create()を使ってproductのインスタンスを作る
type Factory struct {
}

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

func (f *Factory) Create(factory creater, owner string) User {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

// idcard

// IDCard ID card
type IDCard struct {
	owner string
}

// NewIDCard コンストラクタ
func NewIDCard(owner string) *IDCard {
	return &IDCard{owner}
}

// Use Userインターフェイスに必要なメソッド
func (ic *IDCard) Use() string {
	// ID cardを使う何かをする
	return ic.owner
}

// GetOwner オーナー情報取得
func (ic *IDCard) GetOwner() string {
	return ic.owner
}

// IDCardFactory ID cardを作成する
type IDCardFactory struct {
	*Factory
	owners []*string
}

// createProduct ID cardを作成する
func (icf *IDCardFactory) createProduct(owner string) User {
	return NewIDCard(owner)
}

// registerProduct オーナーを登録する
func (icf *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	icf.owners = append(icf.owners, &owner)
}

// GetOwners オーナーリストを返す
func (icf *IDCardFactory) GetOwners() []*string {
	return icf.owners
}
