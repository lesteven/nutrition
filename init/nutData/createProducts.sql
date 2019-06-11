DROP TABLE product;

CREATE TABLE product (
    NDB_No TEXT,
    Long_Name VARCHAR(200),
    Data_Source TEXT,
    GTIN_UPC TEXT,
    Manufacturer TEXT,
    Date_Modified TEXT,
    Date_available TEXT,
    Ingredients TEXT
);

\copy product from '/home/steven/Downloads/nutData/data/Products.csv' DELIMITER ',' CSV
