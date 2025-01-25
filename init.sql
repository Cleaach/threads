CREATE DATABASE forum;

USE forum;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE threads (
    id INT AUTO_INCREMENT PRIMARY KEY,
    author_id INT NOT NULL,
    category_id INT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE categories (
	id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE comments (
	id INT AUTO_INCREMENT PRIMARY KEY,
    author_id INT NOT NULL,
    thread_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (thread_id) REFERENCES threads(id)
);