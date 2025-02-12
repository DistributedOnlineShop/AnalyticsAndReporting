CREATE TABLE "sales_reports" (
  "s_report_id" UUID PRIMARY KEY NOT NULL,
  "report_type" VARCHAR NOT NULL,
  "start_date" TIMESTAMP(0) NOT NULL,
  "end_date" TIMESTAMP(0) NOT NULL,
  "total_sales" NUMERIC(7,2),
  "total_orders" INT,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);

CREATE TABLE "user_reports" (
  "u_report_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "report_type" VARCHAR NOT NULL,
  "total_value" NUMERIC(7,2),
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);
