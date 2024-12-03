CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE
);
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE
);
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
--  inserts
INSERT INTO users (first_name, last_name, email, password_hash, created_at, is_active) VALUES
('Alice', 'Smith', 'alice.smith@example.com', 'hashed_password1', NOW(), TRUE),
('Bob', 'Johnson', 'bob.johnson@example.com', 'hashed_password2', NOW(), TRUE),
('Charlie', 'Williams', 'charlie.williams@example.com', 'hashed_password3', NOW(), TRUE),
('Daisy', 'Brown', 'daisy.brown@example.com', 'hashed_password4', NOW(), TRUE),
('Ethan', 'Jones', 'ethan.jones@example.com', 'hashed_password5', NOW(), TRUE),
('Fiona', 'Garcia', 'fiona.garcia@example.com', 'hashed_password6', NOW(), TRUE),
('George', 'Martinez', 'george.martinez@example.com', 'hashed_password7', NOW(), TRUE),
('Hannah', 'Rodriguez', 'hannah.rodriguez@example.com', 'hashed_password8', NOW(), TRUE),
('Ivy', 'Lopez', 'ivy.lopez@example.com', 'hashed_password9', NOW(), TRUE),
('Jack', 'Lee', 'jack.lee@example.com', 'hashed_password10', NOW(), TRUE),
('Katie', 'Walker', 'katie.walker@example.com', 'hashed_password11', NOW(), TRUE),
('Liam', 'Hall', 'liam.hall@example.com', 'hashed_password12', NOW(), TRUE),
('Mia', 'Allen', 'mia.allen@example.com', 'hashed_password13', NOW(), TRUE),
('Noah', 'Young', 'noah.young@example.com', 'hashed_password14', NOW(), TRUE),
('Olivia', 'Hernandez', 'olivia.hernandez@example.com', 'hashed_password15', NOW(), TRUE),
('Parker', 'King', 'parker.king@example.com', 'hashed_password16', NOW(), TRUE),
('Quinn', 'Wright', 'quinn.wright@example.com', 'hashed_password17', NOW(), TRUE),
('Riley', 'Hill', 'riley.hill@example.com', 'hashed_password18', NOW(), TRUE),
('Sophia', 'Scott', 'sophia.scott@example.com', 'hashed_password19', NOW(), TRUE),
('Tyler', 'Green', 'tyler.green@example.com', 'hashed_password20', NOW(), TRUE);

INSERT INTO products (name, description, price, stock_quantity, created_at, is_active) VALUES
('Product A', 'Description of Product A', 19.99, 10, NOW(), TRUE),
('Product B', 'Description of Product B', 24.99, 15, NOW(), TRUE),
('Product C', 'Description of Product C', 29.99, 20, NOW(), TRUE),
('Product D', 'Description of Product D', 34.99, 5, NOW(), TRUE),
('Product E', 'Description of Product E', 39.99, 25, NOW(), TRUE),
('Product F', 'Description of Product F', 44.99, 30, NOW(), TRUE),
('Product G', 'Description of Product G', 49.99, 8, NOW(), TRUE),
('Product H', 'Description of Product H', 54.99, 12, NOW(), TRUE),
('Product I', 'Description of Product I', 59.99, 18, NOW(), TRUE),
('Product J', 'Description of Product J', 64.99, 22, NOW(), TRUE),
('Product K', 'Description of Product K', 69.99, 14, NOW(), TRUE),
('Product L', 'Description of Product L', 74.99, 9, NOW(), TRUE),
('Product M', 'Description of Product M', 79.99, 13, NOW(), TRUE),
('Product N', 'Description of Product N', 84.99, 6, NOW(), TRUE),
('Product O', 'Description of Product O', 89.99, 11, NOW(), TRUE),
('Product P', 'Description of Product P', 94.99, 7, NOW(), TRUE),
('Product Q', 'Description of Product Q', 99.99, 10, NOW(), TRUE),
('Product R', 'Description of Product R', 104.99, 17, NOW(), TRUE),
('Product S', 'Description of Product S', 109.99, 19, NOW(), TRUE),
('Product T', 'Description of Product T', 114.99, 23, NOW(), TRUE);

INSERT INTO orders (user_id, total, created_at) VALUES
(21, 199.99, NOW()),
(22, 299.99, NOW()),
(23, 399.99, NOW()),
(24, 499.99, NOW()),
(25, 599.99, NOW()),
(26, 699.99, NOW()),
(27, 799.99, NOW()),
(28, 899.99, NOW()),
(29, 999.99, NOW()),
(30, 1099.99, NOW()),
(31, 1199.99, NOW()),
(32, 1299.99, NOW()),
(33, 1399.99, NOW()),
(34, 1499.99, NOW()),
(35, 1599.99, NOW()),
(36, 1699.99, NOW()),
(37, 1799.99, NOW()),
(38, 1899.99, NOW()),
(39, 1999.99, NOW()),
(40, 2099.99, NOW());
