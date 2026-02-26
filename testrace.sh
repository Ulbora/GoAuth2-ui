cd handlers
go test -race 
sleep 15
cd ..
cd services
go test -race 
sleep 15
cd ..
