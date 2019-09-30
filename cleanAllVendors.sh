cd ../engine
./cleanVendor.sh

cd ../installer
./cleanVendor.sh

echo "Compilation of the image sources..."
go build



