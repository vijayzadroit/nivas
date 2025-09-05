package brandRegistrationQuery

var GetBrandRegisterFormData = `SELECT
  ba."refBrandName",
  ba."refProductCatageoryId",
  bd."refLogo",
  ba."refBrandDesciption",
  bs."refWebsiteUrl",
  bs."refInstaUrl",
  bc."refBrandContactPerson",
  bc."refDesignation",
  bc."refBrandMobile",
  bc."refBrandEmail",
  bl."refAddress",
  bl."refCity",
  bl."refZipCode",
  bl."refState",
  bd."refAddressProf",
  ba."refGstin",
  ba."refCin",
  bd."refGstin" AS "refGstinPath",
  bd."refPanCars" AS "refPanCars",
  bw."refIfWareHouse" AS "refIfWareHouse",
  bw."refAddress" AS "wareHouseAddress",
  bw."refCity" AS "wareHouseCity",
  bw."refDistrict" AS "wareHouseDistrict",
  bw."refZipcode" AS "wareHouseZipCode",
  bw."refState" AS "wareHouseState",
  ba."refSaveDraft",
  aps."refApplicationStatusId",
  aps."refApplicationStatus",
  bd."refAddressProfName",
  bd."refPanCarsName",
  bd."refGstinName",
  bd."refLogoName"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refBrandContactDetails" bc ON CAST(ba."refBrandContactId" AS INTEGER) = bc."refBrandContactId"
  LEFT JOIN brand."refBrandLocation" bl ON CAST(bl."refLocationId" AS INTEGER) = ba."refBrandLocationId"
  LEFT JOIN brand."refDocuments" bd ON CAST(bd."refDocumentsId" AS INTEGER) = ba."refDocumentsId"
  LEFT JOIN brand."refBrandWareHouse" bw ON CAST(bw."refWareHouseId" AS INTEGER) = ba."refWareHoueId"
  LEFT JOIN brand."refSocialMedia" bs ON CAST(bs."refSocialMediaId" AS INTEGER) = ba."refSocialMediaId"
  LEFT JOIN brand."refApplicationStatus" aps ON CAST(aps."refApplicationStatusId" AS INTEGER) = ba."refApplicationStatus"
WHERE
  ba."refApplicationCustId" = $1`
