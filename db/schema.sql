CREATE DATABASE technews;

USE technews;

SET time_zone = '+00:00';

CREATE TABLE user (
  user_id INT PRIMARY KEY AUTO_INCREMENT,
  nickname VARCHAR(40) NOT NULL,
  pwhash VARCHAR(40) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE news (
  news_id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  url VARCHAR(300) NOT NULL,
  author_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  approved_at TIMESTAMP,
  FOREIGN KEY(author_id) REFERENCES user(user_id) ON DELETE SET NULL
);

CREATE TABLE comment (
  comment_id INT PRIMARY KEY AUTO_INCREMENT,
  content VARCHAR(500) NOT NULL,
  user_id INT,
  news_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(user_id) REFERENCES user(user_id) ON DELETE SET NULL,
  FOREIGN KEY(news_id) REFERENCES news(news_id) ON DELETE CASCADE
);

CREATE TABLE comment_like (
  user_id INT,
  comment_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, comment_id),
  FOREIGN KEY(user_id) REFERENCES user(user_id) ON DELETE CASCADE,
  FOREIGN KEY(comment_id) REFERENCES comment(comment_id) ON DELETE CASCADE
);

CREATE TABLE news_like (
  user_id INT,
  news_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, news_id),
  FOREIGN KEY(user_id) REFERENCES user(user_id) ON DELETE CASCADE,
  FOREIGN KEY(news_id) REFERENCES news(news_id) ON DELETE CASCADE
);