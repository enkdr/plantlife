-- schema.sql
create schema if not exists plantlife;
create table
     if not exists plantlife.plant (
        id serial primary key,
        name varchar(255) not null,
        description varchar(255) not null,
        image varchar(255) not null,
        water int not null, -- millilitres per day
        sun int not null, -- hours of sunshine per day
        germination int not null, -- days to germinate
        flowering int not null, -- days to flower
        harvest int not null, -- days to harvest
        seed int not null, -- seeds per plant
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plantlife.users (
        id serial primary key,
        name varchar(255) not null,
        email varchar(255) not null,
        password varchar(255) not null,
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plantlife.plant_users (
        id serial primary key,
        plant_id int not null references plantlife.plant (id),
        users_id int not null references plantlife.users (id),
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plantlife.store (
        id serial primary key,
        name varchar(255) not null,
        address varchar(255) not null,
        city varchar(255) not null,
        postcode varchar(255) not null,
        country varchar(255) not null,
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plantlife.plant_store (
        id serial primary key,
        quantity int not null,
        plant_id int not null references plantlife.plant (id),
        store_id int not null references plantlife.store (id),
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );
