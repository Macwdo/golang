DROP TABLE IF EXISTS languages;
CREATE TABLE languages (
    id          INT AUTO_INCREMENT NOT NULL,
    name        VARCHAR(255) NOT NULL,
    rating      DECIMAL(5, 2) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
);


INSERT INTO languages
    (name, rating)
VALUES
    ('PHP', 1.2),
    ('JavaScript', 0),
    ('Python', 5.0),
    ('Java', 3.5),
    ('C#', 3.0),
    ('C++', 4),
    ('Ruby', 4.5),
    ('Swift', 1.5),
    ('Objective-C', 1.0),
    ('Go', 5.0);
