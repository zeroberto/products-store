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

INSERT INTO user_info (id, first_name,last_name,date_of_birth,created_at) VALUES 
  (1, 'User','Test','1980-01-01','2020-09-01 00:00:00')
;
