package engine

// Collider は 衝突判定ができることを表します。
type Collider interface {
	// 衝突判定を行います。
	Collide(c Collider)
}

// 2次元の長方形のコライダーです。
// スケールとか回転とか考えてません。
type boxCollider interface {
	Collider
	// 当たり判定の領域を取得します。
	Bounds() Rectangle
}

// 2次元の長方形のコライダーです。
// スケールとか回転とか考えてません。
type BoxCollider struct {
	hasCollided bool

	// 位置を取得するためのTransformです。
	Transform Transform
	// 当たり判定の幅です。
	Width int
	// 当たり判定の高さです。
	Height int
	// ゲームオブジェクトが衝突判定内に入ってきた場合に
	// 発生するイベントのリスナーのリストです。
	OnEnter EventListenerList
	// ゲームオブジェクトが衝突判定内に出て行った場合に
	// 発生するイベントのリスナーのリストです。
	OnExit EventListenerList
}

// NewBoxCollider はBoxColliderを作成します。
func NewBoxCollider(t Transform, w, h int) *BoxCollider {
	return &BoxCollider{
		Transform: t,
		Width:     w,
		Height:    h,
	}
}

// Collide はCollide.Collideの実装です。
// BoxCollider以外には反応しません。
func (bc *BoxCollider) Collide(c Collider) {
	bc2, ok := c.(boxCollider)
	if !ok {
		return
	}

	obj, ok := c.(GameObject)
	if !ok {
		return
	}

	//　当たり判定
	if bc.Bounds().X0 <= bc2.Bounds().X1 && bc2.Bounds().X0 <= bc.Bounds().X1 &&
		bc.Bounds().Y0 <= bc2.Bounds().Y1 && bc2.Bounds().Y0 <= bc.Bounds().Y1 {
		if !bc.hasCollided {
			bc.hasCollided = true
			EventNotify(bc.OnEnter, obj)
		}
	} else {
		if bc.hasCollided {
			bc.hasCollided = false
			EventNotify(bc.OnExit, obj)
		}
	}
}

// Rectangle は当たり判定の領域を取得します。
func (bc *BoxCollider) Bounds() Rectangle {
	var x, y int
	if bc.Transform != nil {
		x, y = bc.Transform.Position()
	}

	return Rect(x, y, bc.Width, bc.Height)
}
