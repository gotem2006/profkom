CREATE TYPE role AS ENUM ('client', 'admin', 'worker');

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    role role NOT NULL DEFAULT 'client',
    name TEXT,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS invite_token (
    id SERIAL PRIMARY KEY,
    content UUID NOT NULL,
    used BOOLEAN DEFAULT false,
    role role NOT NULL
);

CREATE TABLE IF NOT EXISTS news (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY,
    intro TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    type  text,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TYPE guide_type AS ENUM ('worker', 'student');

CREATE TABLE IF NOT EXISTS guides (
    id SERIAL PRIMARY KEY,
    type guide_type NOT NULL,
    title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS themes (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    guide_id INTEGER REFERENCES guides(id) ON DELETE CASCADE
);

create table if not exists documents(
    id UUID primary key,
    url text,
    title text,
    type text
)