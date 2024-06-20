begin;

UPDATE users
SET name = 'Suxrob', email = 'sxrob@gmail.com'
WHERE id = '0804549a-e92a-4bf9-bef5-b22bdb38ca23';

DELETE FROM users
WHERE id = '0804549a-e92a-4bf9-bef5-b22bdb38ca23';

commit;