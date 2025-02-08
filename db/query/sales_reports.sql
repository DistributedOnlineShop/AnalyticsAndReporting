-- name: CreateSalesReports :one
INSERT INTO sales_reports (
    report_type,
    start_date,
    end_date,
    total_sales,
    total_orders
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetSalesReportsList :many
SELECT 
    * 
FROM 
    sales_reports
ORDER BY 
    created_at DESC;

-- name: GetSalesReportsByDate :many
SELECT 
    * 
FROM 
    sales_reports
WHERE 
    start_date <= $1 AND end_date >= $1;

-- name: GetSalesReportsById :one
SELECT 
    * 
FROM 
    sales_reports
WHERE 
   s_report_id = $1;
