-- Drop the foreign key constraint added to the "accounts" table
ALTER TABLE "accounts" DROP CONSTRAINT "accounts_owner_fkey";

-- Drop the unique index on "accounts" ("owner", "currency")
DROP INDEX IF EXISTS "accounts_owner_currency_idx";

-- Drop the "users" table
DROP TABLE IF EXISTS "users";
