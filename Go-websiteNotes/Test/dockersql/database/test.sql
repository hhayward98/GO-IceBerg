USE sqldock;


CREATE TABLE `users`
 (
	id INT AUTO_INCREMENT,
	username varchar(255) NOT NULL,
	password TEXT NOT NULL,
	email TEXT NOT NULL,
	created_at DATETIME,
	PRIMARY KEY (`id`),
);


