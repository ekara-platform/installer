cd ../model
go build

# copydep in  "engine"
cd ../engine
./copydep.sh
go build

# copydep in  "installer"
cd ../installer
./copydep.sh
go build

# copydep in  "installer_image"
cd ../installer_image
./copydep.sh
cd ./go
go build
