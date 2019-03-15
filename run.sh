dep ensure
cd cmd/mibici/
go build -o ../../mibici
cd ../..

export ENVIRONMENT=LOCAL
./mibici -ConfigDir=config -ConfigType=yaml