INSERT INTO user(nickname, pwhash) VALUES("Jack", "asdasd");
INSERT INTO user(nickname, pwhash) VALUES("John", "asdasd");
INSERT INTO user(nickname, pwhash) VALUES("Mary", "asdasd");
INSERT INTO user(nickname, pwhash) VALUES("Nick", "asdasd");
INSERT INTO user(nickname, pwhash) VALUES("Richard", "asdasd");

INSERT INTO news(title, url, author_id) VALUES("News item 1", "https://www.google.fi", 1);
INSERT INTO news(title, url, author_id) VALUES("News item 2", "https://www.google.fi", 2);
INSERT INTO news(title, url, author_id) VALUES("News item 3", "https://www.google.fi", 3);

INSERT INTO comment(content, user_id, news_id) VALUES("Cool! t. Jack", 1, 3);
INSERT INTO comment(content, user_id, news_id) VALUES("Cool! t. John", 2, 3);
INSERT INTO comment(content, user_id, news_id) VALUES("Cool! t. Mary", 3, 1);
INSERT INTO comment(content, user_id, news_id) VALUES("I prefer JS t. Jack", 1, 3);
INSERT INTO comment(content, user_id, news_id) VALUES("Or maybe C++ t. Jack", 1, 3);

INSERT INTO comment_like VALUES(3, 1);
INSERT INTO comment_like VALUES(3, 2);
INSERT INTO comment_like VALUES(3, 3);
INSERT INTO comment_like VALUES(3, 4);
INSERT INTO comment_like VALUES(3, 5);

INSERT INTO news_like VALUES(1, 1);
INSERT INTO news_like VALUES(3, 2);
INSERT INTO news_like VALUES(3, 3);
INSERT INTO news_like VALUES(5, 1);
INSERT INTO news_like VALUES(4, 1);