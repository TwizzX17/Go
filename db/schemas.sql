CREATE TABLE symbols(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_ts DATE,
    name VARCHAR(255),
    symbol VARCHAR(50) UNIQUE,
    exchange VARCHAR(50),
    asset_type VARCHAR(50),
    status BOOLEAN,
    deleted_ts timestamp
);