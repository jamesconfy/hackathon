-- CREATE TYPE status AS ENUM ('Rejected', 'Done', 'Pending');

ALTER TABLE deposits ADD COLUMN status status DEFAULT 'Pending';