#!/usr/bin/env bash

# SETUP
PATH_TO_SOURCE_REPOS=$1 
# ensure the following repos are checked out to the proper release/branch for doc generation and
# ensure each is a subdirectory of $PATH_TO_SOURCE_REPOS 
# git clone https://github.com/algorandfoundation/algokit-cli
# git clone https://github.com/algorandfoundation/algokit-utils-ts
# git clone https://github.com/algorandfoundation/algokit-utils-py
# git clone https://github.com/algorandfoundation/algokit-client-generator-py
# git clone https://github.com/algorandfoundation/algokit-client-generator-ts

# REMOVE EXISTING ALGOKIT-LIBRARIES CONTENT
PATH_TO_LIBRARIES=../docs/get-details/algokit/libraries
rm -rf $PATH_TO_LIBRARIES

# RENDER ALGOKIT-LIBRARIES
mkdir -p $PATH_TO_LIBRARIES/py/algokit-utils
mkdir -p $PATH_TO_LIBRARIES/ts/algokit-utils
cp -R $PATH_TO_SOURCE_REPOS/algokit-utils-ts/docs/. $PATH_TO_LIBRARIES/ts/algokit-utils
mv $PATH_TO_LIBRARIES/ts/algokit-utils/README.md $PATH_TO_LIBRARIES/ts/algokit-utils/index.md
sed -i '' "s,# AlgoKit TypeScript Utilities,# Overview," $PATH_TO_LIBRARIES/ts/algokit-utils/index.md

cp -R $PATH_TO_SOURCE_REPOS/algokit-utils-py/docs/source/. $PATH_TO_LIBRARIES/py/algokit-utils
sed -i '' "s,# AlgoKit Python Utilities,# Overview," $PATH_TO_LIBRARIES/py/algokit-utils/index.md

mkdir -p $PATH_TO_LIBRARIES/ts/client-generator
cp -R $PATH_TO_SOURCE_REPOS/algokit-client-generator-ts/README.md $PATH_TO_LIBRARIES/ts/client-generator.md
sed -i '' "s,# AlgoKit TypeScript client generator (algokit-client-generator-ts),# algokit-client-generator," $PATH_TO_LIBRARIES/ts/client-generator.md

mkdir -p $PATH_TO_LIBRARIES/py/client-generator
cp -R $PATH_TO_SOURCE_REPOS/algokit-client-generator-py/README.md $PATH_TO_LIBRARIES/py/client-generator.md
sed -i '' "s,# AlgoKit Python client generator,# algokit-client-generator," $PATH_TO_LIBRARIES/py/client-generator.md

# ADD .pages 
echo "title: AlgoKit\n\narrange:\n- index.md\n- features\n- libraries\n- cli-reference.md\n- tutorials\n- architecture-decisions\n- articles" > ../docs/get-details/algokit/.pages
echo "title: Libraries\n\narrange:\n- ts\n- py" > ../docs/get-details/algokit/libraries/.pages
echo "title: for TypeScript\n\n" > $PATH_TO_LIBRARIES/ts/.pages
echo "title: for Python\n\\n" > $PATH_TO_LIBRARIES/py/.pages
echo "title: algokit-utils\n\n" > $PATH_TO_LIBRARIES/ts/algokit-utils/.pages
echo "title: algokit-utils\n\\n" > $PATH_TO_LIBRARIES/py/algokit-utils/.pages

#echo "title: algokit-client-generator-ts\n\narrange:\n- index.md" > $PATH_TO_LIBRARIES/client-generator-ts/.pages
#echo "title: algokit-client-generator-py\n\narrange:\n- index.md" > $PATH_TO_LIBRARIES/client-generator-py/.pages


# CLEANUP
rm -rf $PATH_TO_LIBRARIES/ts/algokit-utils/images
rm -rf $PATH_TO_LIBRARIES/py/algokit-utils/conf.py