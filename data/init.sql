CREATE TABLE task (
    id VARCHAR(255),
    title VARCHAR(255),
    completed BOOLEAN
);

INSERT INTO
    task (id, title, completed)
VALUES
    ('1', 'Task 1', true),
    ('2', 'Task 2', false),
    ('3', 'Task 3', true),
    ('4', 'Task 4', false),
    ('5', 'Task 5', true);