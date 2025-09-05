package query

var CheckBrandEmail = `SELECT
  ba."refApplicationCustId",
  ba."refBrandName",
  ba."refMailId"
FROM
  brand."refBrandApplication" ba
WHERE
  ba."refMailId" = $1`

var CreateNewApplication = `WITH new_contact AS (
    INSERT INTO brand."refBrandContactDetails" ("refBrandEmail")
    VALUES ($2)
    RETURNING "refBrandContactId"
),
new_social AS (
    INSERT INTO brand."refSocialMedia" ("refWebsiteUrl")
    VALUES (DEFAULT)
    RETURNING "refSocialMediaId"
),
new_location AS (
    INSERT INTO brand."refBrandLocation" ("refAddress")
    VALUES (DEFAULT)
    RETURNING "refLocationId"
),
new_warehouse AS (
    INSERT INTO brand."refBrandWareHouse" ("refIfWareHouse")
    VALUES (false)
    RETURNING "refWareHouseId"
),
new_document AS (
    INSERT INTO brand."refDocuments" ("refAddressProf")
    VALUES (DEFAULT)
    RETURNING "refDocumentsId"
)
INSERT INTO brand."refBrandApplication" (
    "refBrandName",
    "refMailId",
    "refApplicationCustId",
    "refApplicationStatus",
    "refSocialMediaId",
    "refDocumentsId",
    "refBrandLocationId",
    "refBrandContactId",
    "refWareHoueId",
    "refCreateAT"
)
VALUES (
    $1,
    $2,
    'BRND-' || TO_CHAR(CURRENT_DATE, 'YYYY') || '-' ||
    LPAD(
        COALESCE(
            (
                SELECT RIGHT("refApplicationCustId", 6)::int + 1
                FROM brand."refBrandApplication"
                WHERE "refApplicationCustId" LIKE 'BRND-' || TO_CHAR(CURRENT_DATE, 'YYYY') || '-%'
                ORDER BY "refApplicationCustId" DESC
                LIMIT 1
            ),
            1
        )::text,
        6,
        '0'
    ),
    1,
    (SELECT "refSocialMediaId" FROM new_social),
    (SELECT "refDocumentsId" FROM new_document),
    (SELECT "refLocationId" FROM new_location),
    (SELECT "refBrandContactId" FROM new_contact),
    (SELECT "refWareHouseId" FROM new_warehouse),
    $3
)
RETURNING "refApplicationCustId";`
