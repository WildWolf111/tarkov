CREATE TABLE "companies" (

"id" integer not null primary key,
"name" varchar not null,
"slug" varchar not null unique,
"inn" integer not null,
"kpp" integer not null

);


ALTER TABLE "companies" ADD(DROP) COLUMN 
