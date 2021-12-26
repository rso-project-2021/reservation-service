CREATE TABLE "reservations" (
    "reservation_id"    BIGSERIAL PRIMARY KEY,
    "station_id"        INT NOT NULL,
    "user_id"           INT NOT NULL,
    "start"             TIMESTAMP NOT NULL,
    "end"               TIMESTAMP NOT NULL,
    "created_at"        TIMESTAMP NOT NULL DEFAULT(now())
);