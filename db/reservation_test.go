package db

import (
	"context"
	"database/sql"
	"reservation-service/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomReservation(t *testing.T) Reservation {
	arg := CreateReservationParam{
		Station_id: util.RandomInt(1261, 654561),
		User_id:    util.RandomInt(1261, 654561),
		Start:      util.RandomTimestamp(),
		End:        util.RandomTimestamp(),
	}

	result, err := testStore.Create(context.Background(), arg)

	// Check if method executed correctly.
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Station_id, result.Station_id)
	require.Equal(t, arg.User_id, result.User_id)
	require.Equal(t, arg.Start.Round(time.Minute).Format("2006-01-02 15:04:05"), result.Start.Round(time.Minute).Format("2006-01-02 15:04:05"))
	require.Equal(t, arg.End.Round(time.Minute).Format("2006-01-02 15:04:05"), result.End.Round(time.Minute).Format("2006-01-02 15:04:05"))

	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)

	return result
}

func TestCreateReservation(t *testing.T) {
	createRandomReservation(t)
}

func TestGetReservation(t *testing.T) {
	reservation1 := createRandomReservation(t)
	reservation2, err := testStore.GetByID(context.Background(), reservation1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, reservation2)

	require.Equal(t, reservation1.ID, reservation2.ID)
	require.Equal(t, reservation1.Station_id, reservation2.Station_id)
	require.Equal(t, reservation1.User_id, reservation2.User_id)
	require.Equal(t, reservation1.Start.Round(time.Minute).Format("2006-01-02 15:04:05"), reservation2.Start.Round(time.Minute).Format("2006-01-02 15:04:05"))
	require.Equal(t, reservation1.End.Round(time.Minute).Format("2006-01-02 15:04:05"), reservation2.End.Round(time.Minute).Format("2006-01-02 15:04:05"))
	require.Equal(t, reservation1.CreatedAt, reservation2.CreatedAt)
}

func TestListReservations(t *testing.T) {

	// Create a list of reservations in database.
	var createdReservations [10]Reservation
	for i := 0; i < 10; i++ {
		createdReservations[i] = createRandomReservation(t)
	}

	arg := ListReservationParam{
		Limit:  10,
		Offset: 0,
	}

	// Retrieve list of reservations.
	reservations, err := testStore.GetAll(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservations)

	for _, u := range reservations {
		require.NotEmpty(t, u)
	}
}

func TestUpdateReservation(t *testing.T) {

	reservation1 := createRandomReservation(t)

	arg := UpdateReservationParam{
		Station_id: util.RandomInt(1261, 654561),
		User_id:    util.RandomInt(1261, 654561),
		Start:      util.RandomTimestamp(),
		End:        util.RandomTimestamp(),
	}

	reservation2, err := testStore.Update(context.Background(), arg, reservation1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, reservation2)

	require.Equal(t, reservation1.ID, reservation2.ID)
	require.Equal(t, arg.Station_id, reservation2.Station_id)
	require.Equal(t, arg.User_id, reservation2.User_id)
	require.Equal(t, arg.Start.Round(time.Minute).Format("2006-01-02 15:04:05"), reservation2.Start.Round(time.Minute).Format("2006-01-02 15:04:05"))
	require.Equal(t, arg.End.Round(time.Minute).Format("2006-01-02 15:04:05"), reservation2.End.Round(time.Minute).Format("2006-01-02 15:04:05"))
	require.Equal(t, reservation1.CreatedAt, reservation2.CreatedAt)

}

func TestDeleteReservation(t *testing.T) {
	reservation1 := createRandomReservation(t)
	err := testStore.Delete(context.Background(), reservation1.ID)
	require.NoError(t, err)

	reservation2, err := testStore.GetByID(context.Background(), reservation1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, reservation2)
}
