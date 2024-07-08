-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE database.public.estate (
 id uuid PRIMARY KEY default  gen_random_uuid()  ,
 width int not null default 0 ,
 length int not null default 0,
 created_at timestamp default current_timestamp
);

create table database.public.estate_tree(
	id uuid primary key default gen_random_uuid(), 
	x int not null, 
	y int not null, 
	height int not null default 1 check (height >= 1 and height <= 30), 
	estate_id uuid references database.public.estate(id), 
	unique(x,y, estate_id)
)
