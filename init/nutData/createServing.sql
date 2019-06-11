DROP TABLE IF EXISTS serving;

CREATE TABLE serving (
    NDB_No TEXT,
    Serving_Size TEXT,
    Serving_Size_UOM VARCHAR(20),
    Househould_Serving TEXT,
    Househould_Serving_UOM TEXT,
    Preparation_State VARCHAR(20)
);


\copy serving from '/home/steven/Downloads/nutData/data/Serving_size.csv' DELIMITER ',' CSV
