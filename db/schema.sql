CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    start_date DATE,
    plan_date DATE,
    detail VARCHAR(50),
    status BOOLEAN
);
