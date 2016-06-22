package engine

import "sync"

// SceneFactory はシーンを作成します。
type SceneFactory interface {
	Create() *Scene
}

// Scene はシーンを表す構造体です。
type Scene struct {
	// メインのカメラ
	MainCamera Camera
	// レンダラ
	Renderer Renderer
	// シーンに登録されているゲームオブジェクト
	gameObjects []GameObject
	// Colliderを実装しているゲームオブジェクト
	colliders map[int]GameObject
	// 削除予定のゲームオブジェクト
	deleted map[int]GameObject
	// 入力判定
	inputMutex sync.RWMutex
	input      map[int]bool
}

// NewScene は新しいシーンを作成します。
func NewScene() *Scene {
	return &Scene{
		deleted:   make(map[int]GameObject),
		colliders: make(map[int]GameObject),
		input:     make(map[int]bool),
	}
}

// SetInput は指定したインデックスの入力があったことを設定します。
// ロックを取ります。
func (s *Scene) SetInput(i int) {
	s.inputMutex.Lock()
	defer s.inputMutex.Unlock()
	s.input[i] = true
}

// UnsetInput は指定したインデックスの入力がなかったことにします。
// ロックを取ります。
func (s *Scene) UnsetInput(i int) {
	s.inputMutex.Lock()
	defer s.inputMutex.Unlock()
	delete(s.input, i)
}

// Input は指定したインデックスの入力があったかどうか取得します。
// ロックを取ります。
func (s *Scene) Input(i int) bool {
	s.inputMutex.RLock()
	defer s.inputMutex.RUnlock()
	return s.input[i]
}

// ClearInput は入力のフラグをすべてOFFにします。
// ロックを取ります。
func (s *Scene) ClearInput() {
	s.inputMutex.Lock()
	defer s.inputMutex.Unlock()
	s.input = map[int]bool{}
}

// Add はシーンに GameObject を追加します。
func (s *Scene) Add(o GameObject) {
	s.gameObjects = append(s.gameObjects, o)
	if _, ok := o.(Collider); ok {
		s.colliders[len(s.gameObjects)-1] = o
	}
}

// Remove はシーンから GameObject を削除します。
func (s *Scene) Remove(o GameObject) {
	for i := range s.gameObjects {
		if s.gameObjects[i] == o {
			s.deleted[i] = s.gameObjects[i]
			if b, ok := s.deleted[i].(Behaviour); ok {
				b.OnDestroy()
			}
			return
		}
	}
}

// Start はシーンの初期化処理を行います。
func (s *Scene) Start() {
	for _, o := range s.gameObjects {
		if b, ok := o.(Behaviour); ok {
			b.Start(s)
		}
	}
}

// Update はシーンの1フレームの処理を行います。
func (s *Scene) Update() {
	if len(s.deleted) > 0 {
		var gameObjects []GameObject
		for i := range s.gameObjects {
			if s.deleted[i] == nil {
				gameObjects = append(gameObjects, s.gameObjects[i])
			} else {
				delete(s.colliders, i)
			}
		}
		s.deleted = map[int]GameObject{}
		s.gameObjects = gameObjects
	}

	for i, o1 := range s.colliders {
		if s.deleted[i] != nil {
			continue
		}

		for j, o2 := range s.colliders {
			if i == j || s.deleted[i] != nil || s.deleted[j] != nil {
				continue
			}
			o1.(Collider).Collide(o2.(Collider))
		}
	}

	for i, o := range s.gameObjects {
		if b, ok := o.(Behaviour); ok && s.deleted[i] == nil {
			b.Update(s)
		}
	}
}

// Render はシーンの描画を行います。
func (s *Scene) Render() {

	if s.Renderer == nil {
		return
	}

	for i, o := range s.gameObjects {
		if s.deleted[i] == nil {
			s.Renderer.Render(s.MainCamera, o)
		}
	}
}
