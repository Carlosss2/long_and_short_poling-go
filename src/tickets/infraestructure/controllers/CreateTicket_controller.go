package controllers

import (
	"net/http"
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/domain"
	"projectC1/src/tickets/infraestructure"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	// Canal para Long Polling con búfer para evitar bloqueo
	ticketUpdates = make(chan []domain.Ticket, 1)
	// Mutex para evitar condiciones de carrera
	mu sync.Mutex
)

type CreateTicketController struct {
	useCaseCreate *application.CreateTicket
	repository    *infraestructure.MySQL // Acceso directo a la capa de infraestructura
}

// Constructor con inicialización de repository
func NewCreateTicketController(useCaseCreate *application.CreateTicket, repository *infraestructure.MySQL) *CreateTicketController {
	return &CreateTicketController{
		useCaseCreate: useCaseCreate,
		repository:    repository,
	}
}

func (createTicket *CreateTicketController) Create(c *gin.Context) {
	var ticket domain.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createTicket.useCaseCreate.Execute(ticket.Client, ticket.Total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Validar que repository no sea nil antes de llamar a GetAll
	if createTicket.repository == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository no inicializado"})
		return
	}

	// Obtener los tickets directamente desde la base de datos
	mu.Lock()
	tickets, err := createTicket.repository.GetAll()
	mu.Unlock()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo tickets"})
		return
	}

	// Enviar de forma no bloqueante al canal
	select {
	case ticketUpdates <- tickets:
	default:
		// Si nadie está escuchando, no bloquea la ejecución
	}
	application.NotifyTicketUpdate()
	c.JSON(http.StatusCreated, gin.H{"message": "Ticket registrado"})
}
