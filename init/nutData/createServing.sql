DROP TABLE IF EXISTS serving;

CREATE TABLE serving (
    NDB_No integer REFERENCES product (NDB_No) ON DELETE CASCADE,
    Serving_Size VARCHAR(10),
    Serving_Size_UOM VARCHAR(4),
    Household_Serving VARCHAR(10),
    Household_Serving_UOM TEXT,
    Preparation_State VARCHAR(20)
);

\copy serving from '/home/steven/Downloads/nutData/data/Serving_size.csv' DELIMITER ',' CSV
