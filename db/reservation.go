package db

import (
	"context"
	"time"
)

type Reservation struct {
	ID         int64     `json:"reservation_id" db:"reservation_id"`
	Station_id int64     `json:"station_id" db:"station_id"`
	User_id    int64     `json:"user_id" db:"user_id"`
	Start      time.Time `json:"start" db:"start"`
	End        time.Time `json:"end" db:"end"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type CreateReservationParam struct {
	Station_id int64
	User_id    int64
	Start      time.Time
	End        time.Time
}

type UpdateReservationParam struct {
	Station_id int64
	User_id    int64
	Start      time.Time
	End        time.Time
}

type ListReservationParam struct {
	Offset int32
	Limit  int32
}

func (store *Store) GetByID(ctx context.Context, id int64) (reservation Reservation, err error) {
	const query = `SELECT * FROM "reservations" WHERE "reservation_id" = $1`
	err = store.db.GetContext(ctx, &reservation, query, id)

	return
}

func (store *Store) GetAll(ctx context.Context, arg ListReservationParam) (reservations []Reservation, err error) {
	const query = `SELECT * FROM "reservations" OFFSET $1 LIMIT $2`
	reservations = []Reservation{}
	err = store.db.SelectContext(ctx, &reservations, query, arg.Offset, arg.Limit)

	return
}

func (store *Store) Create(ctx context.Context, arg CreateReservationParam) (Reservation, error) {
	const query = `
	INSERT INTO "reservations"("station_id", "user_id", "start", "end") 
	VALUES ($1, $2, $3, $4)
	RETURNING "reservation_id", "station_id", "user_id", "start", "end", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, arg.Station_id, arg.User_id, arg.Start, arg.End)

	var reservation Reservation

	err := row.Scan(
		&reservation.ID,
		&reservation.Station_id,
		&reservation.User_id,
		&reservation.Start,
		&reservation.End,
		&reservation.CreatedAt,
	)

	return reservation, err
}

func (store *Store) Update(ctx context.Context, arg UpdateReservationParam, id int64) (Reservation, error) {
	const query = `
	UPDATE "reservations"
	SET "station_id" = $2,
		"user_id" = $3,
		"start" = $4,
		"end" = $5
	WHERE "reservation_id" = $1
	RETURNING "reservation_id", "station_id", "user_id", "start", "end", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, id, arg.Station_id, arg.User_id, arg.Start, arg.End)

	var reservation Reservation

	err := row.Scan(
		&reservation.ID,
		&reservation.Station_id,
		&reservation.User_id,
		&reservation.Start,
		&reservation.End,
		&reservation.CreatedAt,
	)

	return reservation, err
}

func (store *Store) Delete(ctx context.Context, id int64) error {
	const query = `
	DELETE FROM reservations
	WHERE "reservation_id" = $1
	`
	_, err := store.db.ExecContext(ctx, query, id)

	return err
}

func (store *Store) GetAllByUserID(ctx context.Context, userID int64) (reservations []Reservation, err error) {
	const query = `SELECT * FROM "reservations" WHERE "user_id" = $1 ORDER BY "start" DESC`
	reservations = []Reservation{}
	err = store.db.SelectContext(ctx, &reservations, query, userID)

	return
}
