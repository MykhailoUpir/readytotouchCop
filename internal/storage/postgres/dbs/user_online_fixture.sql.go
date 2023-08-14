// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_online_fixture.sql

package dbs

import (
	"context"
)

const userOnlineFixtureCount = `-- name: UserOnlineFixtureCount :one
SELECT COUNT(*) AS total,
       SUM(
               CASE online
                   WHEN TO_TIMESTAMP($1::BIGINT) THEN 0
                   ELSE 1
                   END
           )    AS changed
FROM user_online
`

type UserOnlineFixtureCountRow struct {
	Total   int64
	Changed int64
}

func (q *Queries) UserOnlineFixtureCount(ctx context.Context, online int64) (UserOnlineFixtureCountRow, error) {
	row := q.db.QueryRow(ctx, userOnlineFixtureCount, online)
	var i UserOnlineFixtureCountRow
	err := row.Scan(&i.Total, &i.Changed)
	return i, err
}

const userOnlineFixtureUpsert = `-- name: UserOnlineFixtureUpsert :exec
INSERT INTO user_online (user_id, online)
SELECT generate_series,
       to_timestamp($1::BIGINT)
FROM generate_series(1, $2::BIGINT)
ON CONFLICT (user_id) DO UPDATE
    SET online = excluded.online
`

type UserOnlineFixtureUpsertParams struct {
	Online int64
	Count  int64
}

func (q *Queries) UserOnlineFixtureUpsert(ctx context.Context, arg UserOnlineFixtureUpsertParams) error {
	_, err := q.db.Exec(ctx, userOnlineFixtureUpsert, arg.Online, arg.Count)
	return err
}
