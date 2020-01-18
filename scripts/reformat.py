#!/usr/bin/env python3
# rm -rf goal && mkdir goal && goal generate-docs goal && ./reformat.py -path goal

from os import walk
import argparse    
import os    
import os.path
import re
import tempfile
import shutil
import pprint
pp = pprint.PrettyPrinter(indent=4)

def fix_file_links(new_file, goal_depth):
    """ Update markdown links to use relative paths in different directories. """
    with tempfile.NamedTemporaryFile(mode='w', dir='.', delete=False) as tmp, \
            open(new_file, 'r') as f:
        for line in f:
            result = line
            m = re.search('(^\* \[.*\]\()(.*)(\).*$)', result)
            if m != None:
                link=""
                if m.group(2) == "goal.md":
                    # -1 because we'll move goal.md up a directory at the end.
                    link = ('../'*goal_depth) + 'goal.md'
                else:
                    link = '../' + m.group(2).split('_')[-1].split('.')[0] + '/'
                result = "%s%s%s" % (m.group(1), link, m.group(3))
            tmp.write(result + "\n")
        os.replace(tmp.name, new_file)
    
def process(dirpath):
    """ move files into a directory structure and add .pages files. """
    for f in os.listdir(dirpath):
        parts = f.split('_')
        # Ignore the root "goal.md" file
        new_name=parts.pop()
        root_path='/'.join(parts)

        # If this is a "root" file, append one more directory.
        # For example "goal_account.md" is a root because of commands like "goal_account_new.md"
        root_check = parts.copy()
        root_check.append(re.sub('\.md', '', new_name))
        extended_path = '_'.join(parts) + '_'
        is_root = any(extended_path in s for s in os.listdir(dirpath))
        if is_root:
            parts=root_check
            root_path='/'.join(parts)
        try:
            os.makedirs(dirpath + '/' + root_path)
        except FileExistsError:
            pass # this is expected to happen.

        new_file = dirpath + '/' + root_path + '/' + new_name
        os.rename(dirpath + '/' + f, new_file)

        # Rewrite file with fixed link paths
        fix_file_links(new_file, len(parts)-1)

        # Make sure root file is displayed first in the navigation par
        if is_root:
            with open(dirpath + '/' + root_path + '/.pages', 'w') as f:
                f.write('arrange:\n - %s' % new_name)

def fix_root(path):
    """ the algorithm puts everything one directory too deep, move it up. """
    files=[f for f in os.listdir(path) if os.path.isfile(os.path.join(path, f))]
    directories=[f for f in os.listdir(path) if os.path.isdir(os.path.join(path, f))]

    if len(files) != 0:
        print("WRONG NUMBER OF FILES IN ROOT PATH.")
        return
    if len(directories) != 1:
        print("WRONG NUMBER OF DIRECTORIES IN ROOT PATH.")
        return

    extra_dir = path + '/' + directories[0]
    shutil.move(extra_dir, path + '.tmp')
    os.rmdir(path)
    shutil.move(path + '.tmp', path)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Reformat markdown files to display in mkdocs.')    
    parser.add_argument('-path', required=True, help='Path to reformat.')
    args = parser.parse_args()    

    process(args.path)
    fix_root(args.path)
