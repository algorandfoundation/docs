#!/usr/bin/env python3

from git import Repo, TagReference
from shutil import rmtree, copytree, move, copy
from os import path, walk, rename
import re

###
# initialize
PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS = '../docs/get-details/algokit/utils/'
PATH_TO_TMP_DIR_FOR_SOURCE_REPOS = '/tmp/source_repos/'
rmtree(PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS, ignore_errors=True)
rmtree(PATH_TO_TMP_DIR_FOR_SOURCE_REPOS, ignore_errors=True)
list_of_repos = [
    'https://github.com/algorandfoundation/algokit-utils-ts',
    'https://github.com/algorandfoundation/algokit-utils-py',
    'https://github.com/algorandfoundation/algokit-client-generator-ts',
    'https://github.com/algorandfoundation/algokit-client-generator-py'
]

###
# Clones the repo to local directory 
def clone_and_checkout_latest(repo_url) :
    repo_name = repo_url.split('/')[-1]
    local_dir = PATH_TO_TMP_DIR_FOR_SOURCE_REPOS + repo_name
    repo = Repo.clone_from(repo_url, local_dir) 
    assert not repo.bare
    print("cloned", repo_name)

    ### 
    # Checkout the latest release tag
    tags = repo.tags
    ref_path = None
    i = len(tags) - 1
    while i !=0 :
        if "-" not in tags[i].path :
            ref_path = tags[i].path
            break
        i -= 1
    release_tag = repo.create_head(ref_path, "HEAD")
    repo.head.reference = release_tag
    assert not repo.head.is_detached
    repo.head.reset(index=True, working_tree=True)
    print("  checked out", release_tag)

###
# Function to serach/replace text in file
def search_replace_text(search_text, replace_text, path_to_file): 
    with open(path_to_file,'r+') as f: 
        file = f.read() 
        file = re.sub(search_text, replace_text, file)
        f.seek(0) 
        f.write(file) 
        f.truncate() 

###
# Function to strip the file extention
def strip_file_extentions_recursively(root_dir, exts) :
    for root, dirs, files in walk(root_dir):
        for currentFile in files:
            if currentFile.lower().endswith(exts):
                rename(path.join(root, currentFile), path.join(root, currentFile).rstrip(exts))

###
# Function to replace string in all files within a directory
def replace_string_recursively(search_text, replace_text, root_dir) :
    for root, dirs, files in walk(root_dir):
        for currentFile in files:
            if path.join(root, currentFile).lower().endswith(".md") : 
                search_replace_text(search_text, replace_text, path.join(root, currentFile))

###
# copy algokit-utils-ts docs from source into mkdocs
def fixup_algokit_utils_ts(repo_url) :
    repo_name = repo_url.split('/')[-1]
    src = PATH_TO_TMP_DIR_FOR_SOURCE_REPOS + repo_name + '/docs'
    dst = PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS + repo_name.split('-')[-1]
    copytree(src, dst)
    move(dst + '/README.md' , dst + '/index.md' )
    move(dst + '/code/README.md' , dst + '/code/index.md' )
    replace_string_recursively("README", "index", dst)
    replace_string_recursively("\.\./\.\./src", "https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src", dst)
    search_replace_text("# AlgoKit TypeScript Utilities" , "title: Overview", dst + '/index.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli" , "../../index.md", dst + '/index.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-utils-py" , "../py", dst + '/index.md')
    search_replace_text(" \| \[Reference docs\]\(#reference-documentation\)" , "", dst + '/index.md')
    search_replace_text("(?ms)# Reference documentation.*", "", dst + '/index.md')
    search_replace_text("Note: ", "!!! Note\n\t", dst + '/index.md')
    search_replace_text("\- Note: ", "\n!!! Note\n    ", dst + '/capabilities/account.md')
    search_replace_text("\' note: ", "\n!!! Note\n    ", dst + '/capabilities/account.md')
    search_replace_text("\(note: y", "\n!!! Note\n    Y", dst + '/capabilities/amount.md')
    search_replace_text("Note: t", "!!! Note\n    T", dst + '/capabilities/amount.md')
    search_replace_text(":\n\n```", ":\n    ```", dst + '/capabilities/amount.md')
    search_replace_text("import \{", "    import {", dst + '/capabilities/amount.md')
    search_replace_text("```\n", "    ```\n", dst + '/capabilities/amount.md')
    search_replace_text("> ⚠️ \*\*Note:\*\* ", "!!! Note\n    ", dst + '/capabilities/app-client.md')
    search_replace_text("Note: If", "If", dst + '/capabilities/app-client.md')
    search_replace_text("> \*\*Note\*\*:", "!!! Note \n    ", dst + '/capabilities/asset.md')
    search_replace_text("> Please note that t", "!!! Note \n    T", dst + '/capabilities/transfer.md')
    search_replace_text("(?ms)@.*contents", "title: Overview", dst + '/code/index.md')
    search_replace_text("(?ms)\[@.*contents", "title: Overview", dst + '/code/modules/index.md')
    open(dst + '/.pages', 'a').write('title: TypeScript\n\narrange:\n- index.md\n- capabilities\n- code')
    open(dst + '/code/.pages', 'a').write('title: API Ref') 

###
# copy algokit-utils-py docs from source into mkdocs
def fixup_algokit_utils_py(repo_url) :
    repo_name = repo_url.split('/')[-1]
    src = PATH_TO_TMP_DIR_FOR_SOURCE_REPOS + repo_name + '/docs/html/_sources'
    dst = PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS + repo_name.split('-')[-1]
    copytree(src, dst)
    strip_file_extentions_recursively(dst, '.txt')
    move(dst + '/apidocs/algokit_utils/algokit_utils.md', dst + '/apidocs.md')
    rmtree(dst + '/apidocs', ignore_errors=True)
    search_replace_text("# AlgoKit Python Utilities" , r"title: Overview", dst + '/index.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-utils-ts" , "../ts/index.md", dst + '/index.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli" , "../../index.md", dst + '/index.md')
    search_replace_text(" \| \[Reference docs\]\(#reference-documentation\)" , "", dst + '/index.md')
    search_replace_text("(?ms)# Reference documentation.*", "", dst + '/index.md')
    search_replace_text("\(.*\)\=\n", "", dst + '/index.md')
    search_replace_text("(?ms)```\{toctree.*?```\n\n\n", "", dst + '/index.md')
    search_replace_text("(?ms)```\{note.*?^If", "!!! Note\n\tIf", dst + '/index.md')
    search_replace_text("(?ms)\)\.\n```", ").", dst + '/index.md')
    search_replace_text("/algokit_utils/algokit_utils", "", dst + '/index.md')
    search_replace_text("(?ms)---.*---", "title: API Ref", dst + '/apidocs.md')
    search_replace_text("(?ms)# \{.*## Package Contents", "", dst + '/apidocs.md')
    search_replace_text("\(.*\)\=\n", "", dst + '/capabilities/account.md')
    search_replace_text("exposed:", "exposed:\n", dst + '/capabilities/account.md')
    search_replace_text("   *", "    ", dst + '/capabilities/app-client.md')
    search_replace_text("(?ms)```\{note\}", "!!! Note", dst + '/capabilities/app-client.md')
    search_replace_text("\"\"\"", "", dst + '/capabilities/app-client.md')
    search_replace_text("(?ms)LogicError`\.\n```", ".", dst + '/capabilities/app-client.md')
    search_replace_text("The extended", "\tThe extended", dst + '/capabilities/app-client.md')
    search_replace_text("Remember", "\tRemember", dst + '/capabilities/app-client.md')
    search_replace_text("map`\n```", "map`", dst + '/capabilities/app-client.md')
    search_replace_text("[1-4]\.\) ", "\t- ", dst + '/capabilities/app-client.md')
    search_replace_text("\(.*\)\=\n", "", dst + '/capabilities/app-deploy.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli/blob/main/docs/articles/output_stability.md", "../../../../articles/output_stability", dst + '/capabilities/app-deploy.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli/blob/main/docs/architecture-decisions/2023-01-12_smart-contract-deployment.md", "../../../../architecture-decisions/2023-01-12_smart-contract-deployment", dst + '/capabilities/app-deploy.md')
    search_replace_text("images/", "../../architecture-decisions/", dst + '/capabilities/app-deploy.md')
    search_replace_text("\* \`algokit_utils\.get_purestake.*TestNet\n", "", dst + '/capabilities/client.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md", "../../../../features/localnet", dst + '/capabilities/client.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-utils-py/blob/main/docs/source/capabilities/dispenser-client.md", "dispenser-client.md", dst + '/capabilities/transfer.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/dispenser.md", "../../../../features/dispenser", dst + '/capabilities/transfer.md')
    search_replace_text("> Please note, i", "!!! Note\n\tI", dst + '/capabilities/transfer.md')
    search_replace_text("https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md", "../../../../features/localnet", dst + '/capabilities/transfer.md')
    open(dst + '/.pages', 'a').write('title: Python\n\narrange:\n- index.md\n- capabilities\n- apidocs.md')
    
def fixup_algokit_client_generator_ts(repo_url) :
    repo_name = repo_url.split('/')[-1]
    src = PATH_TO_TMP_DIR_FOR_SOURCE_REPOS + repo_name
    dst = PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS + repo_name.split('-')[-1]
    copy(src + '/README.md', dst + '/client_generator.md')
    search_replace_text("\(./examples\)", "(https://github.com/algorandfoundation/algokit-client-generator-ts/tree/main/examples)", dst + '/client_generator.md')
    open(dst + '/.pages', 'a').write('title: Python\n\narrange:\n- index.md\n- capabilities\n- client_generator.md\n- code')
    

def fixup_algokit_client_generator_py(repo_url) :
    repo_name = repo_url.split('/')[-1]
    src = PATH_TO_TMP_DIR_FOR_SOURCE_REPOS + repo_name
    dst = PATH_TO_MKDOCS_CONTENT_FOR_ALGOKIT_UTILS + repo_name.split('-')[-1]
    copy(src + '/README.md', dst + '/client_generator.md')
    search_replace_text("\(./examples\)", "(https://github.com/algorandfoundation/algokit-client-generator-py/tree/main/examples)", dst + '/client_generator.md')
    open(dst + '/.pages', 'a').write('title: TypeScript\n\narrange:\n- index.md\n- capabilities\n- client_generator.md\n- apidocs.md')

def main() :
    for repo_url in list_of_repos :
        clone_and_checkout_latest(repo_url)
    fixup_algokit_utils_ts(list_of_repos[0])
    fixup_algokit_utils_py(list_of_repos[1])
    fixup_algokit_client_generator_ts(list_of_repos[2])
    fixup_algokit_client_generator_py(list_of_repos[3])

main()

