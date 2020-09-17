CREATE TABLE Account(
    id              BIGSERIAL   NOT NULL    PRIMARY KEY,
    total_amount    FLOAT   
);

CREATE TABLE Customer(
    id               BIGSERIAL       NOT NULL PRIMARY KEY,
    fName            VARCHAR(30)     NOT NULL,
    lName            VARCHAR(30)     NOT NULL,
    mobile           VARCHAR(11)     NOT NULL,
    email            VARCHAR(50)     NOT NULL,
    pass             VARCHAR(255)    NOT NULL,
    UNIQUE(email)
);

CREATE TABLE Bill(
    id                  BIGSERIAL       NOT NULL PRIMARY KEY,
    customer_id         BIGINT          NOT NULL REFERENCES Customer (id),   
    mobile              VARCHAR(11)     NOT NULL,
    bill_type           VARCHAR(30)     NOT NULL,
    equipment_count     INT             NOT NULL,
    amount              FLOAT           NOT NULL,
    account_id          BIGINT          NOT NULL REFERENCES Account (id),
    payment_method      VARCHAR(30)     NOT NULL,
    submit_date         DATE
);

INSERT INTO Account(total_amount) VALUE(0.0);

