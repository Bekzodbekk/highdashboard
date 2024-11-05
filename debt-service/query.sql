-- name: CreateDebt :one
INSERT INTO debts 
    (
        first_name, 
        last_name, 
        phone_number, 
        jshshir, 
        address, 
        price,
        price_paid,
        acquaintance, 
        collateral, 
        deadline
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING
    id,
    first_name, 
    last_name, 
    phone_number, 
    jshshir, 
    address, 
    price,
    price_paid,
    acquaintance, 
    collateral, 
    deadline,
    created_at,
    updated_at,
    deleted_at;


-- name: UpdateDebt :one
UPDATE debts
SET first_name=$2,
    last_name=$3, 
    phone_number=$4, 
    jshshir=$5, 
    address=$6,
    price = $7,
    price_paid = $8,
    acquaintance=$9, 
    collateral=$10, 
    deadline=$11,
    updated_at=$12
WHERE id = $1
RETURNING
    id,
    first_name, 
    last_name, 
    phone_number, 
    jshshir, 
    address,
    price,
    price_paid,
    acquaintance, 
    collateral, 
    deadline,
    created_at,
    updated_at,
    deleted_at;

-- name: DeleteDebt :exec
UPDATE debts
SET deleted_at = $2
WHERE id = $1;

-- name: GetDebtById :one
SELECT * FROM debts
WHERE id = $1;


-- name: GetDebtByFilter :many
SELECT * FROM debts
WHERE 
    ($1::text IS NULL OR 
    first_name ILIKE '%' || $1 || '%' OR
    last_name ILIKE '%' || $1 || '%' OR
    phone_number ILIKE '%' || $1 || '%' OR
    jshshir ILIKE '%' || $1 || '%' OR
    address ILIKE '%' || $1 || '%' OR
    price ILIKE '%' || $1 || '%' OR
    price_paid ILIKE '%' || $1 || '%' OR
    acquaintance ILIKE '%' || $1 || '%' OR
    collateral ILIKE '%' || $1 || '%' OR
    CAST(deadline AS text) ILIKE '%' || $1 || '%');