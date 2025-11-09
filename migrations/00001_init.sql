-- +goose Up
-- PRAGMA journal_mode = WAL;
PRAGMA foreign_keys = ON;

-- PRAGMA synchronous = NORMAL;
PRAGMA cache_size = -64000;

PRAGMA temp_store = MEMORY;
