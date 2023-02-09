package repository

import (
	"time"

	"github.com/dukesp69/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InserReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityByDatesForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}
