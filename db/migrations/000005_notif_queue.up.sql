CREATE TABLE IF NOT EXISTS queue (
	id INTEGER PRIMARY KEY,
	email TEXT NOT NULL,
    subject TEXT NOT NULL,
    body TEXT NOT NULL,
    sent BOOLEAN NOT NULL DEFAULT false,
    retry TEXT NOT NULL,
    last_attempt DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);