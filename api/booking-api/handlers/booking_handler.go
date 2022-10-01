package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/xuanlt1991/flight-booking/api/booking-api/requests"
	"github.com/xuanlt1991/flight-booking/api/booking-api/responses"
	"github.com/xuanlt1991/flight-booking/pb"
)

type BookingApiHandler struct {
	bookingClient pb.BookingServiceClient
}

type IBookingApiHandler interface {
	CreateBooking(c *gin.Context)
	ViewBooking(c *gin.Context)
	BookingHistory(c *gin.Context)
	CancelBooking(c *gin.Context)
}

func NewBookingApiHandler(bookingClient pb.BookingServiceClient) IBookingApiHandler {
	return &BookingApiHandler{
		bookingClient: bookingClient,
	}
}

func (h *BookingApiHandler) CreateBooking(c *gin.Context) {
	req := requests.BookingRequest{}
	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.BookingRequest{}

	err := copier.Copy(&pReq, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pRes, err := h.bookingClient.CreateBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	res := ToApiResponse(pRes.Booking)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": res,
	})
}

func (h *BookingApiHandler) ViewBooking(c *gin.Context) {
	id := c.Param("id")

	pReq := &pb.ViewBookingRequest{
		Id: id,
	}

	pRes, err := h.bookingClient.ViewBooking(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	res := ToApiResponse(pRes.Booking)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": res,
	})
}
func (h *BookingApiHandler) BookingHistory(c *gin.Context) {
	id := c.Param("customer_id")
	pReq := &pb.ViewBookingHistoryRequest{
		CustomerId: id,
	}

	pRes, err := h.bookingClient.BookingHistory(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dtos := make([]*responses.BookingResponse, 0)

	for _, v := range pRes.Bookings {
		dto := ToApiResponse(v.Booking)

		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtos,
	})
}

func (h *BookingApiHandler) CancelBooking(c *gin.Context) {
	id := c.Param("id")

	pReq := &pb.CancelBookingRequest{
		Id: id,
	}

	_, err := h.bookingClient.CancelBooking(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": id,
	})
}

func ToApiResponse(pRes *pb.Booking) *responses.BookingResponse {
	res := &responses.BookingResponse{
		Id:          pRes.Id,
		CustomerId:  pRes.CustomerId,
		FlightId:    pRes.FlightId,
		BookingCode: pRes.BookingCode,
		Status:      pRes.Status,
		BookedDate:  pRes.BookedDate.AsTime(),
		CreatedAt:   pRes.Audit.CreatedAt.AsTime(),
		ModifiedAt:  pRes.Audit.ModifiedAt.AsTime(),
	}

	return res
}
