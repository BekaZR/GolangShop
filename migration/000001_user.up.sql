CREATE TABLE public.users (
    id SERIAL NOT NULL,
    user_name VARCHAR(30) NOT NULL,
    password VARCHAR(600) NOT NULL,
    balance NUMERIC(20, 4) NOT NULL DEFAULT '0.0',
    currency VARCHAR(4) NOT NULL,
    PRIMARY KEY (id)
);