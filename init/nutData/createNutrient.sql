DROP TABLE IF EXISTS nutrient;

CREATE TABLE nutrient (
    NDB_No TEXT,
    Nutrient_Code TEXT,
    Nutrient_Name TEXT,
    Derivation_Code TEXT,
    Output_Value TEXT,
    Output_UOM VARCHAR(10)
);


\copy nutrient from '/home/steven/Downloads/nutData/data/Nutrients.csv' DELIMITER ',' CSV
