package dbrepo

import (
	"errors"
	"time"

	"github.com/dukesp69/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {

	return true
}

// InserReservation insert a reservation into database
func (m *testDBRepo) InserReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// InsertRoomRestriction insert a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// SearchAvailabilityByDatesByRoomID return true if availabilty exits for roomID and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil

}

// SearchAvailabilityByDatesForAllRooms return a slice of available rooms, if any for a given date range
func (m *testDBRepo) SearchAvailabilityByDatesForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil

}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {

	var res models.Reservation

	return res, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {

	return nil
}
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}
