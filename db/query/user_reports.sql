-- name: CreateUserReport :one
INSERT INTO user_reports (
    u_report_id,
    user_id,
    report_type,
    total_value
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetUserReportById :one
SELECT 
    * 
FROM 
    user_reports 
WHERE u_report_id = $1;

-- name: GetUserReportByUserId :many
SELECT 
    * 
FROM 
    user_reports 
WHERE 
    user_id = $1;

-- name: GetUserReportByOverTotalValue :many
SELECT 
    * 
FROM 
    user_reports 
WHERE 
    total_value >= $1;