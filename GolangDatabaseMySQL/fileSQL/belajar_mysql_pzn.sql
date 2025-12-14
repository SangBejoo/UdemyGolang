CREATE TABLE barang (
    id INT not null ,
    nama VARCHAR(100) not null,
    harga INT not null default 0,
    jumlah INT
)engine InnoDB;

describe table barang;

use belajar_mysql;

alter table barang
modify deskripsi text not null ;

alter table barang
modify id int not null ;

describe barang;

show create table barang;


show tables;

describe barang;

show create table barang;

alter table barang
    add column deskripsi text;


show tables;

describe barang;

alter table barang
add column salah text;

describe barang;

alter table barang
drop column salah;

describe barang;


alter table barang
modify harga int not null default 0;

alter table barang
add waktu_dibuat timestamp not null default current_timestamp;


describe barang;

alter table barang
modify deskripsi int default 0;

describe barang;

insert into barang(id, nama) value (1, 'apel');

describe barang;

select * from barang;

truncate barang;

drop table barang;

show tables ;

select * from barang;



create table product (
    id varchar(10) not null,
    name varchar(100) not null,
    description text,
    price int unsigned not null,
    quantity int unsigned not null default 0,
    created_at timestamp not null default current_timestamp
)engine innoDB;

show tables;

insert into product(id, name, description,  price, quantity)
value ('poo1', 'mie ayam', 'mie ayam bakso', 15000, 100);

select * from product;

# batch data
insert into product(id, name, description, price, quantity)
value ('p002', 'mie ayam', 'mie ayam bakso', 20000, 200),
    ('p003', 'mie ayam', 'mie ayam bakso pangsit', 20200, 500),
    ('p004', 'mie ayam', 'mie ayam bakso yamin', 25000, 100);

select * from product;

select id, name, price, quantity from product;

alter table product
add primary key (id);

alter table product
add primary key (id);

show create table product;


select * from product
where price = 15000;

select * from product where quantity = 100;

select * from product where name = 'mie ayam';


update product set description = 'mie ayam pangsit bakso' where description is null;

describe product;

select * from product;

alter table product
add column category enum('makanan', 'minuman', 'lain-lain')
after name;

describe product;

# mengubah satu kolom
update product
set category = 'makanan'
where id = 'poo1';

select * from product where id = 'poo1';

select *
from product;

select * from product where id = 'p003'

update product
set id = 'p001'
where id = 'poo1';

select * from product;

update product
set id = 'p002'
where id = 'poo2';

select * from product;

update product
set category = 'makanan';

select * from product;

# mengubah dengan value di kolom
update product
set price = price + 5000
where id = 'p001';

select * from product;


insert into product (id, name, description, price)
values ('p006', 'mie ayam', 'mie indomie', 1999);


select * from product;

delete
from product
where id = 'p006' ;

#alias kolom

select p.id as kode,
       p.name as nama,
       p.category as kategori,
       p.price as harga,
       p.quantity as stok
from product as p;


insert into product (id, name, description, price)
values ('p007', 'mie ayam', 'mie indomie', 1999),
       ('p008', 'mie ayam', 'mie indomie', 1999),
       ('p009', 'mie ayam', 'mie indomie', 1999),
       ('p010', 'mie ayam', 'mie indomie', 1999),
       ('p011', 'mie ayam', 'mie indomie', 1999),
       ('p012', 'mie ayam', 'mie indomie', 1999),
       ('p013', 'mie ayam', 'mie indomie', 1999),
       ('p014', 'mie ayam', 'mie indomie', 1999),
       ('p015', 'mie ayam', 'mie indomie', 1999);


select *
from product;

select * from product
update product set quantity = 100
where quantity = 0;

select * from product where category = 'makanan';

select * from product where quantity > 100 and price > 20000;

select * from product where quantity < 100 and price >1000;

select * from product where category = 'makanan' and price < 25000;


select * from product where quantity > 100 or price > 1000000;

select * from product where price >= 100000 or quantity <= 500;

select * from product where (price <= 20000 or category = 'makanan') and quantity >= 200;

select * from product where name like '%mie%';
select * from product where name like '%yam%';
select * from product where description like '%it%';


select * from product where description like '%sit';

select * from product where description is not null;

#between

select * from product where price between 10 and 200000;
select * from product where price not between  100 and 20000;

select * from product where category in ('makanan', 'minuman');


select * from product where category not in ('makanan', 'minuman');


select * from product order by price asc , id desc;

select id, category, name  from product order by id;


select category from product where category is null;

#limit clause

select * from product where price > 0
order by price desc
limit 2;

# skip atau offset

select * from product where price > 0
order by id asc
limit 0, 5;

# distinct

select distinct category from product;

select distinct category, id from product;

select *
from product;

select 10, 10, 10 * 10 as hasil;
select id, name, price, price div 1000 as 'price in k' from product;

select pi();

select power(10, 20);

select id, cos(price), sin(price), tan(price) from  product;

select id, name, price
from product
where price div 10000 > 10;


use belajar_mysql;


create table admin(
    id int not null auto_increment,
    first_name varchar(100) not null,
    last_name varchar(100) not null ,
    primary key (id)
)engine InnoDB;

describe admin;

insert into admin(first_name, last_name)
values ('ayub', 'subagiya'),
       ('subagiya', 'ayub');

select * from admin order by id;

delete from admin where id = 1;

insert into admin(first_name, last_name)
values ('ayubasas', 'subagiya');


select * from admin;

select last_insert_id();

select id,
       lower(name) as 'name lower',
       length(name) as ' name lenght',
       upper(name) as 'name upper'
from product;

select id, created_at,
       extract(year from created_at) as year,
       extract(month from created_at) as month
from product;

update product set category = 'minuman' where category is null;

select * from product;





select id,
       case category
            when 'makanan' then 'enak'
            when 'minuman' then 'segar'
            else 'gak ngerti'
            end as 'Category'
from product;


select id,
       price,
       if (price <= 15000, 'murah',
           if (price <= 20000, 'mahal', 'mahal banget')) as 'Mahal?'
from product;

select id, name, ifnull(description, 'kososng') from product;

select count(id) as 'total product' from product;
select max(price) as 'harga termahal' from product;
select min(price) as 'harga termurah' from product;
select avg(price) as 'rata-rata' from product;
select sum(quantity) as 'total stock', quantity from product group by quantity;

select category,
       count(id) as 'total id'
from product
group by category

select
       count(id) as 'total',category
from product
group by category
having total > 1;

create table customer (
    id int not null auto_increment,
    email varchar(100) not null,
    first_name varchar(100) not null ,
    last_name varchar(100),
    primary key (id),
    unique key email_unique(email)
)engine InnoDB;

describe customer;

alter table customer
drop constraint email_unique;


alter table customer
add constraint email_unique unique (email);

insert into customer(email, first_name, last_name)
values ('ayub@gmail.com', 'ayub', 'subagiya');

select * from customer;

insert into customer(email, first_name, last_name)
VALUES ('ay21ub@gmail.com', 'subagiya','ayub')


select * from product;


insert into product(id, name, category, price, quantity)
VALUES ('p016', 'permen', 'lain-lain', 500, 1000)


select * from product;

update product
set price = 1000
where id = 'p0016';

alter table product
add constraint price_check check (price >= 1000);


update product
set price = 1000
where id = 'p0016';

# table index
create table sellers (
    id int not null auto_increment,
    name varchar(100)  ,
    name1 varchar(100)  ,
    email varchar(100) not null ,
    primary key (id),
    unique key email_unique(email),
    index name_name1_index(name, name1)
)engine = InnoDB;

drop table sellers;

desc sellers;

show create table sellers;

select * from sellers where name = 'x';

select * from sellers where name = 'x' and name1 = 'x';

# menambah/ menghapus index
alter table sellers
add index name_index(name);

alter table sellers
drop index name_index;


# full text search, mysql bukan database spesialis search engine
# belajar elastic search dan sol R

# menambah menghapus fulltext search di table yang ada
alter table product
add fulltext product_search(name, description);

show create table product;

select *
from product;


select  * from product where match(name, description)
                                   against ('ayam' in natural language mode);


select * from product where match(name, description)
                                  against ('ayam' in boolean mode );


select * from product where match(name, description)
                                  against ('ayam' with query expansion );


show create table product;

# table relational
create table wishlist(
    id int not null auto_increment,
    id_product varchar(10) not null,
    description text,
    primary key (id),
    constraint fk_wishlist_product
                     foreign key (id_product) references product(id)
)engine = InnoDB;

# mengubah dan menghapus foreign key
alter table wishlist
    drop constraint fk_wishlist_product;

alter table  wishlist
    add constraint fk_wishlist_product
        foreign key (id_product) references product(id);

drop table whislist;


insert into wishlist(id_product, description)
VALUES ('p001', 'makanan kesukaan');

select * from wishlist;

delete from product
where id = 'p001';


alter table wishlist
    drop constraint fk_wishlist_product;

alter table wishlist
add constraint fk_wishlist_product
foreign key (id_product) references product(id)
on delete cascade on update cascade;

show create table product;

insert into product(id, name, category, price, quantity)
VALUES ('pxxxx', 'contoh', 'lain-lain', 1000, 100)

insert into wishlist(id_product, description)
VALUES ('Pxxxx', 'makanan');

select * from wishlist;

delete from product
where id = 'Pxxxx'

select * from wishlist
join product on (wishlist.id_product = product.id);


## melakukan join
select ws.id as id_wishlist, p.id as id_product, p.name, ws.description
from wishlist as ws
join product as p on (ws.id_product = p.id)

desc wishlist;

alter table wishlist
add column id_customer int;

alter table wishlist
add constraint fk_wishlist_customer
foreign key (id_customer) references customer(id);

select * from customer;

update wishlist set id_customer = 1
where id = 1;

select * from wishlist;

select customer.email, product.id, product.name, wishlist.description
from wishlist
join product on (product.id = wishlist.id_product)
join customer on (customer.id = wishlist.id_customer)

# one to one relationship
create table wallet(
    id int not null auto_increment,
    id_customer int not null ,
    balance int not null default 0,
    primary key (id),
    unique key id_customer_unique (id_customer),
    foreign key fk_wallet_customer(id_customer) references customer(id)
)engine InnoDB;

select * from  customer;

insert into wallet(id_customer) VALUES (1), (3)

select * from wallet;


select customer.email, wallet.balance
from wallet join customer on (wallet.id_customer = customer.id)


#table category one to many

create table category (
    id varchar(10) not null,
    name varchar(100) not null ,
    primary key (id)
)engine = InnoDB;


drop table category;

# mengubah table product
alter table product
drop column category;

desc product;

alter table product
add column id_category varchar(10);

alter table product
add constraint fk_product_category
foreign key (id_category) references category(id);


#one to many
show create table product;

select *
from product;

insert into category(id, name) values ('c001', 'makanan'),
                                      ('c002', 'minuman'),
                                      ('c003', 'lain-lain');


update product
set id_category = 'c001'
where id in ('p001', 'p002', 'p003', 'p004', 'p005');

update product
set id_category = 'c002'
where id in ('p006', 'p007', 'p008');

update product
set id_category = 'c003'
where id in ('p009', 'p0010', 'p0011');


select product.id, product.name, category.name
from product
join category on (category.id = product.id_category)


#many to many

create table orders (
    id int not null auto_increment,
    total int    not null ,
    order_date datetime not null default current_timestamp,
    primary key (id)
)engine = InnoDB;

-- Drop the foreign key constraint in orders_detail first
ALTER TABLE orders_detail DROP CONSTRAINT fk_orders_detail_order;

-- Now drop the orders table
DROP TABLE orders;


drop table orders;

describe orders;


create table orders_detail(
    id_product varchar(10) not null,
    id_order int not null ,
    price int not null ,
    quantity int not null ,
    primary key (id_product, id_order)

)engine = InnoDb;

drop table orders_detail;

desc orders_detail;

alter table orders_detail
add constraint fk_orders_detail_product
foreign key (id_product) references product(id);

alter table orders_detail
add constraint fk_orders_detail_order
foreign key (id_order) references orders(id);

show create table orders_detail;

INSERT INTO orders(total) VALUES (50000);  -- id=1
INSERT INTO orders(total) VALUES (60000);  -- id=2
INSERT INTO orders(total) VALUES (70000);  -- id=3


insert into orders_detail(id_product, id_order, price, quantity)
value ('p001', 1, 25000, 1),
    ('p002', 1, 25000, 1);



insert into orders_detail(id_product, id_order, price, quantity)
value ('p003', 2, 25000, 1),
    ('p004', 3, 25000, 1);


insert into orders_detail(id_product, id_order, price, quantity)
value ('p001', 2, 25000, 1),
    ('p003', 3, 25000, 1);

select * from orders_detail;

select * from orders
join orders_detail on (orders_detail.id_order = orders.id)
join product on (product.id = orders_detail.id_product);

select orders.id, product.id, product.name, orders_detail.quantity, orders_detail.price from orders
join orders_detail on (orders_detail.id_order = orders.id)
join product on (product.id = orders_detail.id_product)


#inner join

select * from category
inner join product on (product.id_category = category.id);

insert into category(id, name)
values ('c004', 'oleh oleh');


insert into product(id, name, price, quantity)
values ('x001', 'x satu', 1000, 100),
       ('x002', 'x dua', 1000, 100),
       ('x003', 'x tiga', 1000, 100);

#left join
select * from category
left join product on(product.id_category = category.id);

#right jojin
select * from category
right join product on (product.id_category = category.id);

#cross join
select * from category
cross join product;

create table number(
    id int not null,
    primary key (id)
)engine = InnoDB;

insert into number(id) values (1), (2),(3),(4),(5)

select number1.id, number2.id, (number1.id * number2.id)
from number as number1
    cross join number as number2 order by number1.id, number2.id;


#subquery di where clause
select * from product
where price > (select avg(price) from product);

select max(price) from product;

select price from category
join product on (product.id_category = category.id);

#subquery from

update product
set price = 100000
where id = 'x001';

select max(cp.price)
from (select price from category
    inner join product on (product.id_category = category.id)) as cp;

create table guestbooks(
    id int not null auto_increment,
    email varchar(100),
    title varchar(100),
    content text,
    primary key (id)
)engine = InnoDb;

select * from customer;

insert into guestbooks(email, title, content)
values ('guest@gmail.com', 'Hello', 'World'),
       ('guest1@gmail.com', 'Hello', 'World'),
       ('guest2@gmail.com', 'Hello', 'World'),
       ('guest3@gmail.com', 'Hello', 'World'),
       ('guest4@gmail.com', 'Hello', 'World')

select * from guestbooks;

#set operator

# union
select  email from customer
union
select  email from guestbooks;

#union all
select email, count(emails.email) from
(select  email from customer
union all
select email from guestbooks) as emails
group by emails.email;

#intersect
select distinct email from customer
where email in(select distinct email from guestbooks);

select distinct customer.email from customer
inner join guestbooks on (guestbooks.email = customer.email);

#minus

select distinct customer.email, guestbooks.email from customer
left join guestbooks on (customer.email = guestbooks.email)
where guestbooks.email is null;


#start transaction with commit

start transaction ;

insert into guestbooks(email, title, content)
VALUES ('contoh@gmail.com','contoh', 'contoh'),
       ('contoh1@gmail.com','contoh', 'contoh'),
       ('contoh2@gmail.com','contoh', 'contoh')

select * from guestbooks;

commit;
# sampai commit ini akan di lakukan


#ini untuk roll back
start transaction ;
 o
delete from guestbooks;

rollback ;

#start transaction ini bagus di lakukan


#locking

start transaction;
update guestbooks
set title = 'diubah oleh user 2'
where id = 1;

commit ;

#locking manual
start transaction;

select * from product where id = 'p001' for update;

update product
set quantity = quantity - 10
where id = 'p001';

commit ;

select * from product where id = 'p001';

#dead lock

start transaction ;

select * from product where id = 'p001' for update;

select * from product where id = 'p002' for update;

commit ;

#lock table read
lock tables product read;

update product
set quantity = 100
where id = 'p001';

unlock tables ;


#lock table write
lock tables product write ;

unlock tables ;

update product
set quantity = 100
where id = 'p001';

#lock instance
lock instance for backup;

unlock instance;

# user management
create user 'ayub'@'localhost';
create user 'subagiya'@'localhost';


#menambah hak akses
grant select on belajar_mysql.* to 'ayub'@'localhost';
grant select, insert, update, delete on belajar_mysql.* to 'subagiya'@'localhost';

show grants for 'ayub'@'localhost';
show grants for 'subagiya'@'localhost';

revoke select on belajar_mysql.* from 'ayub'@'localhost';
revoke select, insert, update, delete on belajar_mysql.* from 'subagiya'@'localhost';

#back up database
#mysqldump --user root --password --set-gtid-purged=OFF --single-transaction belajar_mysql --result-file=D:/backup/belajar_mysql.sql

#restore import datab
# cmd /c "mysql --user=root --password belajar_mysql_import < D:/backup/belajar_mysql.sql"

# import database dari data

