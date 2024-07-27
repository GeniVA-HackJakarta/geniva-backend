CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    kyc BOOLEAN NOT NULL,
    saving VARCHAR(255) NOT NULL
);

CREATE INDEX idx_users_email ON Users(email);

INSERT INTO Users (name, email, password, phone_number, kyc, saving) VALUES
('Example User', 'testing@gmail.com', 'example_password', '0811223344', 'FALSE', 'LOW'),
('Austin Gabriel Pardosi', 'gabrielpardosi26@gmail.com', '$2a$10$6CQ8s7jlRBDy420MTbf5HOpoBa57ZdxqEosSP/7mruo2zjJ4t.eXe', '0811223344', 'TRUE', 'LOW'),
('Austin Pardosi', '13521084@std.stei.itb.ac.id', '$2a$10$NT0mtPl74V.yupQbWA0J3u9rFyvNpwL.skt1e3pb2KoXHl1PD1aw6', '0811223344', 'FALSE', 'HIGH'),
('Fikri Naufal Hamdi', 'fiknaufalh@gmail.com', '$2a$10$JC.YCq3aF3etjMF.bFl0vuKB.aHbrnZBy9qpNJv9/gxvybLTCoAde', '0811223344', 'FALSE', 'MEDIUM'),
('Fikri Naufal', 'fiknaufah@gmail.com', '$2a$10$JC.YCq3aF3etjMF.bFl0vuKB.aHbrnZBy9qpNJv9/gxvybLTCoAde', '0811223344', 'FALSE', 'LOW');

CREATE TABLE IF NOT EXISTS Histories (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    order_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    pickup_location VARCHAR(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    fare NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

INSERT INTO Histories (user_id, service_type, pickup_location, destination, fare, status) VALUES
(2, 'GrabBike', '123 Main St', '456 Elm St', 15000.75, 'completed'),
(2, 'GrabCar', '789 Oak St', '101 Pine St', 25000.50, 'completed'),
(2, 'GrabFood', '135 Second St', '246 Third St', 50000.00, 'pending'),
(2, 'GrabBike', '369 Fifth St', '481 Sixth St', 10000.25, 'cancelled'),
(2, 'GrabCar', '753 Seventh St', '864 Eighth St', 30000.00, 'completed');

CREATE TABLE IF NOT EXISTS Drivers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    vehicle_id INT NOT NULL
);

INSERT INTO Drivers (name, vehicle_id) VALUES 
('Driver 1', 1),
('Driver 2', 2),
('Driver 3', 3),
('Driver 4', 4),
('Driver 5', 5);

CREATE TABLE IF NOT EXISTS Vehicles (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL
);

INSERT INTO Vehicles (type) VALUES 
('Motorcycle'),
('Car'),
('Food Delivery');

CREATE TABLE IF NOT EXISTS Orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    driver_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    food_item_id INT,
    total_price NUMERIC(10, 2) NOT NULL,
    discount NUMERIC(10, 2),
    discount_name VARCHAR(255),
    final_price NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    order_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (driver_id) REFERENCES Drivers(id),
    FOREIGN KEY (vehicle_id) REFERENCES Vehicles(id)
);

INSERT INTO Orders (user_id, driver_id, vehicle_id, total_price, discount, discount_name, final_price, status, order_type) VALUES
(2, 1, 1, 20000.00, 5000.00, 'Promo1', 15000.00, 'completed', 'bike'),
(2, 2, 2, 30000.00, 7000.00, 'Promo2', 23000.00, 'completed', 'car'),
(2, 3, 3, 50000.00, 10000.00, 'Promo3', 40000.00, 'pending', 'food'),
(2, 4, 1, 25000.00, 5000.00, 'Promo4', 20000.00, 'cancelled', 'bike'),
(2, 5, 2, 35000.00, 8000.00, 'Promo5', 27000.00, 'completed', 'car');