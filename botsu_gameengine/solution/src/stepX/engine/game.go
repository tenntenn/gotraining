package engine

import "time"

// 各種イベントリスナ
type gameEventListeners struct {
	OnChangeScene EventListener
	BeforeStart   EventListener
	AfterStart    EventListener
	BeforeUpdate  EventListener
	AfterUpdate   EventListener
	BeforeRender  EventListener
	AfterRender   EventListener
}

// Game は1つのゲームを表す構造体です。
type Game struct {
	// イベントリスナ
	gameEventListeners
	// SceneMakerのリスト
	SceneFactories []SceneFactory
	// FPS
	FPS int
	// 現在のシーンのインデックス
	current int
	// 現在のシーン
	scene *Scene
}

// New は新しい Game を作ります。
func New() *Game {
	return &Game{
		SceneFactories: []SceneFactory{},
		FPS:            60,
	}
}

// Start はゲームループを開始します。
func (g *Game) Start() {

	// 初期シーン
	s := g.ChangeScene(0)

	// シーンの初期化処理
	EventNotify(g.BeforeStart, s)
	s.Start()
	EventNotify(g.AfterStart, s)

	// ゲームループ
	for {
		startTime := time.Now()

		// シーンのUpdate
		EventNotify(g.BeforeUpdate, g.Scene())
		g.Scene().Update()
		EventNotify(g.AfterUpdate, g.Scene())

		// シーンの描画
		EventNotify(g.BeforeRender, g.Scene())
		s.Render()
		EventNotify(g.AfterRender, g.Scene())

		// 次のフレームまで停止
		d := time.Since(startTime)
		time.Sleep(time.Duration(1000/float64(g.FPS))*time.Millisecond - d)
	}
}

// Scene は現在のシーンを取得します。
func (g *Game) Scene() *Scene {
	if g.scene != nil {
		return g.scene
	}

	if g.current < len(g.SceneFactories)-1 ||
		g.current > len(g.SceneFactories) {
		return nil
	}

	g.scene = g.SceneFactories[g.current].Create()

	return g.scene
}

// ChangeScene はシーンを変更します。
func (g *Game) ChangeScene(i int) *Scene {
	if i < 0 || i > len(g.SceneFactories) {
		return nil
	}

	g.current = i
	EventNotify(g.OnChangeScene, g.Scene())
	return g.Scene()
}
