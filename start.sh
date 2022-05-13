./stop.sh

cd language-service/
go run main.go >> ../all_logs.log 2>&1 &
cd ../
go run cmd/*.go >> all_logs.log 2>&1 &