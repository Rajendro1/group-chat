package db
var TABLE_CREATION = 
`
CREATE TABLE IF NOT EXISTS public.users (
	id bigserial PRIMARY KEY,
	username varchar NOT NULL,
);

CREATE TABLE IF NOT EXISTS public.messages (
    id bigserial PRIMARY KEY,
    user_id bigserial NOT NULL,
    content VARCHAR,
    created_at TimeStamp TimeNow()
)

`

