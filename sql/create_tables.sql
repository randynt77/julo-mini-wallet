CREATE TABLE IF NOT EXISTS wallet (
    id serial NOT NULL,
    owned_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    enabled_at TIME NULL,
    balance FLOAT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS transaction (
    id serial NOT NULL,
    status VARCHAR(255) NOT NULL,
    transacted_at TIMESTAMP NOT NULL,
    type VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    reference_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS deposit (
    id serial NOT NULL,
    deposited_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    deposited_at TIMESTAMP NOT NULL,
    amount FLOAT NOT NULL,
    reference_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS withdrawal (
    id VARCHAR(255) NOT NULL,
    withdrawn_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    withdrawn_at TIMESTAMP NOT NULL,
    amount FLOAT NOT NULL,
    reference_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);


