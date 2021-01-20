CREATE TABLE IF NOT EXISTS public.sheets
(
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(100) NOT NULL,
    has_positions BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS public.metals
(
    sheets_id SERIAL,
    id SERIAL PRIMARY KEY,
    field VARCHAR(100) NOT NULL,
    value VARCHAR(100) NOT NULL,
    CONSTRAINT fk_sheets
       FOREIGN KEY(sheets_id)
          REFERENCES sheets(id)
);

CREATE TABLE IF NOT EXISTS public.crypto
(
    sheets_id SERIAL,
    id SERIAL PRIMARY KEY,
    field VARCHAR(100) NOT NULL,
    value VARCHAR(100) NOT NULL,
    CONSTRAINT fk_sheets
       FOREIGN KEY(sheets_id)
          REFERENCES sheets(id)
);
