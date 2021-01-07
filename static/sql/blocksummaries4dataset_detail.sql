with t1 as(
     SELECT
         BS.BLOCK_NAME as BLOCK_NAME,
         NVL(SUM(FS.EVENT_COUNT),0) as NUM_EVENT
     FROM
         {(.Owner}}.FILES FS
     {{.BlockJoin}}
     {{.DatasetJoin}}
     {{.WhereClause}}
     group by BS.BLOCK_NAME
 )
 select
     b.block_name as block_name,
     b.file_count as num_file,
     b.block_size as file_size,
     t1.num_event as num_event,
     b.open_for_writing as open_for_writing
from
     {{.Owner}}.blocks b, t1
where
     t1.block_name = b.block_name