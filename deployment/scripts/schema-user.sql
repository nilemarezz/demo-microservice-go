-- CREATE USER 'test'@'localhost' IDENTIFIED BY 'password';

-- GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'test'@'localhost' WITH GRANT OPTION;

-- GRANT ALL PRIVILEGES ON *.* TO 'test'@'localhost' WITH GRANT OPTION;
-- flush privileges; 

CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  username varchar(100) NOT NULL,
  password varchar(1000) NOT NULL,
  PRIMARY KEY (id)
);
