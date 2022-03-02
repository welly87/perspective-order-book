# install pre requisite
https://github.com/weltam/bootstrap/blob/main/al2/al2.sh

sudo yum install python3-devel

# setup cmake

https://github.com/weltam/bootstrap/blob/main/setup-cmake.sh

# install boost
# https://gist.github.com/1duo/2d1d851f76f8297be264b52c1f31a2ab

wget https://boostorg.jfrog.io/artifactory/main/release/1.78.0/source/boost_1_78_0.tar.gz
tar -xzf boost_1_*
cd boost_1_*
./bootstrap.sh --prefix=/opt/boost
sudo ./b2 install --prefix=/opt/boost --with=all

# install flatbuffers

# https://gist.github.com/ankur-anand/389536ccd5accd50143bdacf605dc79b

# install perspective

python3 -m pip install perspective-python websocket-client



# install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash

# install nodejs
nvm install node

# install yarn

npm install --global yarn

# install pers
git clone https://github.com/sc1f/perspective-order-book.git

yarn

yarn run start:client

yarn run start:server
