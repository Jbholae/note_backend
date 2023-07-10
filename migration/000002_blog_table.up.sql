CREATE TABLE IF NOT EXISTS blogs(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(45) NOT NULL,
    content VARCHAR(100) NOT NULL,
    author VARCHAR(45) NOT NULL,
    created_at DATETIME NOT NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY(id)
  ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
)