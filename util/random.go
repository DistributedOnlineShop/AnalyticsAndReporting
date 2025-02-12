package util

import (
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateRandomSalesReportType() string {
	reportTypes := []string{"Daily Sales", "Weekly Sales", "Monthly Sales", "Quarterly Sales", "Annual Sales"}
	return reportTypes[rand.IntN(len(reportTypes))]
}

func GenerateRandomUserReportType() string {
	reportTypes := []string{"New User Report", "Active User Report", "User Retention Report", "User Acquisition Report", "User Demographics Report"}
	return reportTypes[rand.IntN(len(reportTypes))]
}

func GenerateRandomDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateRandomNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateRandomInt4() pgtype.Int4 {
	return pgtype.Int4{
		Int32: rand.Int32N(500) + 10,
		Valid: true,
	}
}
