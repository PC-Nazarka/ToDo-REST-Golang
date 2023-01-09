CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    is_done BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE posts(
      id SERIAL PRIMARY KEY,
      text VARCHAR(255),
      user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
      task_id INT REFERENCES tasks(id) ON DELETE CASCADE NOT NULL,
      created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE comments(
      id SERIAL PRIMARY KEY,
      text VARCHAR(255),
      user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
      post_id INT REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
      created_at TIMESTAMP NOT NULL DEFAULT now()
);