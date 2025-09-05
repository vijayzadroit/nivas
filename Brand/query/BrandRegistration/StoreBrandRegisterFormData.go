package brandRegistrationQuery

var GetReferenceId = ``

var StoreBrandInfoQuery = `UPDATE brand."refBrandApplication"
SET
  "refBrandName" = $1,
  "refProductCatageoryId" = $2,
  "refBrandDesciption" = $3,
  "refApplicationStatus" = $4,
  "refGstin" = $5,
  "refCin" = $6,
  "refUpdateAt" = $7
WHERE
  "refApplicationCustId" = $8
RETURNING
  "refSocialMediaId",
  "refDocumentsId",
  "refBrandLocationId",
  "refBrandContactId",
  "refWareHoueId"
`

var StoreBrandContactQuery = `UPDATE
  "brand"."refBrandContactDetails"
SET
  "refBrandContactPerson" = $2,
  "refBrandEmail" = $3,
  "refBrandMobile" = $4,
  "refDesignation" = $5
WHERE
  "refBrandContactId" = $1`

var StoreLocationQuery = `UPDATE
  "brand"."refBrandLocation"
SET
  "refAddress" = $2,
  "refCity" = $3,
  "refState" = $4,
  "refZipCode" = $5
WHERE
  "refLocationId" = $1`

var StoreWareHouseQuery = `  UPDATE
    "brand"."refBrandWareHouse"
  SET
    "refAddress" = $2,
    "refCity" = $3,
    "refDistrict" = $4,
    "refIfWareHouse" = $5,
    "refState" = $6,
    "refZipcode" = $7
  WHERE
    "refWareHouseId" = $1`

var StoreDocumentQuery = `UPDATE
  "brand"."refDocuments"
SET
  "refAddressProf" = $2,
  "refGstin" = $3,
  "refLogo" = $4,
  "refPanCars" = $5,
  "refAddressProfName" = $6,
  "refPanCarsName" = $7,
  "refGstinName" = $8,
  "refLogoName" = $9
WHERE
  "refDocumentsId" = $1`

var StoreSocialMedia = `UPDATE
  "brand"."refSocialMedia"
SET
  "refInstaUrl" = $2,
  "refWebsiteUrl" = $3
WHERE
  "refSocialMediaId" = $1`
