CREATE DATABASE "golang-test";

CREATE TABLE IF NOT EXISTS public.user_data
(
    id SERIAL NOT NULL,
    user_id character varying(80) COLLATE pg_catalog."default" NOT NULL,
    first_name character varying(80) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(80) COLLATE pg_catalog."default" NOT NULL,
    sex character varying(10) COLLATE pg_catalog."default" NOT NULL,
    email character varying(80) COLLATE pg_catalog."default" NOT NULL,
    phone character varying(80) COLLATE pg_catalog."default" NOT NULL,
    dob date NOT NULL,
    job_title character varying(80) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_data_pkey PRIMARY KEY (id),
    CONSTRAINT user_data_email_key UNIQUE (email),
    CONSTRAINT user_data_phone_key UNIQUE (phone),
    CONSTRAINT user_data_user_id_key UNIQUE (user_id)
);

ALTER TABLE IF EXISTS public.user_data
    OWNER to postgres;