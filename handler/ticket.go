package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"noctiket/constant"
	"noctiket/model/entity"
	"noctiket/model/request"
	"noctiket/repository"
	"noctiket/response"
	"noctiket/util"
	"time"
)

func CreateTicket(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user, _ := userInterface.(entity.User)

	var ticketRequest entity.Ticket
	err := c.ShouldBindJSON(&ticketRequest)
	if err != nil {
		log.Println(err.Error())
		response.MapResponseByError(c, err)
		return
	}

	log.Println(user.Id)

	ticketId, err := repository.GenerateTicketId()
	ticket := entity.Ticket{
		TicketId:    ticketId,
		Title:       ticketRequest.Title,
		Description: ticketRequest.Description,
		Status:      constant.Open,
		Priority:    ticketRequest.Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   user.Email,
	}
	logTicket := util.LogTicket(ticket, user, constant.Open)
	ticket.Log = append(ticket.Log, logTicket)

	err = repository.SaveTicket(ticket)
	if err != nil {
		response.MapResponseByError(c, err)
	}

	response.SuccessResponse(c, ticket)
}

func GetTickets(c *gin.Context) {
	var requestPayload request.TicketRequest
	err := c.ShouldBindQuery(&requestPayload)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	tickets, err := repository.GetTickets(requestPayload)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	response.SuccessResponse(c, tickets)
	return
}

func AssignTicket(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user, _ := userInterface.(entity.User)

	var assignTicketReq request.AssignTicket

	err := c.ShouldBindJSON(&assignTicketReq)
	if err != nil {
		response.MapResponseByError(c, err)
	}

	assignee, err := repository.GetUserByEmail(assignTicketReq.Assignee)

	if err != nil {
		log.Printf("Assignee %s not found\n", assignTicketReq.Assignee)
		response.MapResponseByError(c, err)
		return
	}

	if !constant.Engineers.Contains(entity.Role(assignee.Role)) {
		log.Printf("% is non Engineer, cannot be assigned")
		response.ErrorInvalidRequest(c)
		return
	}

	ticket, err := repository.GetTicketById(assignTicketReq.TicketId)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	ticket.AssignedTo = assignTicketReq.Assignee
	ticket.Status = constant.OnProgress
	ticket.UpdatedAt = time.Now()

	logTicket := util.LogTicket(ticket, user, constant.OnProgress)
	ticket.Log = append(ticket.Log, logTicket)
	err = repository.UpdateTicket(ticket)

	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	response.SuccessResponse(c, nil)
	return
}
