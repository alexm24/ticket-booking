CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE routes
(
    id    UUID         NOT NULL PRIMARY KEY,
    route VARCHAR(255) NOT NULL
);

CREATE TABLE passengers
(
    id       UUID         NOT NULL PRIMARY KEY,
    route_id UUID REFERENCES routes (id) ON DELETE CASCADE,
    fullname VARCHAR(255) NOT NULL,
    age      serial       NOT NULL,
    phone    VARCHAR(50)  NOT NULL,
    email    VARCHAR(100) NOT NULL
);
