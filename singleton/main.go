package singleton

type singleton struct {
}

var instance *singleton

// GetInstance インスタンス取得用メソッドのみ公開することで
// 一意なインスタンスであることを保証する
// lintエラーが出るのは、パブリックな関数がプライベートな型を返しているため
// http://www.songmu.jp/riji/entry/2017-10-28-some-singleton-in-golang.html
func GetInstance() *singleton { // nolint
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
