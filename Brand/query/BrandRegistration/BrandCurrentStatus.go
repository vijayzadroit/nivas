package brandRegistrationQuery

var GetBrandCurrentStatus = `SELECT
  ba."refApplicationStatus"
FROM
  brand."refBrandApplication" ba
WHERE
  ba."refApplicationCustId" = $1`
