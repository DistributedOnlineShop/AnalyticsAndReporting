package db

import (
	"context"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"AnalyticsAndReporting/util"
)

func CreateRandomUserReport(t *testing.T) UserReport {
	data := CreateUserReportParams{
		UReportID:  util.CreateUUID(),
		UserID:     util.CreateUUID(),
		ReportType: util.GenerateUserReportType(),
		TotalValue: util.GenerateNumeric(),
	}

	UserReport, err := testStore.CreateUserReport(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, UserReport.UReportID)
	require.Equal(t, data.UserID, UserReport.UserID)
	require.Equal(t, data.ReportType, UserReport.ReportType)
	require.Equal(t, data.TotalValue, UserReport.TotalValue)
	require.NotZero(t, UserReport.CreatedAt)

	return UserReport
}

func TestCreateUserReport(t *testing.T) {
	CreateRandomUserReport(t)
}

func TestGetUserReportById(t *testing.T) {
	data := CreateRandomUserReport(t)

	UserReport, err := testStore.GetUserReportById(context.Background(), data.UReportID)
	require.NoError(t, err)
	require.NotEmpty(t, UserReport)
}

func TestGetUserReportByTotalValue(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUserReport(t)
	}

	totalValue := pgtype.Numeric{
		Int:   big.NewInt(100),
		Exp:   -2,
		Valid: true,
	}

	UserReport, err := testStore.GetUserReportByOverTotalValue(context.Background(), totalValue)
	require.NoError(t, err)
	require.NotEmpty(t, UserReport)
}

func TestGetUserReportByUserId(t *testing.T) {
	data := CreateRandomUserReport(t)

	UserReport, err := testStore.GetUserReportByUserId(context.Background(), data.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, UserReport)
}
