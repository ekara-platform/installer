echo "INSTALLER_IMAGE copydep..."

# Refresh "model"
rm -rf ./go/vendor/github.com/ekara-platform/model/*.go
mkdir ./go/vendor/github.com/ekara-platform/model/
cp ../model/*.go  ./go/vendor/github.com/ekara-platform/model/

# Refresh "engine"
rm -rf ./go/vendor/github.com/ekara-platform/engine/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/
cp ../engine/*.go  ./go/vendor/github.com/ekara-platform/engine/

rm -rf ./go/vendor/github.com/ekara-platform/engine/ansible/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/ansible/
cp ../engine/ansible/*.go  ./go/vendor/github.com/ekara-platform/engine/ansible/

rm -rf ./go/vendor/github.com/ekara-platform/engine/ssh/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/ssh/
cp ../engine/ssh/*.go  ./go/vendor/github.com/ekara-platform/engine/ssh/


rm -rf ./go/vendor/github.com/ekara-platform/engine/component/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/component/
cp ../engine/component/*.go  ./go/vendor/github.com/ekara-platform/engine/component/

rm -rf ./go/vendor/github.com/ekara-platform/engine/util/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/util/
cp ../engine/util/*.go  ./go/vendor/github.com/ekara-platform/engine/util/

rm -rf ./go/vendor/github.com/ekara-platform/engine/component/scm/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/component/scm/
cp ../engine/component/scm/*.go  ./go/vendor/github.com/ekara-platform/engine/component/scm/

rm -rf ./go/vendor/github.com/ekara-platform/engine/component/scm/file/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/component/scm/file/
cp ../engine/component/scm/file/*.go  ./go/vendor/github.com/ekara-platform/engine/component/scm/file/

rm -rf ./go/vendor/github.com/ekara-platform/engine/component/scm/git/*.go
mkdir ./go/vendor/github.com/ekara-platform/engine/component/scm/git/
cp ../engine/component/scm/git/*.go  ./go/vendor/github.com/ekara-platform/engine/component/scm/git/



# Refresh "installer"
mkdir ./go/vendor/github.com/ekara-platform/installer/
rm -rf ./go/vendor/github.com/ekara-platform/installer/*.go
cp ../installer/*.go  ./go/vendor/github.com/ekara-platform/installer/
