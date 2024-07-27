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
