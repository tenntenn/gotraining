package engine

// GameObject はゲーム中で使われる汎用的なオブジェクトを表します。
type GameObject interface{}

// Behaviour は GameObject が振る舞いを持つことを表します。
type Behaviour interface {
	Start(s *Scene)
	Update(s *Scene)
	OnDestroy(s *Scene)
}

// Transform は GameObject が位置や親子関係を持つことを表します。
type Transform interface {
	// 親の Transform を取得します。
	Parent() Transform
	// 絶対位置を取得します。
	Position() (x, y int)
}

// SimpleObject はシンプルな GameObject です。
// 以下のインタフェースを実装します。
// ・Transform
// ・Behaviour
type SimpleObject struct {
	// 親からの相対のX座標です。
	LocalX int
	// 親からの相対のY座標です。
	LocalY int
	// 親
	parent Transform
}

// Position は Transform.Position の実装です。
func (o *SimpleObject) Position() (int, int) {

	x, y := o.LocalX, o.LocalY
	if p := o.Parent(); p != nil {
		x0, y0 := p.Position()
		x += x0
		y += y0
	}

	return x, y
}

// Parent は Transform.Parent の実装です。
func (o *SimpleObject) Parent() Transform {
	return o.parent
}

// SetParent は親を設定します。
func (o *SimpleObject) SetParent(t Transform) {
	o.parent = t
}

// Start は Behaviour.Start の実装です。
func (o SimpleObject) Start(s *Scene) {}

// Start は Behaviour.Update の実装です。
func (o SimpleObject) Update(s *Scene) {}

// Start は Behaviour.OnDestroy の実装です。
func (o SimpleObject) OnDestroy(s *Scene) {}

// SimpleObjectの非公開用の型です。
// 公開された型を埋め込むと外部パッケージから差し替えられる可能性があるためです。
type simpleObject struct {
	SimpleObject
}
