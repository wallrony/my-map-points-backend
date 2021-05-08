CREATE DATABASE my_map_points;

\use my_map_points;

CREATE TABLE coordinate (
  -- attributes
  id serial primary key,
  latitude varchar(20),
  longitude varchar(20)
);

CREATE TABLE users (
  -- attributes
  id serial primary key,
  name varchar(70) not null,
  birthdate date not null,
  email varchar(60) not null,
  pswd varchar(512) not null,
  salt varchar(512) not null
);

CREATE TABLE map_point (
  -- attributes
  id serial primary key,
  name varchar(100) not null,
  description text,

  -- foreign keys
  user_id int,
  coordinate_id int,

  -- constraints
  constraint map_point_user foreign key (user_id) references users (id),
  constraint map_point_coordinate foreign key (coordinate_id) references coordinate (id)
);
