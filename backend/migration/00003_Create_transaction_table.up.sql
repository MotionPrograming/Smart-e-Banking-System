-- +migrate Up
CREATE TABLE transactions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    from_account_id BIGINT NULL,
    to_account_id BIGINT NULL,
    amount DECIMAL(20, 6) NOT NULL,
    currency VARCHAR(10) DEFAULT 'BDT',
    type VARCHAR(20) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    reference VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),
    FOREIGN KEY (to_account_id) REFERENCES accounts(id),
    INDEX idx_from_account (from_account_id),
    INDEX idx_to_account (to_account_id),
    INDEX idx_created_at (created_at)
);