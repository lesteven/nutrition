
CREATE TABLE derivation (
    Derivation_Code VARCHAR(4) PRIMARY KEY,
    Derivation_Code_Description TEXT
);

\copy derivation from '/home/steven/Downloads/nutData/data/Derivation_Code_Description.csv' DELIMITER ',' CSV
