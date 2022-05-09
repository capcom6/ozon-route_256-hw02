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

TAG=$(git describe --exact-match --tags $(git log -n1 --pretty='%h'))
if [ -z "$TAG" ]; then
    exit 0
fi

if [ -z "$APP" ]; then
    echo "Need APP environment var"
    exit 1
fi

DOCKER_IMAGE="gitlab-registry.ozon.dev/capcom6/homework-2/${APP}"

FULL_VERSION=${TAG:1}
PARTS=$(echo ${FULL_VERSION} | tr "." "\n")
VERSION=

docker build -f build/package/Dockerfile.${APP} --target default -t ${DOCKER_IMAGE}:${FULL_VERSION} .
docker push ${DOCKER_IMAGE}:${FULL_VERSION}

for V in $PARTS
do
    if [ -n "$VERSION" ];
    then
        VERSION=${VERSION}.
    fi
    VERSION=${VERSION}${V}

    if [ "$FULL_VERSION" = "$VERSION" ];
    then
        continue
    fi
    
    docker tag ${DOCKER_IMAGE}:${FULL_VERSION} ${DOCKER_IMAGE}:${VERSION}
    docker push ${DOCKER_IMAGE}:${VERSION}
done

docker tag ${DOCKER_IMAGE}:${FULL_VERSION} ${DOCKER_IMAGE}:latest
docker push ${DOCKER_IMAGE}:latest
