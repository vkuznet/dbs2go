SELECT PE.PROCESSING_VERSION, PE.CREATION_DATE, PE.CREATE_BY, PE.DESCRIPTION
FROM {{.Owner}}.PROCESSING_ERAS PE
