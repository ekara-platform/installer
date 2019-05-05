./copyAllDep.sh

#!/bin/sh
docker run --rm -v "$PWD/go":/go/src/installer -w /go/src/installer iron/go:dev go build -o installer

echo http_proxy \$1:  $1
echo https_proxy \$2:  $2




if [ "$1" = "" ]
then
    docker build -t ekaraplatform/installer .     
else
	if [ "$2" = "" ]
	then
		echo "   using only http_proxy..."
		docker build --build-arg http_proxy="$1" --build-arg https_proxy="$1" -t ekaraplatform/installer .     
	else
		echo "   using http_proxy and https_proxy..."
		docker build --build-arg http_proxy="$1" --build-arg https_proxy="$2" -t ekaraplatform/installer .     
	fi		
fi

