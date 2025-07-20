-- User queries
-- name: SaveUser :exec
INSERT INTO
    users (
        id,
        email,
        native_lang,
        target_lang
    )
VALUES (?, ?, ?, ?)
ON CONFLICT (id) DO
UPDATE
SET
    email = excluded.email,
    native_lang = excluded.native_lang,
    target_lang = excluded.target_lang;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- LearnerProfile queries
-- name: SaveLearnerProfile :exec
INSERT INTO
    learner_profiles (
        user_id,
        theta,
        cefr_level,
        vocab_map,
        grammar_map,
        pragmatics_map
    )
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT (user_id) DO
UPDATE
SET
    theta = excluded.theta,
    cefr_level = excluded.cefr_level,
    vocab_map = excluded.vocab_map,
    grammar_map = excluded.grammar_map,
    pragmatics_map = excluded.pragmatics_map;

-- name: FindLearnerProfileByUserID :one
SELECT * FROM learner_profiles WHERE user_id = ?;

-- Sentence queries
-- name: SaveSentence :exec
INSERT INTO
    sentences (
        id,
        text_native,
        text_target,
        cefr_level,
        topic,
        tags
    )
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT (id) DO
UPDATE
SET
    text_native = excluded.text_native,
    text_target = excluded.text_target,
    cefr_level = excluded.cefr_level,
    topic = excluded.topic,
    tags = excluded.tags;

-- name: FindSentenceByID :one
SELECT * FROM sentences WHERE id = ?;

-- name: FindNextSentenceForUser :one
SELECT *
FROM sentences
WHERE
    cefr_level = ?
ORDER BY RANDOM()
LIMIT 1;

-- Attempt queries
-- name: SaveAttempt :exec
INSERT INTO
    attempts (
        id,
        user_id,
        sentence_id,
        user_input,
        evaluation
    )
VALUES (?, ?, ?, ?, ?)
ON CONFLICT (id) DO
UPDATE
SET
    user_input = excluded.user_input,
    evaluation = excluded.evaluation;

-- name: FindAttemptByID :one
SELECT * FROM attempts WHERE id = ?;

-- name: FindAttemptByUserAndSentence :one
SELECT * FROM attempts WHERE user_id = ? AND sentence_id = ?;