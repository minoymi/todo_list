CREATE DATABASE todo;

CREATE TABLE todolists (
    id SERIAL PRIMARY KEY,
    userID bigint NOT NULL,
    timeDate timestamptz,
    tasks TEXT,
    misc TEXT
);

INSERT INTO todolists (userID, timeDate, tasks, misc)
VALUES (1, '2023-01-02', 'uhh buy grocery', 'not done'),
(2, '2023-01-03', 'walk dog', 'done'),
(2, '2000-09-09', 'buy bitcoin', 'whats a bitcoin haha xd');