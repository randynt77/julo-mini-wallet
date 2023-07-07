CREATE TABLE IF NOT EXISTS wallet (
    id VARCHAR(255) NOT NULL,
    owned_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    enabled_at TIME NULL,
    disabled_at TIME NULL,
    balance FLOAT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS transaction (
    id VARCHAR(255) NOT NULL,
    wallet_id VARCHAR(255) NOT NULL,
    actor_id VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    transacted_at TIMESTAMP NOT NULL,
    type VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    reference_id VARCHAR(255) NOT NULL,
    input_reference_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);


