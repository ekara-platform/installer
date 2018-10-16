echo "INSTALLER_IMAGE copydep..."

# Refresh "model"
rm -rf ./go/vendor/github.com/ekara-platform/model/*.go
cp ../model/*.go  ./go/vendor/github.com/ekara-platform/model/

# Refresh "engine"
rm -rf ./go/vendor/github.com/ekara-platform/engine/*.go
cp ../engine/*.go  ./go/vendor/github.com/ekara-platform/engine/

rm -rf ./go/vendor/github.com/ekara-platform/engine/ansible/*.go
cp ../engine/ansible/*.go  ./go/vendor/github.com/ekara-platform/engine/ansible/

rm -rf ./go/vendor/github.com/ekara-platform/engine/ssh/*.go
cp ../engine/ssh/*.go  ./go/vendor/github.com/ekara-platform/engine/ssh/

rm -rf ./go/vendor/github.com/ekara-platform/engine/component/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/component/
cp ../engine/component/*.go  ./go/vendor/github.com/ekara-platform/engine/component/


rm -rf ./go/vendor/github.com/ekara-platform/engine/util/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/util/
cp ../engine/util/*.go  ./go/vendor/github.com/ekara-platform/engine/util/


# Refresh "installer"
rm -rf ./go/vendor/github.com/ekara-platform/installer/*.go
cp ../installer/*.go  ./go/vendor/github.com/ekara-platform/installer/


rm -rf ./vendor/github.com/ekara-platform/engine/component/*.go
mkdir ./vendor/github.com/ekara-platform/engine/component/
cp ../engine/component/*.go  ./vendor/github.com/ekara-platform/engine/component/


rm -rf ./vendor/github.com/ekara-platform/engine/util/*.go
mkdir ./vendor/github.com/ekara-platform/engine/util/
cp ../engine/util/*.go  ./vendor/github.com/ekara-platform/engine/util/

cd ./go
go build