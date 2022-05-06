# Copyright 2022 Aleksandr Soloshenko
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TAG=$(git tag)
if [ -z "$TAG" ]; then
    exit 0
fi

if [ -z "$APP" ]; then
    echo "Need APP environment var"
    exit 1
fi

DOCKER_IMAGE="capcom.azurecr.io/route256-${APP}"

PARTS=$(echo ${TAG:1} | tr "." "\n")
VERSION=

docker build -f build/package/Dockerfile.${APP} --target default -t ${DOCKER_IMAGE}:latest .
docker push ${DOCKER_IMAGE}:latest

for V in $PARTS
do
    if [ -n "$VERSION" ];
    then
        VERSION=${VERSION}.
    fi
    VERSION=${VERSION}${V}
    
    docker tag ${DOCKER_IMAGE}:latest ${DOCKER_IMAGE}:${VERSION}
    docker push ${DOCKER_IMAGE}:${VERSION}
done
