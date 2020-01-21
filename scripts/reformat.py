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

parser = argparse.ArgumentParser(description='Reformat markdown files to display in mkdocs.')    
parser.add_argument('-path', required=True, help='Path to reformat.')
parser.add_argument('-cmd', required=True, help='Name of the command (goal).')

def fix_file_links(new_file, goal_depth, with_subcommand):
    """ Update markdown links to use relative paths in different directories. """
    with tempfile.NamedTemporaryFile(mode='w', dir='.', delete=False) as tmp, \
            open(new_file, 'r') as f:
        for line in f:
            result = line
            m = re.search('(^\* \[.*\]\()(.*)(\).*$)', result)
            # Most lines wont match, write them unmodified
            if m != None:
                # Grab the command, i.e. 'goal' in 'goal.md', or 'delete' in 'goal_account_multisig_delete.md'

                # Find out how many levels away the link is (siblings need ../, parents need ../../)
                link_parts = m.group(2).split('.')[0].split('_')
                subcommand = link_parts[-1]
                prefix_mul = goal_depth - link_parts.index(subcommand) + 2

                # Subcommands have an extra level of nesting
                if subcommand in with_subcommand:
                    link = '../'*prefix_mul + ((subcommand + '/')*2)
                else:
                    link = '../'*prefix_mul + subcommand + '/'
                result = "%s%s%s" % (m.group(1), link, m.group(3))
            tmp.write(result + "\n")
        os.replace(tmp.name, new_file)
    

def process(dirpath, cmd):
    """ move files into a directory structure and add .pages files. """

    with_subcommand = []
    moved_files = []
    files = os.listdir(dirpath)
    for f in files:
        parts = f.split('_')
        # Ignore the root "goal.md" file
        new_name=parts.pop()
        root_path='/'.join(parts)

        # If this is a "root" file, append one more directory.
        # For example "goal_account.md" is a root because of commands like "goal_account_new.md"
        root_check = parts.copy()
        root_check.append(re.sub('\.md', '', new_name))
        extended_path = '_'.join(root_check) + '_'
        is_root = any(extended_path in s for s in files)
        if is_root:
            with_subcommand.append(root_check[-1])
            parts=root_check
            root_path='/'.join(parts)
        try:
            os.makedirs(dirpath + '/' + root_path)
        except FileExistsError:
            pass # this is expected to happen.

        new_file = dirpath + '/' + root_path + '/' + new_name
        os.rename(dirpath + '/' + f, new_file)
        moved_files.append((new_file, len(parts)-1))

        # Rewrite file with fixed link paths (except for root command)
        #if f != (cmd + '.md'):
        #    fix_file_links(new_file, len(parts)-1, None)

        # Make sure root file is displayed first in the navigation par
        if is_root:
            with open(dirpath + '/' + root_path + '/.pages', 'w') as f:
                f.write('arrange:\n - %s' % new_name)
    # Fix the links at the very end so that we know which ones have subcommands
    for f,depth in moved_files:
        fix_file_links(f, depth, with_subcommand)
    return len(moved_files)

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
    args = parser.parse_args()

    files_modified = process(args.path, args.cmd)
    fix_root(args.path)

    print("Finished formatting %d files." % files_modified)
