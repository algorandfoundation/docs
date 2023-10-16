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

# REMOVE EXISTING ALGOKIT CONTENT
rm -rd ../docs/get-details/algokit

# RENDER ALGOKIT-CLI
cp -R $PATH_TO_SOURCE_REPOS/algokit-cli/docs/ ../docs/get-details/algokit
mv ../docs/get-details/algokit/algokit.md ../docs/get-details/algokit/index.md
sed -i '' "1s/.*/title: Overview/" ../docs/get-details/algokit/index.md
sed -i '' "s,completions.md.*,&\n- [Dispenser](./features/dispenser.md) - interactive TestNet Dispenser for receiveing TestALGO," ../docs/get-details/algokit/index.md
sed -i '' "s/algokit.md/index.md/" ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,look that the,look at the," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,\[AlgoKit install\](https://github.com/algorandfoundation/algokit-cli/blob/main/README.md#install) guide,\[quick start videos\](#quick-start-videos) above," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,\[directly used by certain commands\](../../README.md#prerequisites),directly used by certain commands," ../docs/get-details/algokit/features/doctor.md
sed -i '' "s,Prequisites,Prerequisites," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,\[README\](../../README.md#prerequisites),\[install guide\](../tutorials/intro.md#prerequisites)," ../docs/get-details/algokit/features/init.md
sed -i '' "s,https://github.com/algorandfoundation/algokit-cli/blob/main/docs/algokit,../index," ../docs/get-details/algokit/tutorials/algokit-template.md
sed -i '' "s,../../docs/algokit.md,../index.md," ../docs/get-details/algokit/architecture-decisions/2022-11-14_sandbox-approach.md
sed -i '' "s,../../docs/algokit.md,../index.md," ../docs/get-details/algokit/architecture-decisions/2023-01-12_smart-contract-deployment.md
sed -i '' "s,../../docs/algokit.md,../index.md," ../docs/get-details/algokit/architecture-decisions/2023-06-06_frontend-templates.md
sed -i '' "s,../../docs/algokit.md,../index.md," ../docs/get-details/algokit/architecture-decisions/2023-07-19_advanced_generate_command.md
sed -i '' "s,quick start tutorial: \[https://bit.ly/algokit-intro-tutorial\](https://bit.ly/algokit-intro-tutorial),[quick start tutorial](tutorials/intro.md)," ../docs/get-details/algokit/cli/index.md
sed -i '' "s,https://github.com/algorandfoundation/algokit-cli/blob/main/docs/,../," ../docs/get-details/algokit/features/init.md
sed -i '' "/- \[algokit\](#algokit)/,/# algokit/d" ../docs/get-details/algokit/cli/index.md
sed -i '' "s,# AlgoKit CLI Reference Documentation,# CLI Reference," ../docs/get-details/algokit/cli/index.md
sed -i '' "/## Table of Contents/,/## Quick Start/d" ../docs/get-details/algokit/tutorials/algokit-template.md
sed -i '' "s,For users who are keen,## Quick Start\nFor users who are keen," ../docs/get-details/algokit/tutorials/algokit-template.md
sed -i '' "s,https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/bootstrap,../features/bootstrap," ../docs/get-details/algokit/tutorials/algokit-template.md
sed -i '' "s,https://github.com/algorandfoundation/algokit-cli/blob/deploy-command/docs/features/deploy,../features/deploy," ../docs/get-details/algokit/tutorials/algokit-template.md
sed -i '' "s,dow6U8DxOGc,MzBRef_Res8," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,https://docs.docker.com/desktop/install/mac-install/,https://docker.com/download/," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,\[VSCode\](https://code.visualstudio.com/download),[VSCode](https://code.visualstudio.com/download) (recommended)," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s,\[official installation guide](https://github.com/algorandfoundation/algokit-cli#install),[quick start videos](#quick-start-videos)," ../docs/get-details/algokit/tutorials/intro.md
sed -i '' "s, Feature Documentation,," ../docs/get-details/algokit/features/deploy.md
sed -i '' "s,\[permanently on a given machine\](./cli/index.md#version-prompt) with \`algokit config version-prompt disable\`,permanently on a given machine with [\`algokit config version-prompt disable\`](cli/#-skip-version-check)," ../docs/get-details/algokit/index.md

# REORG CLI REFERENCE
mv ../docs/get-details/algokit/cli/index.md ../docs/get-details/algokit/cli-reference.md
echo "title: AlgoKit\n\narrange:\n- index.md\n- features\n- cli-reference.md\n- tutorials\n- architecture-decisions\n- articles" > ../docs/get-details/algokit/.pages
cd ../docs/get-details/algokit/
find . -type f -name '*.md' -exec sed -i '' "s,cli/index.md,cli-reference.md," {} +
cd ../../../scripts/

# CLEANUP
rm ../docs/get-details/algokit/tutorials/smart-contracts.md # contains "TODO"
rm -rd ../docs/get-details/algokit/sphinx
rm -rd ../docs/get-details/algokit/cli
