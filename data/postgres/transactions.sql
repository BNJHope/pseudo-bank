CREATE TABLE transaction (
   id serial PRIMARY KEY,
   Amount double precision,
   FromAccount VARCHAR ( 255 ),
   ToAccount VARCHAR ( 255 )
);

INSERT into transaction (Amount, FromAccount, ToAccount)
VALUES (20.5, 'Account1', 'Account2')