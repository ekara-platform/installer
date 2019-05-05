cd ../engine
./cleanVendor.sh

cd ../installer
./cleanVendor.sh

cd ../installer_image/go
./cleanVendor.sh
echo "Compilation of the image sources..."
go build



