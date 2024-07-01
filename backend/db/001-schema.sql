-- Create the plant table
create table if not exists plant (
    id integer primary key autoincrement,
    name varchar(255) not null,
    description varchar(255) not null,
    image varchar(255) not null,
    water int not null, -- millilitres per day
    sun int not null, -- hours of sunshine per day
    germination int not null, -- days to germinate
    flowering int not null, -- days to flower
    harvest int not null, -- days to harvest
    seed int not null, -- seeds per plant
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- Create the users table
create table if not exists users (
    id integer primary key autoincrement,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- Create the plant_users table
create table if not exists plant_users (
    id integer primary key autoincrement,
    plant_id int not null references plant(id),
    users_id int not null references users(id),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- Create the store table
create table if not exists store (
    id integer primary key autoincrement,
    name varchar(255) not null,
    address varchar(255) not null,
    city varchar(255) not null,
    postcode varchar(255) not null,
    country varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- Create the plant_store table
create table if not exists plant_store (
    id integer primary key autoincrement,
    quantity int not null,
    plant_id int not null references plant(id),
    store_id int not null references store(id),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
