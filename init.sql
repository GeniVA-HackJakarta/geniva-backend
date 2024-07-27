CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE INDEX idx_users_email ON Users(email);

INSERT INTO Users (name, email, password) VALUES
('Example User', 'testing@gmail.com', 'example_password'),
('Austin Gabriel Pardosi', 'gabrielpardosi26@gmail.com', '$2a$10$6CQ8s7jlRBDy420MTbf5HOpoBa57ZdxqEosSP/7mruo2zjJ4t.eXe'),
('Austin Pardosi', '13521084@std.stei.itb.ac.id', '$2a$10$NT0mtPl74V.yupQbWA0J3u9rFyvNpwL.skt1e3pb2KoXHl1PD1aw6'),
('Fikri Naufal Hamdi', 'fiknaufalh@gmail.com', '$2a$10$JC.YCq3aF3etjMF.bFl0vuKB.aHbrnZBy9qpNJv9/gxvybLTCoAde'),
('Fikri Naufal', 'fiknaufah@gmail.com', '$2a$10$JC.YCq3aF3etjMF.bFl0vuKB.aHbrnZBy9qpNJv9/gxvybLTCoAde');
