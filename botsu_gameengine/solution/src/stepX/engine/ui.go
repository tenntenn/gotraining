package engine

// Button は ボタンを表すゲームオブジェクトです。
type Button struct {
	simpleObject
	// OnPush はボタンが押された時に呼ばれるイベントのリスナ
	OnPush EventListenerList
}

// NewButton はボタンを作ります。
func NewButton() *Button {
	return new(Button)
}

// Push はボタンを押します。
func (b *Button) Push() {
	if b.OnPush != nil {
		b.OnPush.Notify()
	}
}

// Text はテキストラベルを表すゲームオブジェクトです。
type Text struct {
	simpleObject
	// Text は表示する文字列です。
	Text string
}

// NewText はテキストを作ります。
func NewText(text string) *Text {
	return &Text{
		Text: text,
	}
}
