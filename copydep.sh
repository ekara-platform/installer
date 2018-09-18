echo "INSTALLER_IMAGE copydep..."

# Refresh "model"
rm -rf ./go/vendor/github.com/lagoon-platform/model/*.go
cp ../model/*.go  ./go/vendor/github.com/lagoon-platform/model/

# Refresh "engine"
rm -rf ./go/vendor/github.com/lagoon-platform/engine/*.go
cp ../engine/*.go  ./go/vendor/github.com/lagoon-platform/engine/

rm -rf ./go/vendor/github.com/lagoon-platform/engine/ansible/*.go
cp ../engine/ansible/*.go  ./go/vendor/github.com/lagoon-platform/engine/ansible/

rm -rf ./go/vendor/github.com/lagoon-platform/engine/ssh/*.go
cp ../engine/ssh/*.go  ./go/vendor/github.com/lagoon-platform/engine/ssh/

# Refresh "installer"
rm -rf ./go/vendor/github.com/lagoon-platform/installer/*.go
cp ../installer/*.go  ./go/vendor/github.com/lagoon-platform/installer/

cd ./go
go build