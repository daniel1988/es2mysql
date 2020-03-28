ROOT=$(cd "$(dirname "$0")"; cd ../; pwd)
LOG=${ROOT}/Log
PROC=${ROOT}/DcCenter/DcCenter
FLAG='
-dc_center=127.0.0.1:9999
-es_id_key=EsIncrIdKey
-eshost=127.0.0.1
-esindex=dc_index
-esport=9200
-log=/data/github/dc/Log
'
NRPROC=`ps ax | grep -v grep | grep -w $PROC | grep -w "$FLAG" | wc -l`
if [ $NRPROC -lt 1 ]
then
echo $(date +%Y-%m-%d) $(date +%H:%M:%S) $PROC >> $LOG/restart.log
$PROC $FLAG >> $LOG/$(basename $PROC).stderr 2>&1 &
fi
