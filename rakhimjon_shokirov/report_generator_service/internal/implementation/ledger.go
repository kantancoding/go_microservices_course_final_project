package implementation

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	query = `SELECT SUM(amount) AS total_amount FROM ledger WHERE DATE(transaction_date) = ?`
)

func GetDailyAmount(ctx context.Context, ledgerDB *sql.DB) (int64, error) {
	var (
		today       = time.Now().Format("2006-01-02")
		dailyAmount int64
	)

	err := ledgerDB.QueryRowContext(ctx, query, today).Scan(&dailyAmount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}

	return dailyAmount, nil
}
