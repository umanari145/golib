
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE persons (
  id  SERIAL PRIMARY KEY,
  person_name varchar(50) DEFAULT '',
  email varchar(100) DEFAULT '',
  age smallint DEFAULT NULL,  
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO persons (person_name, email, age) VALUES 
('山田太郎', 'tarou@sample.com',20);
INSERT INTO persons (person_name, email, age) VALUES 
('山田花子', 'hanako@sample.com',25);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE persons;
