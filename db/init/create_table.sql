DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS sample_table;

CREATE TABLE sample_table
(
  id    INT(10) NOT NULL AUTO_INCREMENT,
  name  text,
  PRIMARY   KEY (`id`)
);

INSERT INTO sample_table (name) VALUES ("sample1");
INSERT INTO sample_table (name) VALUES ("sample2");
