-- Create "users" table
CREATE TABLE "public"."users" (
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" text NOT NULL,
  "email" text NOT NULL,
  "password_hash" text NOT NULL,
  "role" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("user_id"),
  CONSTRAINT "chk_users_role" CHECK (role = ANY (ARRAY['manager'::text, 'member'::text]))
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create "folders" table
CREATE TABLE "public"."folders" (
  "folder_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "owner_id" uuid NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("folder_id"),
  CONSTRAINT "fk_users_folders" FOREIGN KEY ("owner_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "folder_shares" table
CREATE TABLE "public"."folder_shares" (
  "folder_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "access" text NOT NULL,
  PRIMARY KEY ("folder_id", "user_id"),
  CONSTRAINT "fk_folders_folder_shares" FOREIGN KEY ("folder_id") REFERENCES "public"."folders" ("folder_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_folder_shares" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_folder_shares_access" CHECK (access = ANY (ARRAY['read'::text, 'write'::text]))
);
-- Create "notes" table
CREATE TABLE "public"."notes" (
  "note_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "title" text NOT NULL,
  "body" text NULL,
  "folder_id" uuid NULL,
  "owner_id" uuid NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("note_id"),
  CONSTRAINT "fk_folders_notes" FOREIGN KEY ("folder_id") REFERENCES "public"."folders" ("folder_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_notes" FOREIGN KEY ("owner_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "note_shares" table
CREATE TABLE "public"."note_shares" (
  "note_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "access" text NOT NULL,
  PRIMARY KEY ("note_id", "user_id"),
  CONSTRAINT "fk_notes_note_shares" FOREIGN KEY ("note_id") REFERENCES "public"."notes" ("note_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_note_shares" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_note_shares_access" CHECK (access = ANY (ARRAY['read'::text, 'write'::text]))
);
-- Create "teams" table
CREATE TABLE "public"."teams" (
  "team_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "team_name" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("team_id")
);
-- Create "team_managers" table
CREATE TABLE "public"."team_managers" (
  "team_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("team_id", "user_id"),
  CONSTRAINT "fk_teams_managers" FOREIGN KEY ("team_id") REFERENCES "public"."teams" ("team_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_team_manager" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "team_members" table
CREATE TABLE "public"."team_members" (
  "team_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("team_id", "user_id"),
  CONSTRAINT "fk_teams_members" FOREIGN KEY ("team_id") REFERENCES "public"."teams" ("team_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_team_memberships" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
