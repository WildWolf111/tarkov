CREATE TABLE gtds(

id BIGSERIAL PRIMARY KEY,
country_id BIGINT,
number integer,


CONSTRAINT fk_country_id FOREIGN KEY (country_id) REFERENCES countries(id)


);