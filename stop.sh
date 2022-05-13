LPID=`cat language-service/language_service.pid`
FPID=`cat filter_service.pid`

PS_OUT=$(ps -p "$LPID")
if [ "$?" -eq 0 ]; then
        LANG_IS_UP=1
fi

PS_OUT=$(ps -p "$FPID")
if [ "$?" -eq 0 ]; then
        FILTER_IS_UP=1
fi

if [ $LANG_IS_UP ]; then
    echo `date` "Language service is already running. Killing it"
    kill -TERM $LPID
fi

if [ $FILTER_IS_UP ]; then
    echo `date` "Filter service is already running. Killing it"
    kill -TERM $FPID
fi