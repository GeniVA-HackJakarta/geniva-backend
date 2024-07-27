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
    food_item_id INT,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

INSERT INTO Histories (user_id, service_type, pickup_location, destination, fare, status, food_item_id) VALUES
(2, 'GrabBike', '123 Main St', '456 Elm St', 15000.75, 'completed', NULL),
(2, 'GrabCar', '789 Oak St', '101 Pine St', 25000.50, 'completed', NULL),
(2, 'GrabFood', '135 Second St', '246 Third St', 50000.00, 'pending', 1),
(2, 'GrabBike', '369 Fifth St', '481 Sixth St', 10000.25, 'cancelled', NULL),
(2, 'GrabCar', '753 Seventh St', '864 Eighth St', 30000.00, 'completed', NULL);

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

CREATE TABLE IF NOT EXISTS Menus (
    menu_id INT PRIMARY KEY,
    menu_name VARCHAR(255) NOT NULL,
    restaurant_id INT NOT NULL,
    price VARCHAR(10) NOT NULL,
    glucose_levels VARCHAR(10) NOT NULL
);

INSERT INTO Menus (menu_id, menu_name, restaurant_id, price, glucose_levels) VALUES
	(1, 'Oil - Truffle, White', 7421901, '$48.31', 'high'),
	(2, 'Pasta - Penne Primavera, Single', 7419572, '$62.36', 'low'),
	(3, 'Otomegusa Dashi Konbu', 18524271, '$23.23', 'low'),
	(4, 'Mushroom - Morel Frozen', 18346266, '$67.46', 'medium'),
	(5, 'Cocoa Powder - Natural', 18368134, '$65.61', 'low'),
	(6, 'Cheese - Bocconcini', 18368132, '$77.28', 'medium'),
	(7, 'Cheese - Brick With Pepper', 7401990, '$62.01', 'medium'),
	(8, 'Mahi Mahi', 7422631, '$22.36', 'medium'),
	(9, 'Crab - Claws, Snow 16 - 24', 7404738, '$31.17', 'medium'),
	(10, 'Food Colouring - Orange', 18917980, '$7.16', 'medium'),
	(11, 'Mackerel Whole Fresh', 7402859, '$31.51', 'medium'),
	(12, 'Scallops 60/80 Iqf', 7419154, '$46.99', 'high'),
	(13, 'Bread - Wheat Baguette', 7422160, '$23.02', 'low'),
	(14, 'External Supplier', 18478836, '$3.12', 'high'),
	(15, 'Sesame Seed Black', 7420287, '$57.15', 'low'),
	(16, 'Shichimi Togarashi Peppeers', 18665235, '$51.87', 'medium'),
	(17, 'Ham Black Forest', 18571667, '$28.22', 'medium'),
	(18, 'Bread - English Muffin', 7413472, '$1.90', 'high'),
	(19, 'Soup - Campbells Asian Noodle', 7401969, '$3.75', 'high'),
	(20, 'Sponge Cake Mix - Chocolate', 7410695, '$73.51', 'medium'),
	(21, 'Longos - Chicken Wings', 7410290, '$70.04', 'medium'),
	(22, 'Gin - Gilbeys London, Dry', 7405302, '$33.54', 'low'),
	(23, 'Relish', 7423350, '$39.25', 'high'),
	(24, 'Pepsi, 355 Ml', 18323073, '$74.34', 'high'),
	(25, 'Trout - Smoked', 18358536, '$76.99', 'high'),
	(26, 'Wine - Barolo Fontanafredda', 7404585, '$22.21', 'high'),
	(27, 'Chicken - Diced, Cooked', 18615392, '$22.01', 'high'),
	(28, 'Cheese - Victor Et Berthold', 7402935, '$28.36', 'medium'),
	(29, 'Versatainer Nc - 888', 7422751, '$44.27', 'medium'),
	(30, 'Ice Cream Bar - Oreo Sandwich', 7407828, '$13.91', 'low'),
	(31, 'Butter Ripple - Phillips', 7423352, '$51.36', 'high'),
	(32, 'Taro Leaves', 7404571, '$68.83', 'high'),
	(33, 'Jagermeister', 18421271, '$66.02', 'high'),
	(34, 'Steampan Lid', 19022690, '$14.66', 'low'),
	(35, 'Cheese - Pont Couvert', 18538909, '$76.05', 'medium'),
	(36, 'Chicken Breast Wing On', 18445561, '$66.64', 'low'),
	(37, 'Chocolate - Milk', 7424424, '$78.22', 'high'),
	(38, 'Soup - Campbellschix Stew', 18339380, '$68.44', 'high'),
	(39, 'Southern Comfort', 18286687, '$55.21', 'low'),
	(40, 'Split Peas - Green, Dry', 18534498, '$42.19', 'low'),
	(41, 'Chinese Lemon Pork', 7404551, '$42.40', 'medium'),
	(42, 'Chocolate - Unsweetened', 18370496, '$33.17', 'low'),
	(43, 'Okra', 18166724, '$25.75', 'medium'),
	(44, 'Ham - Black Forest', 7423164, '$5.85', 'high'),
	(45, 'Pork - Suckling Pig', 18421275, '$2.76', 'medium'),
	(46, 'Beef - Kindney, Whole', 7417570, '$27.95', 'medium'),
	(47, 'Pears - Bosc', 7403024, '$15.99', 'high'),
	(48, 'Soap - Mr.clean Floor Soap', 18236985, '$63.72', 'medium'),
	(49, 'Campari', 18236984, '$39.55', 'medium'),
	(50, 'Nacho Chips', 7424538, '$62.04', 'medium'),
	(51, 'Pork - Bacon, Sliced', 7401371, '$70.34', 'high'),
	(52, 'Soup - Campbells Tomato Ravioli', 7402602, '$20.15', 'medium'),
	(53, 'Tofu - Firm', 18498083, '$5.10', 'low'),
	(54, 'Bread - Bistro White', 7405273, '$16.01', 'medium'),
	(55, 'Water - Tonic', 18563957, '$79.86', 'medium'),
	(56, 'Horseradish Root', 7404491, '$9.85', 'medium'),
	(57, 'Beef - Tenderloin', 18407558, '$78.79', 'medium'),
	(58, 'Chevril', 18458483, '$42.63', 'medium'),
	(59, 'Liqueur Banana, Ramazzotti', 7403833, '$70.40', 'high'),
	(60, 'Lettuce - Boston Bib', 7400286, '$76.05', 'low'),
	(61, 'Trout - Smoked', 18266425, '$13.09', 'high'),
	(62, 'Tofu - Soft', 7405344, '$67.61', 'high'),
	(63, 'Lentils - Green Le Puy', 7401164, '$53.14', 'low'),
	(64, 'Cheese - Stilton', 18563955, '$18.48', 'low'),
	(65, 'Appetiser - Bought', 18312549, '$29.55', 'high'),
	(66, 'Lettuce - Green Leaf', 7424480, '$53.52', 'low'),
	(67, 'Cake - Lemon Chiffon', 7402466, '$12.88', 'high'),
	(68, 'Soup - Knorr, Chicken Gumbo', 18225791, '$78.31', 'medium'),
	(69, 'Chicken - Livers', 7422971, '$57.00', 'low'),
	(70, 'Doilies - 12, Paper', 18535568, '$62.83', 'medium'),
	(71, 'Cheese - Gouda', 18335732, '$11.91', 'high'),
	(72, 'Wine - Lou Black Shiraz', 7406486, '$68.93', 'medium'),
	(73, 'Syrup - Chocolate', 7402060, '$47.25', 'low'),
	(74, 'Lemonade - Strawberry, 591 Ml', 18264516, '$60.41', 'high'),
	(75, 'Dasheen', 7400576, '$66.07', 'medium'),
	(76, 'Extract - Raspberry', 18836779, '$70.12', 'high'),
	(77, 'Horseradish Root', 7423519, '$40.56', 'high'),
	(78, 'Chocolate - Semi Sweet', 18446348, '$2.53', 'low'),
	(79, 'Tomatoes Tear Drop', 18597282, '$6.75', 'high'),
	(80, 'Milk - Chocolate 500ml', 7404855, '$21.89', 'low'),
	(81, 'Honey - Comb', 18407754, '$17.02', 'medium'),
	(82, 'Beef Striploin Aaa', 7420388, '$34.95', 'high'),
	(83, 'Beef Dry Aged Tenderloin Aaa', 7404953, '$37.40', 'low'),
	(84, 'Tart Shells - Sweet, 3', 18819628, '$60.03', 'high'),
	(85, 'Cheese - Comte', 19307150, '$53.25', 'medium'),
	(86, 'Cheese - Cottage Cheese', 7404295, '$38.37', 'low'),
	(87, 'Melon - Cantaloupe', 18623698, '$4.46', 'low'),
	(88, 'Wine - Red, Cabernet Merlot', 18386858, '$13.14', 'medium'),
	(89, 'Creme De Cacao Mcguines', 18258558, '$34.77', 'medium'),
	(90, 'Spinach - Frozen', 7420079, '$58.77', 'low'),
	(91, 'Coffee - French Vanilla Frothy', 7420284, '$69.20', 'high'),
	(92, 'Pepper - Julienne, Frozen', 7422953, '$58.89', 'high'),
	(93, 'Transfer Sheets', 7415963, '$38.01', 'low'),
	(94, 'Beef - Cow Feet Split', 7419773, '$71.87', 'low'),
	(95, 'Coffee - Decaffeinato Coffee', 18352452, '$25.12', 'low'),
	(96, 'Beef - Top Butt', 18157488, '$8.91', 'medium'),
	(97, 'Potatoes - Peeled', 7402002, '$65.59', 'high'),
	(98, 'Turkey - Oven Roast Breast', 7401166, '$32.41', 'high'),
	(99, 'Trout - Rainbow, Frozen', 7405156, '$57.94', 'low'),
	(100, 'Soup - Campbells Beef Noodle', 7423897, '$48.57', 'medium');



    