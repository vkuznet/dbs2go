{{if .RunNumber}}
SELECT DISTINCT
{{else}}
SELECT 
{{end}}
        F.FILE_ID, F.LOGICAL_FILE_NAME, F.IS_FILE_VALID,
        F.DATASET_ID, D.DATASET,
        F.BLOCK_ID, B.BLOCK_NAME,
        F.FILE_TYPE_ID, FT.FILE_TYPE,
        F.CHECK_SUM, F.EVENT_COUNT, F.FILE_SIZE,
        F.BRANCH_HASH_ID, F.ADLER32, F.MD5,
        F.AUTO_CROSS_SECTION,
        F.CREATION_DATE, F.CREATE_BY,
        F.LAST_MODIFICATION_DATE, F.LAST_MODIFIED_BY
{{if .RunNumber}}
        , FL.RUN_NUM
{{end}}
FROM {{.Owner}}.FILES F
JOIN {{.Owner}}.FILE_DATA_TYPES FT ON  FT.FILE_TYPE_ID = F.FILE_TYPE_ID
JOIN {{.Owner}}.DATASETS D ON  D.DATASET_ID = F.DATASET_ID
JOIN {{.Owner}}.BLOCKS B ON B.BLOCK_ID = F.BLOCK_ID
JOIN {{.Owner}}.DATASET_ACCESS_TYPES DT ON  DT.DATASET_ACCESS_TYPE_ID = D.DATASET_ACCESS_TYPE_ID
{{if .RunNumber}}
JOIN {{.Owner}}.FILE_LUMIS FL on FL.FILE_ID=F.FILE_ID
{{end}}
{{if .Addition}}
LEFT OUTER JOIN {{.Owner}}.FILE_OUTPUT_MOD_CONFIGS FOMC ON FOMC.FILE_ID = F.FILE_ID
LEFT OUTER JOIN {{.Owner}}.OUTPUT_MODULE_CONFIGS OMC ON OMC.OUTPUT_MOD_CONFIG_ID = FOMC.OUTPUT_MOD_CONFIG_ID
LEFT OUTER JOIN {{.Owner}}.RELEASE_VERSIONS RV ON RV.RELEASE_VERSION_ID = OMC.RELEASE_VERSION_ID
LEFT OUTER JOIN {{.Owner}}.PARAMETER_SET_HASHES PSH ON PSH.PARAMETER_SET_HASH_ID = OMC.PARAMETER_SET_HASH_ID
LEFT OUTER JOIN {{.Owner}}.APPLICATION_EXECUTABLES AEX ON AEX.APP_EXEC_ID = OMC.APP_EXEC_ID
{{end}}
