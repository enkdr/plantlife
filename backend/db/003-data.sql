-- Insert data into the plant table and capture the generated UUIDs
WITH inserted_plants AS (
    INSERT INTO plant (id, name, description, image, water, sun, germination, flowering, harvest, seed)
    VALUES
        (uuid_generate_v4(), 'Tomato', 'A juicy, red fruit often used in salads and cooking. High in vitamin C and antioxidants.', 'tomato.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Potato', 'A starchy tuber that is a staple food in many cultures. Rich in carbohydrates and potassium.', 'potato.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Carrot', 'A crunchy, orange root vegetable known for its high beta-carotene content.', 'carrot.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Lettuce', 'A leafy green vegetable often used in salads and sandwiches. Low in calories and rich in fiber.', 'lettuce.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Pepper', 'A versatile vegetable that comes in various colors and can be sweet or spicy. High in vitamins A and C.', 'pepper.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Tulip', 'A popular spring-blooming flower known for its bright, cup-shaped blossoms.', 'tulip.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Rose', 'A classic flowering plant admired for its beauty and fragrance. Often associated with love and romance.', 'rose.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Daffodil', 'A bright yellow flower that is a symbol of spring. Known for its trumpet-shaped blooms.', 'daffodil.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Orchid', 'An exotic flowering plant known for its intricate and delicate blossoms. Symbolizes beauty and strength.', 'orchid.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Sunflower', 'A tall plant known for its large, yellow blooms that follow the sun. High in vitamin E and often used to produce oil.', 'sunflower.jpg', 100, 6, 10, 20, 30, 5),
        (uuid_generate_v4(), 'Lavender', 'A fragrant herb known for its purple flowers and calming scent. Often used in aromatherapy and cooking.', 'lavender.jpg', 100, 6, 10, 20, 30, 5)
    RETURNING id, name
),

inserted_stores AS (
    INSERT INTO store (id, name, address, city, postcode, country)
    VALUES
        (uuid_generate_v4(), 'Green Thumb Nursery', '123 Plant St', 'Springfield', '12345', 'USA'),
        (uuid_generate_v4(), 'Urban Jungle', '456 Garden Ave', 'Metropolis', '67890', 'USA'),
        (uuid_generate_v4(), 'Natures Best', '789 Forest Rd', 'Gotham', '10112', 'USA'),
        (uuid_generate_v4(), 'Sunset Gardens', '890 Sunset Blvd', 'Sunset City', '54321', 'USA'),
        (uuid_generate_v4(), 'Evergreen Plants', '234 Pine Rd', 'Evergreen', '98765', 'USA')
    RETURNING id, name
)

-- Insert data into the plant_store table using the captured UUIDs
INSERT INTO plant_store (id, quantity, plant_id, store_id)
SELECT uuid_generate_v4(), quantity, plant_id, store_id
FROM (
    SELECT
        50 AS quantity,
        (SELECT id FROM inserted_plants WHERE name = 'Tomato') AS plant_id,
        (SELECT id FROM inserted_stores WHERE name = 'Green Thumb Nursery') AS store_id
    UNION ALL
    SELECT
        75 AS quantity,
        (SELECT id FROM inserted_plants WHERE name = 'Lettuce') AS plant_id,
        (SELECT id FROM inserted_stores WHERE name = 'Green Thumb Nursery') AS store_id
    UNION ALL
    SELECT
        100 AS quantity,
        (SELECT id FROM inserted_plants WHERE name = 'Carrot') AS plant_id,
        (SELECT id FROM inserted_stores WHERE name = 'Urban Jungle') AS store_id
    UNION ALL
    SELECT
        30 AS quantity,
        (SELECT id FROM inserted_plants WHERE name = 'Sunflower') AS plant_id,
        (SELECT id FROM inserted_stores WHERE name = 'Natures Best') AS store_id
    UNION ALL
    SELECT
        45 AS quantity,
        (SELECT id FROM inserted_plants WHERE name = 'Pepper') AS plant_id,
        (SELECT id FROM inserted_stores WHERE name = 'Natures Best') AS store_id
) AS subquery;
