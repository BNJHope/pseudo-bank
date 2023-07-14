DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
   id uuid PRIMARY KEY,
   firstname VARCHAR ( 255 ),
   surname VARCHAR ( 255 )
);

CREATE TABLE transactions (
   id serial PRIMARY KEY,
   Amount double precision,
   FromAccount uuid,
   ToAccount uuid,
   CONSTRAINT fk_fromaccount
      FOREIGN KEY(FromAccount) 
         REFERENCES users(id),
   CONSTRAINT fk_toaccount
      FOREIGN KEY(ToAccount) 
         REFERENCES users(id)
);

INSERT into users (id, firstname, surname)
VALUES ('d2e19190-59c8-4a43-8bb7-a729ea2b5173', 'Ben', 'Hope');

INSERT into users (id, firstname, surname)
VALUES ('1a8580b6-fb6c-4f3a-8254-3c19e638f385', 'Second', 'User');

INSERT into transactions (Amount, FromAccount, ToAccount)
VALUES (20.5, 'd2e19190-59c8-4a43-8bb7-a729ea2b5173', '1a8580b6-fb6c-4f3a-8254-3c19e638f385');
