CREATE TABLE tasks (
    id TINYINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    completed BOOLEAN,
    PRIMARY KEY (id)
);

INSERT INTO
    tasks (title, completed)
VALUES
    ('Task 1', true),
    ('Task 2', false),
    ('Task 3', true),
    ('Task 4', false),
    ('Task 5', true);