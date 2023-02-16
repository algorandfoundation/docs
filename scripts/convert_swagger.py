#!/usr/bin/env python3

import argparse
import os
from os.path import join, exists
import re
import subprocess
import shutil
import sys
import tempfile
import urllib.request

convert_command = "swagger2markup-cli.jar"
config = "swagger2markup.properties"

parser = argparse.ArgumentParser(
    description="Generate markdown files from swagger specfile."
)
parser.add_argument(
    "-processors",
    required=False,
    help="Optional file full of commands to fix broken swagger files. Replaces occurrances of '{specfile}' with the specfile",
)
parser.add_argument(
    "-target", required=True, help="Full path to output file, including the name."
)
parser.add_argument("-specfile", required=True, help="Swagger specfile.")


def apply_commands(command_file, specfile):
    """
    Given a file full of commands, applies them line by line to the specfile, replacing the file with stdout.
    """
    with open(command_file, "r") as commands:
        for command in commands:
            command = command.strip().replace("{specfile}", specfile)
            print("Running command: %s" % command)
            with tempfile.NamedTemporaryFile(mode="w", dir=".", delete=False) as tmp:
                subprocess.check_call([command], shell=True, cwd=".", stdout=tmp)
                os.replace(tmp.name, specfile)


def convert_swagger_to_markdown(specfile, directory):
    """
    Given a specfile generate markdown files in the specified directory
    """
    if not exists(convert_command):
        print("downloading %s..." % convert_command)
        url = "https://repo1.maven.org/maven2/io/github/swagger2markup/swagger2markup-cli/1.3.3/swagger2markup-cli-1.3.3.jar"
        filedata = urllib.request.urlretrieve(url, convert_command)
    if not exists(config):
        print("initializing %s..." % config)
        with open(config, "w") as f:
            f.write("swagger2markup.markupLanguage=MARKDOWN")

    # Generate markdown files
    command = "java -jar %s convert --swaggerInput %s --config %s --outputDir %s" % (
        convert_command,
        specfile,
        config,
        directory,
    )
    subprocess.check_call([command], shell=True)


def reformat_markdown(filename, outputfilename):
    """
    Takes output algod and kmd md files and reformats headings to look better with mkdocs.
    """

    raw = open(filename, "r").read()

    # Change headers to bolded so they do not show up in ToC
    step1 = re.sub("\#\#\#\# (\S+)", r"**\1**", raw)

    # Make the header the description text and the request path the header so it shows up in ToC
    step2 = re.sub(
        "\#\#\# (.+?)\n```\n((GET|PUT|POST|DELETE|PATCH) \S+)\n```",
        r"### \2\n\1\n```\n\2\n```",
        step1,
    )

    # Bring every header up a level
    matches = re.finditer("^\s*#+", step2)
    subtract = 0
    for match in matches:
        replace = match.group()
        start = match.start()
        end = match.end()
        step2 = step2[: start - subtract] + replace[:-1] + step2[end - subtract :]
        subtract += 1
    outputfile = open(outputfilename, "w")
    outputfile.write(step2)


if __name__ == "__main__":
    args = parser.parse_args()

    if args.processors is not None:
        apply_commands(args.processors, args.specfile)

    with tempfile.TemporaryDirectory() as tmpdir:
        # Generate markdown files.
        convert_swagger_to_markdown(args.specfile, tmpdir)

        title = args.target.split(os.sep)[-1].split(".")[0].lower()
        # Merge paths.md and definitions.md.
        merged_filename = join(tmpdir, "merged.md")
        with open(merged_filename, "wb") as merged:
            merged.write(str("title: %s\n---\n" % title).encode("utf-8"))
            for filename in ["paths.md", "definitions.md"]:
                markdown_file = join(tmpdir, filename)
                if not exists(markdown_file):
                    print("Missing '%s'" % filename)
                with open(markdown_file, "rb") as f:
                    shutil.copyfileobj(f, merged)

        # Process the markdown and update target.
        reformat_markdown(merged_filename, args.target)
