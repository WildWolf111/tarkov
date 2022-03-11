CREATE TABLE warehouses_cells (

id BIGSERIAL PRIMARY KEY,
name varchar,
slug varchar,
warehouses_id BIGINT,

CONSTRAINT fk_warehouses_id FOREIGN KEY (warehouses_id) REFERENCES warehouses(id)


);