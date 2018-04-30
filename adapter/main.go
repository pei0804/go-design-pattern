package adpter

// Banner バナー
type Banner struct {
	name string
}

// NewBanner コンストラクタ
func NewBanner(name string) *Banner {
	return &Banner{
		name: name,
	}
}

func (b *Banner) getWithParen() string {
	return "(" + b.name + ")"
}

func (b *Banner) getWithAster() string {
	return "*" + b.name + "*"
}

// Print 必要とされているインターフェイス
type Print interface {
	getWeek() string
	getStrong() string
}

//
// フィールド保持
//

// PrintBannerComposition アダプター 構造体フィールドとしてBannerを保持
type PrintBannerComposition struct {
	Banner *Banner
}

// NewPrintBannerComposition コンストラクタ
func NewPrintBannerComposition(name string) *PrintBannerComposition {
	return &PrintBannerComposition{
		Banner: &Banner{
			name: name,
		},
	}
}

func (pb *PrintBannerComposition) getWeek() string {
	return pb.Banner.getWithParen()
}

func (pb *PrintBannerComposition) getStrong() string {
	return pb.Banner.getWithAster()
}

//
// 埋め込み
//

// PrintBannerEmbedding アダプター 埋め込みによる継承
type PrintBannerEmbedding struct {
	*Banner
}

// NewPrintBannerEmbedding コンストラクタ
func NewPrintBannerEmbedding(name string) *PrintBannerEmbedding {
	return &PrintBannerEmbedding{
		Banner: &Banner{
			name: name,
		},
	}
}

func (pb *PrintBannerEmbedding) getWeek() string {
	return pb.getWithParen()
}

func (pb *PrintBannerEmbedding) getStrong() string {
	return pb.getWithAster()
}
