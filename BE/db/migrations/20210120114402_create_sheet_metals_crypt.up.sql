CREATE TABLE IF NOT EXISTS public.sheets
(
    id SERIAL PRIMARY KEY,
    has_metals BOOLEAN NOT NULL,
    has_crypto BOOLEAN NOT NULL
);
