CREATE TABLE users (
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE todo_lists (
    id          serial       not null UNIQUE,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);
CREATE TABLE users_lists (
    id      serial                                           not null UNIQUE,
    user_id int REFERENCES users (id) ON DELETE CASCADE      NOT NULL,
    list_id int REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);
CREATE table todo_items (
    id          serial          NOT NULL UNIQUE,
    title       VARCHAR(255)    NOT NULL,
    description VARCHAR(255),
    done        BOOLEAN         NOT NULL DEFAULT false
);

CREATE TABLE lists_items (
    id      serial                                           not null UNIQUE,
    item_id int REFERENCES todo_items (id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);