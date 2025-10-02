-- Step 1: Remove the constraints added to the 'token' column.
ALTER TABLE "session"
DROP CONSTRAINT "session_token_key";

ALTER TABLE "session"
ALTER COLUMN "token" DROP NOT NULL;

-- Step 2: Drop the 'id' column.
ALTER TABLE "session"
DROP COLUMN "id";

-- Step 3: Re-establish 'token' as the primary key.
ALTER TABLE "session"
ADD PRIMARY KEY ("token");
