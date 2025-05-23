CREATE TABLE review_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID REFERENCES users(id) ON DELETE SET NULL,
    box_id UUID REFERENCES users(id) ON DELETE SET NULL,
    pattern_id UUID REFERENCES review_patterns(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    detail TEXT,
    learned_date DATE NOT NULL,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)