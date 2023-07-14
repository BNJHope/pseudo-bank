DROP TABLE IF EXISTS users;

CREATE TABLE users (
   id uuid PRIMARY KEY,
   firstname VARCHAR ( 255 ),
   surname VARCHAR ( 255 )
);

INSERT into users (id, firstname, surname)
VALUES ('d2e19190-59c8-4a43-8bb7-a729ea2b5173', 'Ben', 'Hope');

INSERT into users (id, firstname, surname)
VALUES ('1a8580b6-fb6c-4f3a-8254-3c19e638f385', 'Second', 'User');
