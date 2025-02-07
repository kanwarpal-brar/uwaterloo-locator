package handler

import (
	"net/http"
	"strconv"

	"washroom-data-service/models"
	"washroom-data-service/service"

	"github.com/gin-gonic/gin"
)

type WashroomHandler struct {
	washroomService service.WashroomService
}

func NewWashroomHandler(s service.WashroomService) *WashroomHandler {
	return &WashroomHandler{washroomService: s}
}

// Create godoc
// @Summary Create a new washroom
// @Description Add a new washroom to the system
// @Accept json
// @Produce json
// @Param washroom body models.Washroom true "Washroom information"
// @Success 201 {object} models.Washroom
// @Router /washrooms [post]
func (h *WashroomHandler) Create(c *gin.Context) {
	var washroom models.Washroom
	if err := c.ShouldBindJSON(&washroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.washroomService.Create(c.Request.Context(), &washroom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, washroom)
}

// GetByID godoc
// @Summary Get a washroom by ID
// @Description Retrieve a washroom's details by its ID
// @Produce json
// @Param id path string true "Washroom ID"
// @Success 200 {object} models.Washroom
// @Router /washrooms/{id} [get]
func (h *WashroomHandler) GetByID(c *gin.Context) {
	washroom, err := h.washroomService.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Washroom not found"})
		return
	}
	c.JSON(http.StatusOK, washroom)
}

// FindNearby godoc
// @Summary Find nearby washrooms
// @Description Find washrooms within a specified radius of a location
// @Produce json
// @Param lat query number true "Latitude"
// @Param lng query number true "Longitude"
// @Param radius query number true "Radius in meters"
// @Success 200 {array} models.Washroom
// @Router /washrooms/nearby [get]
func (h *WashroomHandler) FindNearby(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)
	radius, _ := strconv.ParseFloat(c.Query("radius"), 64)

	washrooms, err := h.washroomService.FindNearby(c.Request.Context(), lat, lng, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, washrooms)
}
