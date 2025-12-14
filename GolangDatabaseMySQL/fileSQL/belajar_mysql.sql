-- ===============================================
-- MYSQL COMPLETE CHEAT SHEET
-- ===============================================

-- ===============================================
-- 1. DATABASE MANAGEMENT
-- ===============================================

-- Membuat database baru
CREATE DATABASE nama_database;

-- Menggunakan database tertentu
USE nama_database;

-- Melihat semua database
SHOW DATABASES;

-- Menghapus database
DROP DATABASE nama_database;


-- ===============================================
-- 2. TABLE MANAGEMENT - CREATE & STRUCTURE
-- ===============================================

-- Membuat tabel dasar
CREATE TABLE nama_tabel (
                            id INT NOT NULL AUTO_INCREMENT,
                            nama VARCHAR(100) NOT NULL,
                            harga INT NOT NULL DEFAULT 0,
                            jumlah INT,
                            PRIMARY KEY (id)
) ENGINE = InnoDB;

-- Membuat tabel dengan berbagai constraint
CREATE TABLE product (
                         id VARCHAR(10) NOT NULL,
                         name VARCHAR(100) NOT NULL,
                         description TEXT,
                         price INT UNSIGNED NOT NULL,
                         quantity INT UNSIGNED NOT NULL DEFAULT 0,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (id)
) ENGINE = InnoDB;

-- Melihat struktur tabel
DESCRIBE nama_tabel;
DESC nama_tabel;

-- Melihat detail pembuatan tabel (termasuk constraint)
SHOW CREATE TABLE nama_tabel;

-- Melihat semua tabel dalam database
SHOW TABLES;

-- Menghapus tabel
DROP TABLE nama_tabel;


-- ===============================================
-- 3. TABLE ALTERATION - MODIFY STRUCTURE
-- ===============================================

-- Menambah kolom baru
ALTER TABLE product ADD COLUMN deskripsi TEXT;

-- Menambah kolom dengan posisi tertentu
ALTER TABLE product
    ADD COLUMN category ENUM('makanan', 'minuman', 'lain-lain')
AFTER name;

-- Mengubah tipe data kolom
ALTER TABLE product MODIFY deskripsi INT DEFAULT 0;

-- Mengubah kolom menjadi NOT NULL
ALTER TABLE product MODIFY id INT NOT NULL;

-- Menghapus kolom
ALTER TABLE product DROP COLUMN salah;

-- Menambahkan primary key
ALTER TABLE product ADD PRIMARY KEY (id);

-- Menambahkan constraint check
ALTER TABLE product ADD CONSTRAINT price_check CHECK (price >= 1000);


-- ===============================================
-- 4. INSERT DATA - Menambah Data
-- ===============================================

-- Insert single row
INSERT INTO product(id, name, description, price, quantity)
VALUES ('p001', 'Mie Ayam', 'Mie ayam bakso', 15000, 100);

-- Insert multiple rows (batch insert)
INSERT INTO product(id, name, description, price, quantity)
VALUES
    ('p002', 'Mie Ayam Original', 'Mie ayam bakso', 20000, 200),
    ('p003', 'Mie Ayam Pangsit', 'Mie ayam bakso pangsit', 20200, 500),
    ('p004', 'Mie Ayam Yamin', 'Mie ayam bakso yamin', 25000, 100);

-- Insert dengan auto increment
INSERT INTO admin(first_name, last_name)
VALUES ('Ayub', 'Subagiya');

-- Mendapatkan ID terakhir yang di-insert
SELECT LAST_INSERT_ID();


-- ===============================================
-- 5. SELECT - Menampilkan Data
-- ===============================================

-- Select semua kolom
SELECT * FROM product;

-- Select kolom tertentu
SELECT id, name, price, quantity FROM product;

-- Select dengan alias kolom
SELECT
    id AS kode,
    name AS nama,
    category AS kategori,
    price AS harga,
    quantity AS stok
FROM product;

-- Select dengan alias tabel
SELECT p.id, p.name, p.price
FROM product AS p;


-- ===============================================
-- 6. WHERE CLAUSE - Filtering Data
-- ===============================================

-- Kondisi sederhana
SELECT * FROM product WHERE price = 15000;
SELECT * FROM product WHERE quantity = 100;
SELECT * FROM product WHERE name = 'Mie Ayam';

-- Menggunakan operator perbandingan
SELECT * FROM product WHERE price > 20000;
SELECT * FROM product WHERE price >= 15000;
SELECT * FROM product WHERE price < 25000;
SELECT * FROM product WHERE price <= 20000;
SELECT * FROM product WHERE price != 15000;

-- Kondisi AND (semua kondisi harus true)
SELECT * FROM product
WHERE quantity > 100 AND price > 20000;

SELECT * FROM product
WHERE category = 'makanan' AND price < 25000;

-- Kondisi OR (salah satu kondisi true)
SELECT * FROM product
WHERE quantity > 100 OR price > 1000000;

SELECT * FROM product
WHERE price >= 100000 OR quantity <= 500;

-- Kombinasi AND dan OR dengan parentheses
SELECT * FROM product
WHERE (price <= 20000 OR category = 'makanan')
  AND quantity >= 200;


-- ===============================================
-- 7. LIKE - Pattern Matching
-- ===============================================

-- Mencari yang mengandung kata (% = wildcard untuk karakter apapun)
SELECT * FROM product WHERE name LIKE '%mie%';
SELECT * FROM product WHERE name LIKE '%ayam%';
SELECT * FROM product WHERE description LIKE '%it%';

-- Mencari yang diakhiri dengan kata tertentu
SELECT * FROM product WHERE description LIKE '%sit';

-- Mencari yang dimulai dengan kata tertentu
SELECT * FROM product WHERE name LIKE 'Mie%';


-- ===============================================
-- 8. NULL HANDLING
-- ===============================================

-- Mencari data yang NULL
SELECT * FROM product WHERE description IS NULL;

-- Mencari data yang NOT NULL
SELECT * FROM product WHERE description IS NOT NULL;

-- Mengganti NULL dengan nilai default
SELECT id, name, IFNULL(description, 'kosong') FROM product;


-- ===============================================
-- 9. BETWEEN - Range Values
-- ===============================================

-- Mencari dalam rentang nilai (inclusive)
SELECT * FROM product WHERE price BETWEEN 10 AND 200000;

-- Mencari di luar rentang nilai
SELECT * FROM product WHERE price NOT BETWEEN 100 AND 20000;


-- ===============================================
-- 10. IN - Multiple Values
-- ===============================================

-- Mencari yang cocok dengan salah satu nilai
SELECT * FROM product WHERE category IN ('makanan', 'minuman');

-- Mencari yang tidak cocok dengan nilai-nilai tertentu
SELECT * FROM product WHERE category NOT IN ('makanan', 'minuman');


-- ===============================================
-- 11. ORDER BY - Sorting Data
-- ===============================================

-- Sort ascending (kecil ke besar)
SELECT * FROM product ORDER BY price ASC;

-- Sort descending (besar ke kecil)
SELECT * FROM product ORDER BY price DESC;

-- Sort multiple columns
SELECT * FROM product ORDER BY price ASC, id DESC;

-- Sort dengan kolom tertentu
SELECT id, category, name FROM product ORDER BY id;


-- ===============================================
-- 12. LIMIT - Membatasi Hasil
-- ===============================================

-- Mengambil N data pertama
SELECT * FROM product WHERE price > 0
ORDER BY price DESC
    LIMIT 2;

-- Menggunakan OFFSET (skip data)
-- Format: LIMIT offset, jumlah_data
SELECT * FROM product WHERE price > 0
ORDER BY id ASC
    LIMIT 0, 5;  -- Ambil 5 data mulai dari index 0

-- Pagination contoh
LIMIT 5, 5;   -- Ambil 5 data mulai dari index 5 (halaman 2)
LIMIT 10, 5;  -- Ambil 5 data mulai dari index 10 (halaman 3)


-- ===============================================
-- 13. DISTINCT - Menghilangkan Duplikat
-- ===============================================

-- Mendapatkan nilai unik dari satu kolom
SELECT DISTINCT category FROM product;

-- Distinct dari kombinasi kolom
SELECT DISTINCT category, id FROM product;


-- ===============================================
-- 14. UPDATE - Mengubah Data
-- ===============================================

-- Update single column
UPDATE product
SET category = 'makanan'
WHERE id = 'p001';

-- Update multiple columns
UPDATE product
SET category = 'makanan', price = 25000
WHERE id = 'p001';

-- Update semua row (tanpa WHERE)
UPDATE product SET category = 'makanan';

-- Update dengan kalkulasi dari value existing
UPDATE product
SET price = price + 5000
WHERE id = 'p001';

-- Update berdasarkan kondisi
UPDATE product
SET quantity = 100
WHERE quantity = 0;


-- ===============================================
-- 15. DELETE - Menghapus Data
-- ===============================================

-- Delete dengan kondisi
DELETE FROM product WHERE id = 'p006';

-- Delete semua data (structure table tetap ada)
DELETE FROM product;

-- Truncate - hapus semua data dan reset auto increment
TRUNCATE TABLE product;


-- ===============================================
-- 16. NUMERIC FUNCTIONS - Fungsi Matematika
-- ===============================================

-- Operasi matematika dasar
SELECT 10, 10, 10 * 10 AS hasil;
SELECT id, price, price DIV 1000 AS 'price_in_k' FROM product;

-- Fungsi matematika
SELECT PI();                    -- Nilai pi
SELECT POWER(10, 2);           -- Pangkat: 10^2 = 100
SELECT SQRT(16);               -- Akar kuadrat
SELECT ABS(-10);               -- Nilai absolut
SELECT CEIL(4.3);              -- Pembulatan ke atas
SELECT FLOOR(4.7);             -- Pembulatan ke bawah
SELECT ROUND(4.567, 2);        -- Pembulatan 2 desimal

-- Trigonometri
SELECT COS(price), SIN(price), TAN(price) FROM product;


-- ===============================================
-- 17. STRING FUNCTIONS - Fungsi String
-- ===============================================

-- Mengubah case
SELECT
    id,
    LOWER(name) AS 'name_lower',
    UPPER(name) AS 'name_upper'
FROM product;

-- Panjang string
SELECT id, LENGTH(name) AS 'name_length' FROM product;

-- Substring
SELECT SUBSTRING(name, 1, 5) FROM product;

-- Concat
SELECT CONCAT(first_name, ' ', last_name) AS full_name FROM customer;

-- Trim whitespace
SELECT TRIM('  spasi  ');


-- ===============================================
-- 18. DATE & TIME FUNCTIONS
-- ===============================================

-- Mendapatkan tanggal/waktu sekarang
SELECT NOW();                   -- Datetime sekarang
SELECT CURDATE();              -- Date sekarang
SELECT CURTIME();              -- Time sekarang

-- Extract bagian dari date
SELECT
    id,
    created_at,
    EXTRACT(YEAR FROM created_at) AS year,
    EXTRACT(MONTH FROM created_at) AS month,
    EXTRACT(DAY FROM created_at) AS day
FROM product;

-- Format date
SELECT DATE_FORMAT(created_at, '%Y-%m-%d') FROM product;


-- ===============================================
-- 19. CONTROL FLOW - CASE & IF
-- ===============================================

-- CASE statement (seperti switch-case)
SELECT
    id,
    CASE category
        WHEN 'makanan' THEN 'Enak'
        WHEN 'minuman' THEN 'Segar'
        ELSE 'Tidak tahu'
        END AS 'kategori_desc'
FROM product;

-- IF function (kondisi sederhana)
SELECT
    id,
    price,
    IF(price <= 15000, 'Murah',
       IF(price <= 20000, 'Sedang', 'Mahal')
    ) AS 'price_category'
FROM product;


-- ===============================================
-- 20. AGGREGATE FUNCTIONS - Fungsi Agregasi
-- ===============================================

-- COUNT - menghitung jumlah row
SELECT COUNT(id) AS 'total_product' FROM product;
SELECT COUNT(*) AS 'total_rows' FROM product;

-- MAX - nilai maksimum
SELECT MAX(price) AS 'harga_termahal' FROM product;

-- MIN - nilai minimum
SELECT MIN(price) AS 'harga_termurah' FROM product;

-- AVG - rata-rata
SELECT AVG(price) AS 'rata_rata_harga' FROM product;

-- SUM - total/jumlah
SELECT SUM(quantity) AS 'total_stock' FROM product;


-- ===============================================
-- 21. GROUP BY - Mengelompokkan Data
-- ===============================================

-- Group by single column
SELECT
    category,
    COUNT(id) AS 'total_product'
FROM product
GROUP BY category;

-- Group by dengan multiple aggregate
SELECT
    category,
    COUNT(id) AS 'total',
    MAX(price) AS 'termahal',
    MIN(price) AS 'termurah',
    AVG(price) AS 'rata_rata'
FROM product
GROUP BY category;


-- ===============================================
-- 22. HAVING - Filter Setelah Group By
-- ===============================================

-- HAVING digunakan untuk filter hasil aggregate
-- (WHERE tidak bisa digunakan dengan aggregate functions)
SELECT
    category,
    COUNT(id) AS 'total'
FROM product
GROUP BY category
HAVING total > 5;

SELECT
    category,
    AVG(price) AS 'rata_rata'
FROM product
GROUP BY category
HAVING rata_rata > 15000;


-- ===============================================
-- 23. CONSTRAINTS - Batasan Data
-- ===============================================

-- PRIMARY KEY - Unique identifier untuk setiap row
CREATE TABLE example (
                         id INT NOT NULL AUTO_INCREMENT,
                         PRIMARY KEY (id)
);

-- UNIQUE - Nilai harus unik
CREATE TABLE customer (
                          id INT NOT NULL AUTO_INCREMENT,
                          email VARCHAR(100) NOT NULL,
                          PRIMARY KEY (id),
                          UNIQUE KEY email_unique(email)
);

-- Menambah/menghapus UNIQUE constraint
ALTER TABLE customer ADD CONSTRAINT email_unique UNIQUE (email);
ALTER TABLE customer DROP CONSTRAINT email_unique;

-- CHECK constraint - Validasi nilai
ALTER TABLE product
    ADD CONSTRAINT price_check CHECK (price >= 1000);

-- NOT NULL - Kolom wajib diisi
ALTER TABLE product MODIFY name VARCHAR(100) NOT NULL;

-- DEFAULT - Nilai default
ALTER TABLE product
    MODIFY quantity INT NOT NULL DEFAULT 0;


-- ===============================================
-- 24. INDEX - Mempercepat Pencarian
-- ===============================================

-- Membuat index saat create table
CREATE TABLE sellers (
                         id INT NOT NULL AUTO_INCREMENT,
                         name VARCHAR(100),
                         email VARCHAR(100) NOT NULL,
                         PRIMARY KEY (id),
                         UNIQUE KEY email_unique(email),
                         INDEX name_index(name)
) ENGINE = InnoDB;

-- Menambah index pada tabel existing
ALTER TABLE sellers ADD INDEX name_index(name);

-- Composite index (multiple columns)
ALTER TABLE sellers ADD INDEX name_email_index(name, email);

-- Menghapus index
ALTER TABLE sellers DROP INDEX name_index;

-- Melihat index dalam tabel
SHOW INDEX FROM sellers;


-- ===============================================
-- 25. FULL TEXT SEARCH - Pencarian Teks
-- ===============================================

-- Menambah full text index
ALTER TABLE product
    ADD FULLTEXT product_search(name, description);

-- Natural language mode (default)
SELECT * FROM product
WHERE MATCH(name, description)
    AGAINST ('ayam' IN NATURAL LANGUAGE MODE);

-- Boolean mode (untuk operator AND, OR, NOT)
SELECT * FROM product
WHERE MATCH(name, description)
    AGAINST ('+ayam -goreng' IN BOOLEAN MODE);

-- Query expansion (pencarian diperluas)
SELECT * FROM product
WHERE MATCH(name, description)
    AGAINST ('ayam' WITH QUERY EXPANSION);


-- ===============================================
-- 26. FOREIGN KEY - Relasi Antar Tabel
-- ===============================================

-- Membuat foreign key saat create table
CREATE TABLE wishlist (
                          id INT NOT NULL AUTO_INCREMENT,
                          id_product VARCHAR(10) NOT NULL,
                          description TEXT,
                          PRIMARY KEY (id),
                          CONSTRAINT fk_wishlist_product
                              FOREIGN KEY (id_product) REFERENCES product(id)
) ENGINE = InnoDB;

-- Menambah foreign key pada tabel existing
ALTER TABLE wishlist
    ADD CONSTRAINT fk_wishlist_product
        FOREIGN KEY (id_product) REFERENCES product(id);

-- Foreign key dengan CASCADE (auto update/delete)
ALTER TABLE wishlist
    ADD CONSTRAINT fk_wishlist_product
        FOREIGN KEY (id_product) REFERENCES product(id)
            ON DELETE CASCADE
            ON UPDATE CASCADE;

-- Menghapus foreign key
ALTER TABLE wishlist DROP CONSTRAINT fk_wishlist_product;

-- Opsi ON DELETE dan ON UPDATE:
-- CASCADE    : Ikut update/delete
-- RESTRICT   : Tolak jika ada referensi
-- SET NULL   : Set NULL jika parent dihapus
-- NO ACTION  : Sama dengan RESTRICT


-- ===============================================
-- 27. JOIN - Menggabungkan Tabel
-- ===============================================

-- INNER JOIN - Hanya data yang match di kedua tabel
SELECT * FROM category
                  INNER JOIN product ON (product.id_category = category.id);

-- Join dengan select specific columns
SELECT
    ws.id AS id_wishlist,
    p.id AS id_product,
    p.name,
    ws.description
FROM wishlist AS ws
         INNER JOIN product AS p ON (ws.id_product = p.id);

-- Multiple joins
SELECT
    customer.email,
    product.id,
    product.name,
    wishlist.description
FROM wishlist
         INNER JOIN product ON (product.id = wishlist.id_product)
         INNER JOIN customer ON (customer.id = wishlist.id_customer);

-- LEFT JOIN - Semua data dari tabel kiri + matching dari kanan
SELECT * FROM category
                  LEFT JOIN product ON (product.id_category = category.id);

-- RIGHT JOIN - Semua data dari tabel kanan + matching dari kiri
SELECT * FROM category
                  RIGHT JOIN product ON (product.id_category = category.id);

-- CROSS JOIN - Cartesian product (semua kombinasi)
SELECT * FROM category CROSS JOIN product;

-- Self join - Join tabel dengan dirinya sendiri
SELECT
    n1.id,
    n2.id,
    (n1.id * n2.id) AS multiplication
FROM number AS n1
         CROSS JOIN number AS n2
ORDER BY n1.id, n2.id;


-- ===============================================
-- 28. RELATIONSHIP TYPES - Jenis Relasi
-- ===============================================

-- ONE TO ONE (1:1)
-- Satu customer punya satu wallet
CREATE TABLE wallet (
                        id INT NOT NULL AUTO_INCREMENT,
                        id_customer INT NOT NULL,
                        balance INT NOT NULL DEFAULT 0,
                        PRIMARY KEY (id),
                        UNIQUE KEY id_customer_unique(id_customer),
                        FOREIGN KEY (id_customer) REFERENCES customer(id)
) ENGINE = InnoDB;

-- ONE TO MANY (1:N)
-- Satu category punya banyak product
CREATE TABLE category (
                          id VARCHAR(10) NOT NULL,
                          name VARCHAR(100) NOT NULL,
                          PRIMARY KEY (id)
) ENGINE = InnoDB;

ALTER TABLE product
    ADD COLUMN id_category VARCHAR(10),
ADD CONSTRAINT fk_product_category
    FOREIGN KEY (id_category) REFERENCES category(id);

-- MANY TO MANY (N:M)
-- Menggunakan junction/pivot table
CREATE TABLE orders (
                        id INT NOT NULL AUTO_INCREMENT,
                        total INT NOT NULL,
                        order_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE orders_detail (
                               id_product VARCHAR(10) NOT NULL,
                               id_order INT NOT NULL,
                               price INT NOT NULL,
                               quantity INT NOT NULL,
                               PRIMARY KEY (id_product, id_order),
                               FOREIGN KEY (id_product) REFERENCES product(id),
                               FOREIGN KEY (id_order) REFERENCES orders(id)
) ENGINE = InnoDB;


-- ===============================================
-- 29. SUBQUERY - Query di Dalam Query
-- ===============================================

-- Subquery di WHERE clause
SELECT * FROM product
WHERE price > (SELECT AVG(price) FROM product);

-- Subquery untuk mendapatkan max price
SELECT * FROM product
WHERE price = (SELECT MAX(price) FROM product);

-- Subquery di FROM clause
SELECT MAX(cp.price)
FROM (
         SELECT price FROM category
                               INNER JOIN product ON (product.id_category = category.id)
     ) AS cp;

-- Subquery dengan IN
SELECT * FROM product
WHERE id_category IN (
    SELECT id FROM category WHERE name = 'makanan'
);


-- ===============================================
-- 30. SET OPERATORS - Operasi Set
-- ===============================================

-- UNION - Gabung hasil tanpa duplikat
SELECT email FROM customer
UNION
SELECT email FROM guestbooks;

-- UNION ALL - Gabung hasil dengan duplikat
SELECT email FROM customer
UNION ALL
SELECT email FROM guestbooks;

-- Count dengan UNION ALL
SELECT email, COUNT(emails.email)
FROM (
         SELECT email FROM customer
         UNION ALL
         SELECT email FROM guestbooks
     ) AS emails
GROUP BY emails.email;

-- INTERSECT - Data yang ada di kedua query
SELECT DISTINCT email FROM customer
WHERE email IN (SELECT DISTINCT email FROM guestbooks);

-- Atau menggunakan INNER JOIN
SELECT DISTINCT customer.email
FROM customer
         INNER JOIN guestbooks ON (guestbooks.email = customer.email);

-- MINUS/EXCEPT - Data di query pertama tapi tidak di kedua
SELECT DISTINCT customer.email
FROM customer
         LEFT JOIN guestbooks ON (customer.email = guestbooks.email)
WHERE guestbooks.email IS NULL;


-- ===============================================
-- 31. TRANSACTION - Transaksi Database
-- ===============================================

-- Start transaction
START TRANSACTION;

-- Lakukan operasi database
INSERT INTO guestbooks(email, title, content)
VALUES ('contoh@gmail.com', 'Contoh', 'Contoh');

UPDATE product SET quantity = quantity - 1 WHERE id = 'p001';

-- Commit - Simpan perubahan
COMMIT;

-- Atau Rollback - Batalkan perubahan
ROLLBACK;

-- Contoh transaction lengkap
START TRANSACTION;
INSERT INTO orders(total) VALUES (50000);
INSERT INTO orders_detail(id_product, id_order, price, quantity)
VALUES ('p001', LAST_INSERT_ID(), 25000, 2);
UPDATE product SET quantity = quantity - 2 WHERE id = 'p001';
COMMIT;


-- ===============================================
-- 32. LOCKING - Mengunci Data
-- ===============================================

-- FOR UPDATE - Lock row untuk update
START TRANSACTION;
SELECT * FROM product WHERE id = 'p001' FOR UPDATE;
UPDATE product SET quantity = quantity - 10 WHERE id = 'p001';
COMMIT;

-- FOR SHARE - Lock row untuk read (shared lock)
START TRANSACTION;
SELECT * FROM product WHERE id = 'p001' FOR SHARE;
-- Row bisa dibaca tapi tidak bisa diupdate oleh session lain
COMMIT;

-- LOCK TABLES - Lock entire table
-- READ lock - Hanya bisa read
LOCK TABLES product READ;
SELECT * FROM product;
UNLOCK TABLES;

-- WRITE lock - Bisa read dan write
LOCK TABLES product WRITE;
UPDATE product SET quantity = 100 WHERE id = 'p001';
UNLOCK TABLES;

-- LOCK INSTANCE - Lock seluruh instance
LOCK INSTANCE FOR BACKUP;
-- Backup database
UNLOCK INSTANCE;


-- ===============================================
-- 33. USER MANAGEMENT - Manajemen User
-- ===============================================

-- Membuat user baru
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'username'@'%' IDENTIFIED BY 'password';  -- Dari host manapun

-- Memberikan hak akses (GRANT)
-- Akses SELECT saja
GRANT SELECT ON database_name.* TO 'username'@'localhost';

-- Multiple privileges
GRANT SELECT, INSERT, UPDATE, DELETE ON database_name.*
    TO 'username'@'localhost';

-- All privileges
GRANT ALL PRIVILEGES ON database_name.* TO 'username'@'localhost';

-- Grant pada tabel tertentu
GRANT SELECT ON database_name.table_name TO 'username'@'localhost';

-- Melihat hak akses user
SHOW GRANTS FOR 'username'@'localhost';

-- Mencabut hak akses (REVOKE)
REVOKE SELECT ON database_name.* FROM 'username'@'localhost';
REVOKE ALL PRIVILEGES ON database_name.* FROM 'username'@'localhost';

-- Menghapus user
DROP USER 'username'@'localhost';

-- Mengubah password user
ALTER USER 'username'@'localhost' IDENTIFIED BY 'new_password';

-- Reload privileges
FLUSH PRIVILEGES;


-- ===============================================
-- 34. BACKUP & RESTORE
-- ===============================================

-- BACKUP DATABASE (jalankan di terminal/cmd, bukan di MySQL)
-- Format: mysqldump -u username -p database_name > backup_file.sql

-- Backup single database
mysqldump --user=root --password --single-transaction database_name > backup.sql

-- Backup dengan opsi tambahan
mysqldump --user=root --password --set-gtid-purged=OFF --single-transaction
database_name --result-file=/path/to/backup.sql

-- Backup specific tables
mysqldump --user=root --password database_name table1 table2 > backup.sql

-- Backup all databases
mysqldump --user=root --password --all-databases > all_backup.sql

-- RESTORE DATABASE (jalankan di terminal/cmd)
-- Format: mysql -u username -p database_name < backup_file.sql

-- Restore database
mysql --user=root --password database_name < backup.sql

-- Atau dari dalam MySQL
SOURCE /path/to/backup.sql;

-- Restore dengan create database
mysql --user=root --password < backup.sql


-- ===============================================
-- 35. TIPS & BEST PRACTICES
-- ===============================================

/*
1. NAMING CONVENTIONS
   - Gunakan lowercase untuk nama tabel dan kolom
   - Gunakan underscore untuk pemisah (snake_case)
   - Nama tabel gunakan singular atau plural konsisten
   - Primary key biasanya bernama 'id'
   - Foreign key: id_nama_tabel_referensi

2. DATA TYPES
   - INT: untuk angka bulat
   - VARCHAR: untuk string dengan panjang variabel
   - TEXT: untuk string panjang/artikel
   - DECIMAL: untuk angka desimal/uang
   - TIMESTAMP/DATETIME: untuk tanggal dan waktu
   - ENUM: untuk pilihan terbatas
   - BOOLEAN: simpan sebagai TINYINT(1)

3. INDEXING
   - Buat index pada kolom yang sering di-search/join
   - Jangan terlalu banyak index (memperlambat INSERT/UPDATE)
   - Primary key otomatis terindex
   - Foreign key sebaiknya di-index

4. PERFORMANCE
   - Gunakan LIMIT untuk membatasi hasil query
   - Hindari SELECT * jika tidak perlu semua kolom
   - Gunakan EXPLAIN untuk analisis query
   - Buat index pada kolom WHERE dan JOIN
   - Hindari fungsi pada WHERE clause
   - Gunakan JOIN dibanding subquery jika memungkinkan

5. SECURITY
   - Jangan simpan password plain text
   - Gunakan prepared statement untuk prevent SQL injection
   - Beri hak akses user sesuai kebutuhan (principle of least privilege)
   - Backup database secara regular

6. TRANSACTIONS
   - Gunakan transaction untuk operasi yang harus atomic
   - Selalu COMMIT atau ROLLBACK
   - Jangan transaction terlalu lama (lock resources)

7. RELATIONSHIPS
   - Gunakan foreign key untuk data integrity
   - CASCADE hati-hati, bisa hapus banyak data
   - Untuk many-to-many, buat junction table

8. MAINTENANCE
   - Backup database secara teratur
   - Monitor slow queries
   - Optimize dan analyze table secara berkala
   - Update MySQL version untuk bug fix & performance
*/


-- ===============================================
-- 36. COMMON ERRORS & SOLUTIONS
-- ===============================================

/*
ERROR 1062: Duplicate entry
- Solusi: Data sudah ada (PRIMARY KEY atau UNIQUE constraint)

ERROR 1064: Syntax error
- Solusi: Cek syntax SQL, kata kunci reserved, tanda kutip

ERROR 1451: Cannot delete parent row (foreign key constraint)
- Solusi: Hapus child rows dulu, atau gunakan ON DELETE CASCADE

ERROR 1452: Cannot add child row (foreign key constraint)
- Solusi: Pastikan parent row exists, atau nonaktifkan sementara FK check

ERROR 1146: Table doesn't exist
- Solusi: Pastikan nama tabel benar dan USE database yang tepat

ERROR 1054: Unknown column
- Solusi: Cek nama kolom, pastikan exists di tabel

ERROR 1406: Data too long for column
- Solusi: Perbesar size kolom atau kurangi data

ERROR 1366: Incorrect integer value
- Solusi: Pastikan input sesuai tipe data kolom
*/


-- ===============================================
-- 37. USEFUL QUERIES - Query Berguna
-- ===============================================

-- Melihat ukuran database
SELECT
    table_schema AS 'Database',
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)'
FROM information_schema.TABLES
GROUP BY table_schema;

-- Melihat ukuran setiap tabel
SELECT
    table_name AS 'Table',
    ROUND(((data_length + index_length) / 1024 / 1024), 2) AS 'Size (MB)'
FROM information_schema.TABLES
WHERE table_schema = 'database_name'
ORDER BY (data_length + index_length) DESC;

-- Melihat jumlah row setiap tabel
SELECT
    table_name AS 'Table',
    table_rows AS 'Rows'
FROM information_schema.TABLES
WHERE table_schema = 'database_name'
ORDER BY table_rows DESC;

-- Melihat semua constraint dalam database
SELECT
    CONSTRAINT_NAME,
    TABLE_NAME,
    CONSTRAINT_TYPE
FROM information_schema.TABLE_CONSTRAINTS
WHERE table_schema = 'database_name';

-- Melihat column info detail
SELECT
    COLUMN_NAME,
    COLUMN_TYPE,
    IS_NULLABLE,
    COLUMN_DEFAULT
FROM information_schema.COLUMNS
WHERE table_schema = 'database_name'
  AND table_name = 'table_name';

-- Find duplicate data
SELECT
    column_name,
    COUNT(*) AS jumlah
FROM table_name
GROUP BY column_name
HAVING jumlah > 1;

-- Explain query untuk optimization
EXPLAIN SELECT * FROM product WHERE price > 10000;

-- Show running processes
SHOW PROCESSLIST;

-- Kill process
KILL process_id;

-- Show variables
SHOW VARIABLES LIKE 'max_connections';

-- Set variable
SET GLOBAL max_connections = 200;


-- ===============================================
-- END OF CHEAT SHEET
-- ===============================================