-- CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
	id uuid DEFAULT uuid_generate_v4(),
	first_name VARCHAR(250),
	last_name VARCHAR(250),
	date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	
	PRIMARY KEY (id)
);

INSERT INTO users(first_name, last_name) VALUES('Maryam', 'Yahya');

