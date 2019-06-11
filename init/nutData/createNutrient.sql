
CREATE TABLE nutrient (
    NDB_No INTEGER REFERENCES product (NDB_No),
    Nutrient_Code SMALLINT,
    Nutrient_Name TEXT,
    Derivation_Code VARCHAR(4) REFERENCES derivation (Derivation_Code),
    Output_Value REAL,
    Output_UOM VARCHAR(4)
);
