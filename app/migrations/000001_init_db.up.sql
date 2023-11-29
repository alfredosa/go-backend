-- type User struct {
-- 	ID           int
-- 	Username     string
-- 	PasswordHash string
-- 	Email        string
-- 	FirstName    string
-- 	LastName     string
-- 	IsActive     bool
-- 	IsAdmin      bool
-- 	CreatedAt    time.Time
-- 	UpdatedAt    time.Time
-- }
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid default uuid_generate_v4() primary key,
    username varchar not null unique,
    email varchar not null,
    password_hash varchar not null,
    full_name varchar null,
    first_name varchar null,
    last_name varchar null,
    is_active boolean not null default true,
    is_admin boolean not null default false,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);