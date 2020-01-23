CREATE TABLE author(
    id serial primary key,
    first_name text,
    last_name text,
    UNIQUE (first_name, last_name)
);

CREATE TABLE book(
    id serial PRIMARY KEY,
    author_id integer REFERENCES author(id) ON UPDATE CASCADE ON DELETE CASCADE,
    title varchar(255),
    number_pages integer,
    UNIQUE (author_id, title)
);

CREATE TABLE genre (
  id serial PRIMARY KEY,
  name varchar(64) unique
);

CREATE TABLE book_genre (
    book_id INTEGER REFERENCES book(id) ON UPDATE CASCADE ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genre(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT book_genre_pk PRIMARY KEY (book_id, genre_id)
);

INSERT INTO genre (name) VALUES
('Fantasy'),
('ScienceFiction'),
('Horror'),
('Western'),
('Romance'),
('Thriller'),
('Mystery'),
('Detective'),
('Dystopia'),
('Newspaper');
