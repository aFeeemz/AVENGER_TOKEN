CREATE TABLE recipe (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    nama_resep VARCHAR(255) NOT NULL,
    deskripsi_resep VARCHAR(255) NOT NULL,
    waktu_masak VARCHAR(255),
    rating FLOAT
);

INSERT INTO recipe (nama_resep, deskripsi_resep, waktu_masak, rating) VALUES
('Spaghetti Bolognese', 'Classic Italian pasta dish with rich meat sauce', '45 minutes', 4.8),
('Chicken Curry', 'Spicy and creamy chicken curry with coconut milk', '60 minutes', 4.5),
('Vegetable Stir Fry', 'Quick and easy stir-fried vegetables with soy sauce', '20 minutes', 4.2),
('Beef Stew', 'Hearty beef stew with potatoes and carrots', '120 minutes', 4.9),
('Chocolate Cake', 'Moist and rich chocolate cake with chocolate frosting', '90 minutes', 4.7);
