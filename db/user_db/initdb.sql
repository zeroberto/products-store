CREATE TABLE user_info (
	id bigserial NOT NULL,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	date_of_birth date NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	deactivated_at timestamp NULL,
	CONSTRAINT user_info_pk PRIMARY KEY (id)
);

INSERT INTO user_info (first_name,last_name,date_of_birth) VALUES 
  ('Luke','Skywalker','1951-09-25')
  ,('Han','Solo','1942-07-13')
  ,('Darth','Vader','1981-04-19')
  ,('Princess Leia','Organa','1956-10-21')
;
