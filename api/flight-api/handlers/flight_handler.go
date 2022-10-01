package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/xuanlt1991/flight-booking/api/flight-api/requests"
	"github.com/xuanlt1991/flight-booking/api/flight-api/responses"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightApiHandler struct {
	flightClient pb.FlightServiceClient
}

type IFlightApiHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	ViewFlight(c *gin.Context)
	SearchFlight(c *gin.Context)
}

func NewFlightApiHandler(flightClient pb.FlightServiceClient) IFlightApiHandler {
	return &FlightApiHandler{
		flightClient: flightClient,
	}
}

func (h *FlightApiHandler) CreateFlight(c *gin.Context) {
	req := requests.FlightRequest{}
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
	log.Printf("Departure Date: %T\n", req.DepatureDate)
	log.Printf("Departure Time: %T\n", req.DepartureTime)
	pReq := &pb.FlightRequest{}
	err := copier.Copy(&pReq, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	CopyReqToPReq(&req, pReq)
	log.Printf("request: %v\n - pReq: %v\n", req, pReq)

	log.Println("Start calling gRPC create flight method")
	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	res := ToApiResponse(pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": res,
	})

}

func (h *FlightApiHandler) UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	req := requests.FlightRequest{}
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

	pReq := &pb.FlightRequest{}

	err := copier.Copy(&pReq, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	pReq.Id = id
	CopyReqToPReq(&req, pReq)

	pRes, err := h.flightClient.UpdateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	res := ToApiResponse(pRes)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": res,
	})
}

func (h *FlightApiHandler) ViewFlight(c *gin.Context) {
	id := c.Param("id")

	pReq := &pb.ViewFlightRequest{
		Id: id,
	}

	pRes, err := h.flightClient.ViewFlight(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	res := ToApiResponse(pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": res,
	})
}

func (h *FlightApiHandler) SearchFlight(c *gin.Context) {

	req := requests.SearchFlightRequest{}

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

	pReq := &pb.SearchFlightRequest{}

	err := copier.Copy(&pReq, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	pReq.DepartureDate = ConvertDate(req.DepatureDate)
	pReq.ArrivalDate = ConvertDate(req.ArrivalDate)

	pRes, err := h.flightClient.SearchFlight(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dtos := make([]*responses.FlightResponse, 0)

	for _, v := range pRes.Flights {
		dto := ToApiResponse(v)

		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtos,
	})
}

func ToApiResponse(pRes *pb.FlightResponse) *responses.FlightResponse {
	res := &responses.FlightResponse{
		Id:            pRes.Flight.Id,
		Name:          pRes.Flight.Name,
		From:          pRes.Flight.From,
		To:            pRes.Flight.To,
		Status:        pRes.Flight.Status,
		AvailableSlot: pRes.Flight.AvailableSlot,
		DepatureDate:  time.Date(int(pRes.Flight.DepartureDate.Year), time.Month(pRes.Flight.DepartureDate.Month), int(pRes.Flight.DepartureDate.Day), 0, 0, 0, 0, time.UTC),
		ArrivalDate:   time.Date(int(pRes.Flight.ArrivalDate.Year), time.Month(pRes.Flight.ArrivalDate.Month), int(pRes.Flight.ArrivalDate.Day), 0, 0, 0, 0, time.UTC),
		DepartureTime: pRes.Flight.DepartureTime.AsTime(),
		ArrivalTime:   pRes.Flight.ArrivalTime.AsTime(),
		CreatedAt:     pRes.Flight.Audit.CreatedAt.AsTime(),
		ModifiedAt:    pRes.Flight.Audit.ModifiedAt.AsTime(),
	}

	return res
}

func CopyReqToPReq(req *requests.FlightRequest, pReq *pb.FlightRequest) {
	pReq.DepartureDate = ConvertDate(req.DepatureDate)
	pReq.ArrivalDate = ConvertDate(req.ArrivalDate)
	pReq.ArrivalTime = timestamppb.New(req.ArrivalTime)
	pReq.DepartureTime = timestamppb.New(req.DepartureTime)
}

func ConvertDate(date time.Time) *pb.Date {
	return &pb.Date{
		Year:  int32(date.Year()),
		Month: int32(date.Month()),
		Day:   int32(date.Day()),
	}
}
