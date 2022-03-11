CREATE TABLE companies
 (
id BIGSERIAL not null PRIMARY KEY,
name varchar not null,
slug varchar not null unique,
inn integer not null,
kpp integer not null
);