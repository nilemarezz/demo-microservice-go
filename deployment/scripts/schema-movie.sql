-- CREATE USER 'test'@'localhost' IDENTIFIED BY 'password';

-- GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'test'@'localhost' WITH GRANT OPTION;

-- GRANT ALL PRIVILEGES ON *.* TO 'test'@'localhost' WITH GRANT OPTION;
-- flush privileges; 

CREATE TABLE celebrities (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  age int NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE movies (
  movie_id int NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  description varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  screen_date date NOT NULL,
  PRIMARY KEY (movie_id)
);

CREATE TABLE movie_celebritry (
  movie_id int NOT NULL,
  celebritry_id int NOT NULL,
  KEY movie_id (movie_id),
  KEY celebritry_id (celebritry_id),
  CONSTRAINT movie_celebritry_ibfk_1 FOREIGN KEY (movie_id) REFERENCES movies (movie_id),
  CONSTRAINT movie_celebritry_ibfk_2 FOREIGN KEY (celebritry_id) REFERENCES celebrities (id)
);


INSERT INTO movies (name,description,screen_date) VALUES
	 ('The Shawshank Redemption','Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.','1994-10-22'),
	 ('The Godfather','The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.','1972-05-10'),
	 ('The Dark Knight','When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.','2008-01-01');

INSERT INTO celebrities (name,age) VALUES
	 ('Christian Bale',52),
	 ('Heath Ledger',49),
	 ('Marlon Brando',71),
	 ('Al Pacino',71),
	 ('Robert Duvall',30),
	 ('Tim Robbins',65);

INSERT INTO movie_celebritry (movie_id,celebritry_id) VALUES
	 (3,1),
	 (3,2),
	 (2,3),
	 (2,4),
	 (2,5),
	 (1,6);