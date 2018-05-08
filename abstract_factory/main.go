package abstractFactory

import (
	"fmt"

	"github.com/yosssi/gohtml"
)

// factory
// 抽象的な工場・部品・製品を含むパッケージ

// Item LinkとTrayのスーパークラス
type Item interface {
	makeHTML() string
}

// Link HTMLのハイパーリンクを抽象的に表現したクラス
type Link interface {
	Item
}

// Tray 複数のlinkやtrayを集めてひとまとまりにしたものを表す（お盆のようなイメージ）
type Tray interface {
	Item
	AddItem(item Item)
}

// baseTray trayに必要な必須要素
type baseTray struct {
	traies []Item
}

func (bt *baseTray) AddItem(item Item) {
	bt.traies = append(bt.traies, item)
}

// Page HTMLページ全体を抽象的に表現したクラス
type Page interface {
	AddContent(item Item)
	Output() string
}
type basePage struct {
	contents []Item
}

// AddContent Itemインターフェイスに準拠したものをコンテンツに追加する
func (bp *basePage) AddContent(item Item) {
	bp.contents = append(bp.contents, item)
}

// Factory 抽象的な部品をまとめる
type Factory interface {
	CreateLink(caption string, link string) Link
	CreateTray(caption string) Tray
	CreatePage(title string, author string) Page
}

// listFactory
// 具体的な工場・部品・製品を含むパッケージ

// HTML

// HTMLFactory 具体的な動作
type HTMLFactory struct {
}

// CreateLink linkインターフェイスに準拠したHTMLLink作成
func (hf *HTMLFactory) CreateLink(caption string, link string) Link {
	return &HTMLLink{captipon: caption, url: link}
}

// CreateTray trayインターフェイスに準拠したHTMLTray作成
func (hf *HTMLFactory) CreateTray(caption string) Tray {
	return &HTMLTray{caption: caption}
}

// CreatePage pageインターフェイスに準拠したHTMLPage作成
func (hf *HTMLFactory) CreatePage(title string, author string) Page {
	return &HTMLPage{title: title, author: author}
}

// HTMLLink リンク
type HTMLLink struct {
	captipon string
	url      string
}

func (hl *HTMLLink) makeHTML() string {
	return fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", hl.url, hl.captipon)
}

// HTMLTray コンテンツ
type HTMLTray struct {
	baseTray
	caption string
}

func (hl *HTMLTray) makeHTML() string {
	result := fmt.Sprintf("<li>%s<ul>", hl.caption)
	for _, v := range hl.traies {
		result += v.makeHTML()
	}
	result += "</ul></li>"
	return result
}

// HTMLPage HTMLの外枠
type HTMLPage struct {
	basePage
	title  string
	author string
}

// Output HTMLの外枠を作成し、内包しているcontentsを展開し、フォーマット済みのHTMLを返す
func (hp *HTMLPage) Output() string {
	result := fmt.Sprintf("<html><head><title>%s</title></head>", hp.title)
	result += fmt.Sprintf("<body><h1>%s</h1>", hp.title)
	result += "<ul>"
	for _, v := range hp.contents {
		result += v.makeHTML()
	}
	result += "</ul>"
	result += fmt.Sprintf("<hr><address>%s</address>", hp.author)
	result += "</body></html>"
	return gohtml.Format(result)
}

// 以下面倒過ぎるのでやめたけど、いい感じに出来る。はいおｋ

// // Table
//
// // TableFactory 具体的な動作
// type TableFactory struct {
// }
//
// // CreateLink linkインターフェイスに準拠したHTMLLink作成
// func (hf *TableFactory) CreateLink(caption string, link string) Link {
// 	return &TableLink{captipon: caption, url: link}
// }
//
// // CreateTray trayインターフェイスに準拠したHTMLTray作成
// func (hf *TableFactory) CreateTray(caption string) Tray {
// 	return &TableTray{caption: caption}
// }
//
// // CreatePage pageインターフェイスに準拠したHTMLPage作成
// func (hf *TableFactory) CreatePage(title string, author string) Page {
// 	return &TablePage{title: title, author: author}
// }
//
// // TableLink リンク
// type TableLink struct {
// 	captipon string
// 	url      string
// }
//
// func (hl *TableLink) makeHTML() string {
// 	return fmt.Sprintf("<td><a href=\"%s\">%s</a></td>", hl.url, hl.captipon)
// }
//
// // TableTray コンテンツ
// type TableTray struct {
// 	baseTray
// 	caption string
// }
//
// func (tt *TableTray) makeHTML() string {
// 	result := "<td><table width=\"\100\" border=\"1\"><tr>"
// 	result += fmt.Sprintf("<td bgcolor=\"#cccccc\" align=\"center\" colspan=\"%s\"><b>%s</b></td>", len(tt.baseTray.traies), tt.caption)
// 	for _, v := range hl.traies {
// 		result += v.makeHTML()
// 	}
// 	result += "</ul></li>"
// 	return result
// }
//
// // TablePage HTMLの外枠
// type TablePage struct {
// 	basePage
// 	title  string
// 	author string
// }
//
// // Output HTMLの外枠を作成し、内包しているcontentsを展開し、フォーマット済みのHTMLを返す
// func (hp *TablePage) Output() string {
// 	result := fmt.Sprintf("<html><head><title>%s</title></head>", hp.title)
// 	result += fmt.Sprintf("<body><h1>%s</h1>", hp.title)
// 	result += "<ul>"
// 	for _, v := range hp.contents {
// 		result += v.makeHTML()
// 	}
// 	result += "</ul>"
// 	result += fmt.Sprintf("<hr><address>%s</address>", hp.author)
// 	result += "</body></html>"
// 	return gohtml.Format(result)
// }
