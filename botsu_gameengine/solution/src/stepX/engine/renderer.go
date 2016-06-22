package engine

// Renderer は GameObject がレンダリングできることを表します。
type Renderer interface {
	// 指定したカメラを使ってレンダリングを行います。
	Render(c Camera, o GameObject)
}

// Rectangle は矩形を表します。
type Rectangle struct {
	X0, Y0 int
	X1, Y1 int
}

// Rect は左上の座標と幅と高さを指定して矩形を作成します。
func Rect(x0, y0, width, height int) Rectangle {
	return Rectangle{x0, y0, x0 + width, y0 + height}
}

// Width は矩形の幅を取得します。
func (r Rectangle) Width() int {
	return r.X1 - r.X0
}

// Height は矩形の高さを取得します。
func (r Rectangle) Height() int {
	return r.Y1 - r.Y0
}

// Camera はグローバル座標をカメラ座標に変換します。
type Camera interface {
	// グローバル座標をカメラ座標に変換
	Map(r Rectangle) (Rectangle, bool)
}

// SimpleCamera はシンプルな Camera の実装です。
type SimpleCamera struct {
	*simpleObject
	width, height int
}

// Map は Camera.Map の実装です。
func (c *SimpleCamera) Map(r Rectangle) (Rectangle, bool) {
	var to Rectangle

	cx, cy := c.Position()

	to.X0 = max(r.X0-cx, 0)
	to.X1 = min(r.X1-cx, c.width)

	to.Y0 = max(r.Y0-cy, 0)
	to.Y1 = min(r.Y1-cy, c.height)

	if to.X0 <= to.X1 && to.X0 >= 0 && to.X0 < cx+c.width && to.Width() <= c.width &&
		to.Y0 <= to.Y1 && to.Y0 >= 0 && to.Y0 < cy+c.height && to.Height() <= c.height {
		return to, true
	}

	return Rectangle{}, false
}

// NewSimpleCamera は SimpleCamera を作成します。
func NewSimpleCamera(width, height int) *SimpleCamera {
	return &SimpleCamera{
		simpleObject: new(simpleObject),
		width:        width,
		height:       height,
	}
}

// MultiRenderer は 複数の Renderer をまとめます。
type MultiRenderer []Renderer

// Render は Renderer.Render の実装です。
func (mr MultiRenderer) Render(c Camera, o GameObject) {
	for _, r := range mr {
		r.Render(c, o)
	}
}

func max(n, m int) int {
	if n > m {
		return n
	}

	return m
}

func min(n, m int) int {
	if n < m {
		return n
	}

	return m
}
