CREATE TABLE "warehouses" (

"id" BIGSERIAL not null PRIMARY KEY,
"name" varchar not null,
"slug" varchar not null unique,
"company_id" integer not null,
"address" varchar not null
);
