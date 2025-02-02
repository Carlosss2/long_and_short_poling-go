package application

import "sync"

type ProductEvent struct {
	mu        sync.Mutex
	listeners []chan struct{}
}

var productEvent = &ProductEvent{}

func (e *ProductEvent) Notify() {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.listeners {
		ch <- struct{}{} // Envía un mensaje para notificar cambios
	}
	e.listeners = nil // Limpia la lista de listeners
}

func (e *ProductEvent) Wait() <-chan struct{} {
	e.mu.Lock()
	defer e.mu.Unlock()
	ch := make(chan struct{}, 1) // Canal con buffer para evitar bloqueos
	e.listeners = append(e.listeners, ch)
	return ch
}

func NotifyProductUpdate() {
	productEvent.Notify()
}

func WaitForProductUpdate() <-chan struct{} {
	return productEvent.Wait()
}