CREATE TABLE "kompany" (

"id" integer not null primary key,
"name" varchar not null,
"slug" varchar not null unique,
"inn" integer not null,
"kpp" integer not null

);