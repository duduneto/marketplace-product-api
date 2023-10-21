CREATE TABLE Product (
    id UUID PRIMARY KEY
);

-- Table for Product Details
CREATE TABLE ProductDetails (
    detail_id UUID PRIMARY KEY,
    product_id UUID,
    version UUID,
    version_description TEXT,
    title TEXT,
    price DECIMAL(10, 2),
    FOREIGN KEY (product_id) REFERENCES Product(id)
);