CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    native_lang TEXT NOT NULL,
    target_lang TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS learner_profiles (
    user_id TEXT PRIMARY KEY REFERENCES users (id),
    theta REAL DEFAULT 0.0,
    cefr_level TEXT DEFAULT 'A1',
    vocab_map TEXT DEFAULT '{}',
    grammar_map TEXT DEFAULT '{}',
    pragmatics_map TEXT DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS sentences (
    id TEXT PRIMARY KEY,
    text_native TEXT NOT NULL,
    text_target TEXT,
    cefr_level TEXT,
    topic TEXT,
    tags TEXT DEFAULT '[]'
);

CREATE TABLE IF NOT EXISTS attempts (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users (id),
    sentence_id TEXT REFERENCES sentences (id),
    user_input TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    evaluation TEXT DEFAULT '{}'
);

CREATE INDEX IF NOT EXISTS idx_attempts_user_id ON attempts (user_id);

CREATE INDEX IF NOT EXISTS idx_attempts_sentence_id ON attempts (sentence_id);