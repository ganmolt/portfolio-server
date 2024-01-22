
-- +migrate Up
CREATE TABLE IF NOT EXISTS works (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  url varchar(4095) NOT NULL,
  description varchar(4095) NOT NULL,
	encoded_img text NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
	PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS works;
