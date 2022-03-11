BEGIN;

CREATE TABLE companies_warehouses (

company_id BIGINT,
warehouses_id BIGINT,

CONSTRAINT fk_sender FOREIGN KEY (company_id) REFERENCES companies(id),
CONSTRAINT fk_recepient FOREIGN KEY (warehouses_id) REFERENCES warehouses(id)

);
COMMIT;