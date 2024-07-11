-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the plant table
CREATE TABLE IF NOT EXISTS plant (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    image VARCHAR NOT NULL,
    water INT NOT NULL, -- millilitres per day
    sun INT NOT NULL, -- hours of sunshine per day
    germination INT NOT NULL, -- days to germinate
    flowering INT NOT NULL, -- days to flower
    harvest INT NOT NULL, -- days to harvest
    seed INT NOT NULL, -- seeds per plant
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the profiles table
CREATE TABLE IF NOT EXISTS profiles (
    id UUID NOT NULL DEFAULT uuid_generate_v4() REFERENCES auth.users(id) ON DELETE CASCADE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- Create the plant_profiles table
CREATE TABLE IF NOT EXISTS plant_profiles (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    plant_id UUID NOT NULL REFERENCES plant(id),
    profiles_id UUID NOT NULL REFERENCES profiles(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the store table
CREATE TABLE IF NOT EXISTS store (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    postcode VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the plant_store table
CREATE TABLE IF NOT EXISTS plant_store (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    quantity INT NOT NULL,
    plant_id UUID NOT NULL REFERENCES plant(id),
    store_id UUID NOT NULL REFERENCES store(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
    
