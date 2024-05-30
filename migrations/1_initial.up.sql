CREATE TABLE sites (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    delete_flag BOOLEAN
);

CREATE TABLE payment_to_company_infos (
    id VARCHAR(36) PRIMARY KEY,
    site_id VARCHAR(36),
    amount INT,
    payed_at DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (site_id) REFERENCES sites(id)
);

-- Insert initial data into sites table
INSERT INTO sites (id, name, url, created_at, updated_at, delete_flag)
VALUES
('site1', 'Example Site', 'http://example.com', NOW(), NOW(), false);

-- Insert initial data into payment_to_company_infos table
INSERT INTO payment_to_company_infos (id, site_id, amount, payed_at, created_at, updated_at)
VALUES
('payment1', 'site1', 1000, NOW(), NOW(), NOW()),
('payment2', 'site1', 2000, NOW(), NOW(), NOW());