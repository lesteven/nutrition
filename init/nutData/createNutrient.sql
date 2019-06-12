DROP TABLE IF EXISTS nutrient;

CREATE TABLE nutrient (
    NDB_No integer REFERENCES product (NDB_No) ON DELETE CASCADE,
    Nutrient_Code SMALLINT,
    Nutrient_Name TEXT,
    Derivation_Code VARCHAR(4) REFERENCES derivation (Derivation_Code)\
     ON DELETE CASCADE,
    Output_Value REAL,
    Output_UOM VARCHAR(10)
);

\copy nutrient from '/home/steven/Downloads/nutData/data/Nutrients.csv' DELIMITER ',' CSV
