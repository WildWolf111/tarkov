CREATE TABLE "stocs" (

"id" BIGSERIAL not null primary key,
"company_sender_id" integer not null,
"company_recipient_id" integer not null,
"product_id" integer not null,
"quantity" varchar not null,
"warehouse_cell_id" integer not null,
"gtd_id" integer not null
);
