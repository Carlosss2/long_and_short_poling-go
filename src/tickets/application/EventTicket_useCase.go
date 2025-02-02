package application

import "sync"

type TicketEvent struct {
	mu        sync.Mutex
	listeners []chan struct{}
}

var ticketEvent = &TicketEvent{}

func (e *TicketEvent) Notify() {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.listeners {
		ch <- struct{}{} // Envía un mensaje en lugar de cerrar el canal
	}
	e.listeners = nil // Limpia la lista de listeners después de notificar
}

func (e *TicketEvent) Wait() <-chan struct{} {
	e.mu.Lock()
	defer e.mu.Unlock()
	ch := make(chan struct{}, 1) // Usa un canal con buffer para evitar bloqueos
	e.listeners = append(e.listeners, ch)
	return ch
}

func NotifyTicketUpdate() {
	ticketEvent.Notify()
}

func WaitForTicketUpdate() <-chan struct{} {
	return ticketEvent.Wait()
}