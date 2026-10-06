CREATE TYPE copy_status   AS ENUM ('AVAILABLE', 'BORROWED', 'LOST');
CREATE TYPE member_status AS ENUM ('ACTIVE', 'SUSPENDED');

CREATE TABLE IF NOT EXISTS members (
    id         UUID          PRIMARY KEY,
    name       VARCHAR(100)  NOT NULL CHECK (name <> ''),
    email      TEXT          NOT NULL,
    status     member_status NOT NULL DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ   NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ   NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_members_email      ON members (email) WHERE deleted_at IS NULL;
CREATE        INDEX IF NOT EXISTS idx_members_deleted_at ON members (deleted_at);

CREATE TABLE IF NOT EXISTS books (
    id         UUID        PRIMARY KEY,
    isbn       CHAR(13)    NOT NULL CHECK (isbn ~ '^[0-9]{13}$'),
    title      TEXT        NOT NULL,
    author     TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_books_isbn       ON books (isbn) WHERE deleted_at IS NULL;
CREATE        INDEX IF NOT EXISTS idx_books_deleted_at ON books (deleted_at);

CREATE TABLE IF NOT EXISTS book_copies (
    id         UUID        PRIMARY KEY,
    book_id    UUID        NOT NULL REFERENCES books (id),
    barcode    TEXT        NOT NULL CHECK (barcode <> ''),
    status     copy_status NOT NULL DEFAULT 'AVAILABLE',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_book_copies_barcode    ON book_copies (barcode) WHERE deleted_at IS NULL;
CREATE        INDEX IF NOT EXISTS idx_book_copies_book_id    ON book_copies (book_id);
CREATE        INDEX IF NOT EXISTS idx_book_copies_deleted_at ON book_copies (deleted_at);

CREATE TABLE IF NOT EXISTS loans (
    id          UUID        PRIMARY KEY,
    member_id   UUID        NOT NULL REFERENCES members (id),
    copy_id     UUID        NOT NULL REFERENCES book_copies (id),
    borrowed_at TIMESTAMPTZ NOT NULL,
    due_at      TIMESTAMPTZ NOT NULL,
    returned_at TIMESTAMPTZ,
    renew_count INTEGER     NOT NULL DEFAULT 0 CHECK (renew_count BETWEEN 0 AND 1),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ,
    CHECK (due_at > borrowed_at),
    CHECK (returned_at IS NULL OR returned_at >= borrowed_at)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_loans_active_copy ON loans (copy_id) WHERE returned_at IS NULL;
CREATE        INDEX IF NOT EXISTS idx_loans_member_id   ON loans (member_id);
CREATE        INDEX IF NOT EXISTS idx_loans_deleted_at  ON loans (deleted_at);

CREATE TABLE IF NOT EXISTS fines (
    id         UUID        PRIMARY KEY,
    loan_id    UUID        NOT NULL REFERENCES loans (id),
    member_id  UUID        NOT NULL REFERENCES members (id),
    amount     BIGINT      NOT NULL CHECK (amount > 0),
    paid_at    TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_fines_loan_id    ON fines (loan_id) WHERE deleted_at IS NULL;
CREATE        INDEX IF NOT EXISTS idx_fines_member_id  ON fines (member_id);
CREATE        INDEX IF NOT EXISTS idx_fines_deleted_at ON fines (deleted_at);
