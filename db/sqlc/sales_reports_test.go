package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"AnalyticsAndReporting/util"
)

func CreateRandomSalesReports(t *testing.T) SalesReport {
	data := CreateSalesReportsParams{
		SReportID:   util.CreateUUID(),
		ReportType:  util.GenerateRandomSalesReportType(),
		StartDate:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		EndDate:     util.GenerateRandomDate(),
		TotalSales:  util.GenerateRandomNumeric(),
		TotalOrders: util.GenerateRandomInt4(),
	}

	SalesReport, err := testStore.CreateSalesReports(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, SalesReport.SReportID)
	require.NotEmpty(t, data.StartDate, SalesReport.StartDate)
	require.NotEmpty(t, data.EndDate, SalesReport.EndDate)
	require.Equal(t, data.ReportType, SalesReport.ReportType)
	require.Equal(t, data.TotalSales, SalesReport.TotalSales)
	require.Equal(t, data.TotalOrders, SalesReport.TotalOrders)
	require.NotZero(t, SalesReport.CreatedAt)

	return SalesReport
}

func TestCreateSalesReports(t *testing.T) {
	CreateRandomSalesReports(t)
}

func TestGetSalesReportsList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomSalesReports(t)
	}

	DataList, err := testStore.GetSalesReportsList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, DataList)
	require.GreaterOrEqual(t, len(DataList), 10)
}

func TestGetSalesReportsByDate(t *testing.T) {
	data := CreateRandomSalesReports(t)
	newT := data.StartDate.Time.Add(2 * time.Minute)

	Data, err := testStore.GetSalesReportsByDate(context.Background(), pgtype.Timestamp{Time: newT, Valid: true})
	require.NoError(t, err)
	require.NotEmpty(t, Data)
}

func TestGetSalesReportsById(t *testing.T) {
	data := CreateRandomSalesReports(t)

	Data, err := testStore.GetSalesReportsById(context.Background(), data.SReportID)
	require.NoError(t, err)
	require.NotEmpty(t, Data)
}
