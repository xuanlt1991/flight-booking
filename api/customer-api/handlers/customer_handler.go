package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/xuanlt1991/flight-booking/api/customer-api/requests"
	"github.com/xuanlt1991/flight-booking/api/customer-api/responses"
	"github.com/xuanlt1991/flight-booking/pb"
)

type CustomerApiHandler struct {
	customerClient pb.CustomerServiceClient
}

type ICustomerApiHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
	ViewCustomer(c *gin.Context)
}

func NewCustomerApiHandler(customerClient pb.CustomerServiceClient) ICustomerApiHandler {
	return &CustomerApiHandler{
		customerClient: customerClient,
	}
}

func (h *CustomerApiHandler) CreateCustomer(c *gin.Context) {
	req := requests.CustomerRequest{}
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

	pReq := &pb.CustomerRequest{}

	err := copier.Copy(&pReq, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pRes, err := h.customerClient.CreateCustomer(c.Request.Context(), pReq)
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

func (h *CustomerApiHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	req := requests.CustomerRequest{}
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

	pReq := &pb.CustomerRequest{}

	err := copier.Copy(&pReq, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pReq.Id = id

	pRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), pReq)
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

func (h *CustomerApiHandler) ChangePassword(c *gin.Context) {
	id := c.Param("id")
	req := requests.ChangePasswordRequest{}

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

	pReq := &pb.ChangePasswordRequest{}

	err := copier.Copy(&pReq, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	pReq.Id = id
	pRes, err := h.customerClient.ChangePassword(c.Request.Context(), pReq)

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

func (h *CustomerApiHandler) ViewCustomer(c *gin.Context) {
	id := c.Param("id")

	pReq := &pb.ViewCustomerRequest{
		Id: id,
	}

	pRes, err := h.customerClient.ViewCustomer(c.Request.Context(), pReq)

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

func ToApiResponse(pRes *pb.CustomerResponse) *responses.CustomerResponse {
	res := &responses.CustomerResponse{
		Id:          pRes.Customer.Id,
		FirstName:   pRes.Customer.FirstName,
		LastName:    pRes.Customer.LastName,
		Address:     pRes.Customer.Address,
		License:     pRes.Customer.License,
		PhoneNumber: pRes.Customer.PhoneNumber,
		Email:       pRes.Customer.Email,
		Status:      pRes.Customer.Status,
		CreatedAt:   pRes.Customer.Audit.CreatedAt.AsTime(),
		ModifiedAt:  pRes.Customer.Audit.ModifiedAt.AsTime(),
	}

	return res
}
