package engine

// EventListener はイベントをハンドルするためのリスナーです。
type EventListener interface {
	// イベントが発生したことを伝える
	Notify(args ...interface{})
}

// EventListenerFunc は関数自体に EventListener を実装させるための型です。
type EventListenerFunc func(args ...interface{})

// Notify は EventListener.Notifyの実装です。
func (l EventListenerFunc) Notify(args ...interface{}) {
	l(args...)
}

// EventListenerList は EventListenerをまとめて扱うための型です。
// EventListenerを実装します。
type EventListenerList []EventListener

// Notify は EventListener.Notifyの実装です。
func (l EventListenerList) Notify(args ...interface{}) {
	for _, elm := range l {
		EventNotify(elm, args...)
	}
}

// EventNotify は l が nil の時には Notify しないヘルパー関数です。j
func EventNotify(l EventListener, args ...interface{}) {
	if l != nil {
		l.Notify(args...)
	}
}
