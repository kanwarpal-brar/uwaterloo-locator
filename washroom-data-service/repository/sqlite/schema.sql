CREATE TABLE IF NOT EXISTS washrooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,
    location_updates INTEGER NOT NULL DEFAULT 1,
    building TEXT NOT NULL,
    floor INTEGER NOT NULL,
    gender TEXT NOT NULL,
    is_accessible BOOLEAN NOT NULL,
    version INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_washrooms_building ON washrooms(building);
CREATE INDEX IF NOT EXISTS idx_washrooms_building_floor ON washrooms(building, floor);

CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    aggregate_id TEXT NOT NULL,
    type TEXT NOT NULL,
    data BLOB NOT NULL,
    version INTEGER NOT NULL,
    timestamp DATETIME NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_events_aggregate_id ON events(aggregate_id);
