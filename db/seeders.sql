-- =============================================================================
-- SEED DATA FOR CINEMA TICKETING SYSTEM (POSTGRESQL)
-- =============================================================================

-- 1. Users (Password: 'password123' bcrypt hashed)
INSERT INTO users (name, email, password, role)
VALUES 
    ('Admin', 'admin@example.com', '$2a$10$Q7iWk8kHQy/sZ6L6q7l1e.lJvJm1fK4t0hK1Z1z0L0jG5R8h5c7G6', 'admin'),
    ('User', 'user@example.com', '$2a$10$Q7iWk8kHQy/sZ6L6q7l1e.lJvJm1fK4t0hK1Z1z0L0jG5R8h5c7G6', 'customer')
ON CONFLICT (email) DO NOTHING;

-- 2. Cinemas
INSERT INTO cinemas (id, name, city, address)
VALUES 
    (1, 'XXI Grand Indonesia', 'Jakarta', 'Grand Indonesia Shopping Town, Jl. M.H. Thamrin No.1, Jakarta Pusat'),
    (2, 'Cinepolis Paragon Mall', 'Semarang', 'Pollux Mall Paragon, Jl. Pemuda No.118, Sekayu, Kota Semarang'),
    (3, 'XXI Tunjungan Plaza', 'Surabaya', 'Tunjungan Plaza 5, Jl. Embong Malang No.1-30, Surabaya'),
    (4, 'XXI Beachwalk Bali', 'Denpasar', 'Beachwalk Shopping Center, Jl. Pantai Kuta, Badung, Bali')
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for cinemas
SELECT setval('cinemas_id_seq', (SELECT MAX(id) FROM cinemas));

-- 3. Studios
INSERT INTO studios (id, cinema_id, name, total_seats)
VALUES 
    -- XXI Grand Indonesia (Jakarta)
    (1, 1, 'Studio 1', 60),
    (2, 1, 'Studio 2', 50),
    (3, 1, 'IMAX', 100),
    (4, 1, 'Premiere', 30),
    -- Cinepolis Paragon Mall (Semarang)
    (5, 2, 'Studio 1', 50),
    (6, 2, 'Macro XE', 100),
    -- XXI Tunjungan Plaza (Surabaya)
    (7, 3, 'Studio 1', 60),
    (8, 3, 'IMAX', 100),
    -- XXI Beachwalk Bali (Denpasar)
    (9, 4, 'Studio 1', 50),
    (10, 4, 'Premiere', 30)
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for studios
SELECT setval('studios_id_seq', (SELECT MAX(id) FROM studios));

-- 4. Movies
INSERT INTO movies (id, title, duration_minutes, genre, poster_url)
VALUES 
    (1, 'Inception', 148, 'Sci-Fi, Action', 'https://image.tmdb.org/t/p/w500/edv5CZvWj09upOsy2Y6IwDhK8bt.jpg'),
    (2, 'Interstellar', 169, 'Sci-Fi, Adventure, Drama', 'https://image.tmdb.org/t/p/w500/gEU2QniE6E77NI6lCU6MxlNBvIx.jpg'),
    (3, 'The Dark Knight', 152, 'Action, Crime, Drama', 'https://image.tmdb.org/t/p/w500/qJ2tW6WMUDux911r6m7haRef0WH.jpg'),
    (4, 'Dune: Part Two', 166, 'Sci-Fi, Adventure', 'https://image.tmdb.org/t/p/w500/1pdfLvkbY9ohJlCjQH2CZjjYVvJ.jpg'),
    (5, 'Agak Laen', 119, 'Comedy, Horror', 'https://image.tmdb.org/t/p/w500/agak_laen_poster.jpg')
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for movies
SELECT setval('movies_id_seq', (SELECT MAX(id) FROM movies));
