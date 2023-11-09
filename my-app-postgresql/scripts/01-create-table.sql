-- Database: m1_myapp_archi_log
-- DROP DATABASE m1_myapp_archi_log;

-- CREATE DATABASE m1_myapp_archi_log;

\c m1_myapp_archi_log

CREATE TABLE IF NOT EXISTS articles(
    id serial PRIMARY KEY NOT NULL,
    title text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    date text COLLATE pg_catalog."default"
);

