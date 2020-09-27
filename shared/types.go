package shared

import (
	"github.com/gwuah/api/database/models"
	"github.com/uber/h3-go"
)

type Meta struct {
	Type string `json:"type"`
}

type Coord struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
}

type User struct {
	Id             string     `json:"id"`
	LastKnownIndex h3.H3Index `json:"lastKnownIndex"`
	Coord
}

type UserLocationUpdate struct {
	Id string `json:"id"`
	Coord
}

type DeliveryRequest struct {
	Meta        Meta   `json:"meta"`
	Origin      Coord  `json:"origin"`
	Destination Coord  `json:"destination"`
	ProductId   uint   `json:"productId"`
	Notes       string `json:"notes"`
	CustomerID  uint   `json:"customerId"`
}

type CancelDeliveryRequest struct {
	Meta   Meta `json:"meta"`
	TripId uint `json:"tripId"`
}

type GetClosestElectronsRequest struct {
	Meta   Meta   `json:"meta"`
	Id     string `json:"id"`
	Origin Coord  `json:"origin"`
}

type NewDeliveryOrder struct {
	Meta     Meta             `json:"meta"`
	Delivery *models.Delivery `json:"delivery"`
}

type AcceptDeliveryRequest struct {
	Meta       Meta `json:"meta"`
	DeliveryId uint `json:"deliveryId"`
}