package brandRegistrationQuery

var GetBrandRegistrationStatusData = `SELECT
  aps."refApplicationStatus",
  aps."refApplicationStatusId",
  ba."refApplicationCustId",
  '3-5 Business Days' AS "refProcessTime",
  bd."refLogo",
  bd."refAddressProf",
  bd."refPanCars" AS "refPanDoc",
  bd."refGstin" AS "refGstinDoc",
  ba."refBrandName",
  bc."refBrandCategoryName",
  ba."refCin",
  ba."refGstin",
  cd."refBrandContactPerson",
  cd."refBrandMobile",
  cd."refBrandEmail",
  ba."refCreateAT",
  ba."refUpdateAt",
  aps."refApplicationReMark" AS "currentStatus",
  bd."refAddressProfName",
  bd."refPanCarsName",
  bd."refGstinName",
  bd."refLogoName",
  CASE
    WHEN ba."refApplicationStatus" = 5 THEN ba."refApplicationStatusDesciption"
    ELSE aps."refApplicationStatusDesciption"
  END AS "refApplicationStatusDesciption"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refApplicationStatus" aps ON CAST(aps."refApplicationStatusId" AS INTEGER) = ba."refApplicationStatus"
  LEFT JOIN brand."refDocuments" bd ON CAST(bd."refDocumentsId" AS INTEGER) = ba."refDocumentsId"
  LEFT JOIN brand."refBrandCategory" bc ON CAST(bc."refBrandCategoryId" AS INTEGER) = ba."refProductCatageoryId"
  LEFT JOIN brand."refBrandContactDetails" cd ON CAST(cd."refBrandContactId" AS INTEGER) = ba."refBrandContactId"
WHERE
  ba."refApplicationCustId" = $1`
