-- name: InsertSymbol :exec
INSERT INTO symbols (created_ts, name, symbol, exchange, asset_type, status, deleted_ts)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (symbol) DO UPDATE
SET created_ts = EXCLUDED.created_ts,
    name = EXCLUDED.name,
    exchange = EXCLUDED.exchange,
    asset_type = EXCLUDED.asset_type,
    status = EXCLUDED.status,
    deleted_ts = EXCLUDED.deleted_ts;


-- name: ListSymbols :many
SELECT id, created_ts, name, symbol, exchange, asset_type, status
FROM symbols
WHERE deleted_ts IS NULL;