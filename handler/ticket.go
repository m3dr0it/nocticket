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

	logTicket := util.LogTicketCreated(ticketId, user.Email)
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
		log.Printf("%s is non Engineer, cannot be assigned", assignTicketReq.Assignee)
		response.ErrorInvalidRequest(c)
		return
	}

	ticket, err := repository.GetTicketById(assignTicketReq.TicketId)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	changes := false

	if ticket.Status != constant.OnProgress {
		logTicketStatus := util.LogUpdateStatus(ticket.TicketId, user.Email, ticket.Status, constant.OnProgress)
		ticket.Status = constant.OnProgress
		ticket.Log = append(ticket.Log, logTicketStatus)
		ticket.UpdatedAt = time.Now()
		changes = true
	}

	if ticket.AssignedTo != assignTicketReq.Assignee {
		ticket.AssignedTo = assignTicketReq.Assignee
		logTicketAssign := util.LogAssignTicket(ticket.TicketId, user.Email, assignTicketReq.Assignee)
		ticket.Log = append(ticket.Log, logTicketAssign)
		ticket.UpdatedAt = time.Now()
		changes = true
	}

	if changes {
		err = repository.UpdateTicket(ticket)

		if err != nil {
			response.MapResponseByError(c, err)
			return
		}
	}

	response.SuccessResponse(c, nil)
	return
}

func UpdateProgress(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user, _ := userInterface.(entity.User)

	var updateProgressReq request.TicketProgressRequest
	err := c.ShouldBindJSON(&updateProgressReq)

	if err != nil {
		log.Println(err.Error())
		response.MapResponseByError(c, err)
	}

	ticket, err := repository.GetTicketById(updateProgressReq.TicketId)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	if user.Role != constant.RoleAdmin && !constant.Engineers.Contains(user.Role) {
		log.Printf("%s is non Engineer, cannot update ticket \n", updateProgressReq.TicketId)
		response.UnauthorizedResponse(c)
		return
	}

	if ticket.Status != updateProgressReq.Status {
		logTicket := util.LogUpdateStatus(ticket.TicketId, user.Email, ticket.Status, updateProgressReq.Status)
		ticket.Log = append(ticket.Log, logTicket)
	}

	var logTicket entity.LogEntry

	if updateProgressReq.Status == constant.Close {
		logTicket = util.LogCloseTicket(ticket.TicketId, user.Email, updateProgressReq.Message)
		ticket.Resolution = updateProgressReq.Message
		ticket.ResolvedAt = time.Now()
	} else {
		logTicket = util.LogUpdateProgress(ticket.TicketId, user.Email, updateProgressReq.Message)
	}

	ticket.Status = updateProgressReq.Status
	ticket.Log = append(ticket.Log, logTicket)
	err = repository.UpdateTicket(ticket)
	if err != nil {
		response.MapResponseByError(c, err)
	}

	response.SuccessResponse(c, nil)
	return
}
