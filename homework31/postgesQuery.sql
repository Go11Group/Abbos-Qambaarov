create table product
(
    id       uuid primary key DEFAULT gen_random_uuid(),
    name     varchar,
    category varchar,
    cost     int
);
DROP table product;
SELECT
    tablename,
    indexname,
    indexdef
FROM
    pg_indexes
WHERE
    schemaname = 'public'
ORDER BY
    tablename,
    indexname;
-- single index
CREATE INDEX idx_name ON product (name);

-- multi index
-- create index product_id_idx_cat on product (name);
create unique index product_id_idx on product (id, cost);
--brin index
CREATE INDEX index_brin ON product USING brin (name);
explain (analyse )
select * from product where name  = 'cristal' 

