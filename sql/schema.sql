CREATE TABLE request (
    id INT AUTO_INCREMENT PRIMARY KEY,
    d_t DATE(50),
    server_name VARCHAR(100)
);

INSERT INTO req_log (d_t, server_name) VALUES
('00/00/00:00:00:00', '0.0.0.0');
