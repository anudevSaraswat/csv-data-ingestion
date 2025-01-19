CREATE DATABASE "golang-test"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'C'
    LC_CTYPE = 'C'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public.user_data
(
    id integer NOT NULL,
    user_id character varying(80) COLLATE pg_catalog."default" NOT NULL,
    name character varying(80) COLLATE pg_catalog."default" NOT NULL,
    email character varying(80) COLLATE pg_catalog."default" NOT NULL,
    dob date NOT NULL,
    city character varying(80) COLLATE pg_catalog."default",
    CONSTRAINT user_data_pkey PRIMARY KEY (id),
    CONSTRAINT user_data_email_key UNIQUE (email),
    CONSTRAINT user_data_name_key UNIQUE (name),
    CONSTRAINT user_data_user_id_key UNIQUE (user_id)
)

ALTER TABLE IF EXISTS public."user-data"
    OWNER to postgres;