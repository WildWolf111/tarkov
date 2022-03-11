CREATE TABLE "stocs" (

"id" BIGSERIAL not null PRIMARY KEY,
"company_sender_id" integer not null,
"company_recipient_id" integer not null,
"product_id" integer not null,
"quantity" varchar not null,
"warehouse_cell_id" integer not null,
"gtd_id" integer not null,

CONSTRAINT fk_warehouses_id FOREIGN KEY (warehouse_cell_id) REFERENCES warehouses_cells(id),

CONSTRAINT fk_recepient FOREIGN KEY (company_recipient_id) REFERENCES companies(id),
CONSTRAINT fk_sender FOREIGN KEY (company_sender_id) REFERENCES companies(id)

);
