USE test_db;

CREATE TABLE users
(
    id         INTEGER NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    age        INTEGER,
    created    DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE posts
(
    id       INTEGER NOT NULL AUTO_INCREMENT,
    user_id  INTEGER NOT NULL,
    content  TEXT    NOT NULL,
    created  DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO users (first_name, last_name, age) 
VALUES 
    ("rina", "mikami", 43),
    ("jun", "kusano", 34),
    ("hideki", "yamada", 23);