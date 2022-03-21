CREATE TABLE "products" (

	Id       BIGSERIAL not null PRIMARY KEY,
	Name      varchar not null,
	Slug     varchar not null,
	Sku       BIGINT not null,
	ShortDesc varchar not null,
	FullDesc  varchar not null,
	Sort      BIGINT   

CONSTRAINT fk_warehouses_id FOREIGN KEY (warehouse_cell_id) REFERENCES warehouses_cells(id),

CONSTRAINT fk_recepient FOREIGN KEY (company_recipient_id) REFERENCES companies(id),
CONSTRAINT fk_sender FOREIGN KEY (company_sender_id) REFERENCES companies(id)

);