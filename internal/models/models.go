package models

import (
	"time"
)

// Reservation holds reservation data
// type Reservation struct {
// 	FirstName string
// 	LastName  string
// 	Email     string
// 	Phone     string
// }

// User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Restriction is the restiction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdateAt        time.Time
}

// Reservation is the  reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdateAt  time.Time
	Processed int
	Room      Room
}

// RoomRestriction is the room restiction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdateAt      time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Reservation
}

// MailData holds an email message
type MailData struct {
	To      string
	From    string
	Subject string
	Content string
}
