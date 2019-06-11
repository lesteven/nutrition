
CREATE TABLE serving (
    NDB_No integer REFERENCES product (NDB_No),
    Serving_Size REAL,
    Serving_Size_UOM VARCHAR(4),
    Househould_Serving REAL,
    Househould_Serving_UOM TEXT,
    Preparation_State VARCHAR(20)
);


