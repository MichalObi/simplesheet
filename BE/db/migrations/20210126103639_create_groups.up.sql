CREATE TABLE IF NOT EXISTS public.groups
(
    sheet_id INT,
    group_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    positions JSONB NOT NULL,
    CONSTRAINT fk_sheets
       FOREIGN KEY(sheet_id)
          REFERENCES sheets(id)
);
