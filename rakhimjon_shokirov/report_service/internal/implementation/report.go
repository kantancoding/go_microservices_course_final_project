package implementation

import (
	"context"
	"database/sql"
)

func Insert(ctx context.Context, reportDB *sql.DB, dailyAmount int64) error {
	stmt, err := reportDB.PrepareContext(ctx, "insert into report (daily_amount) values(?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dailyAmount)
	return err
}
