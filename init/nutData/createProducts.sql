
CREATE TABLE product (
    NDB_No INTEGER PRIMARY KEY,
    Long_Name VARCHAR(200),
    Data_Source TEXT,
    GTIN_UPC BIGINT,
    Manufacturer TEXT,
    Date_Modified TIMESTAMP,
    Date_available TIMESTAMP,
    Ingredients TEXT
);

\copy product from '/home/steven/Downloads/nutData/data/Products.csv' DELIMITER ',' CSV
