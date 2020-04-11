SET timezone = '+00:00';

CREATE TABLE member (
  member_id SERIAL PRIMARY KEY,
  username VARCHAR(40) NOT NULL UNIQUE,
  pwhash VARCHAR(60) NOT NULL,
  is_admin BOOLEAN NOT NULL, 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE news (
  news_id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  url VARCHAR(300) NOT NULL,
  author_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  approved_at TIMESTAMP,
  FOREIGN KEY(author_id) REFERENCES member(member_id) ON DELETE SET NULL
);

CREATE TABLE comment (
  comment_id SERIAL PRIMARY KEY,
  content VARCHAR NOT NULL,
  member_id INT,
  news_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(member_id) REFERENCES member(member_id) ON DELETE SET NULL,
  FOREIGN KEY(news_id) REFERENCES news(news_id) ON DELETE CASCADE
);

CREATE TABLE comment_like (
  member_id INT,
  comment_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(member_id, comment_id),
  FOREIGN KEY(member_id) REFERENCES member(member_id) ON DELETE CASCADE,
  FOREIGN KEY(comment_id) REFERENCES comment(comment_id) ON DELETE CASCADE
);

CREATE TABLE news_like (
  member_id INT,
  news_id INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(member_id, news_id),
  FOREIGN KEY(member_id) REFERENCES member(member_id) ON DELETE CASCADE,
  FOREIGN KEY(news_id) REFERENCES news(news_id) ON DELETE CASCADE
);