alter table "public"."pet" add column "inventoried_at" timestamptz
 null default now();
