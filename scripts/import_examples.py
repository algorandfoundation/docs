#!/usr/bin/env python3

import os
from dataclasses import dataclass

SKIP_DIRS = [".venv", "__pycache__", "node_modules"]


@dataclass
class ExampleSource:
    """Represents a source for examples"""

    #: url to the github repo
    github_url: str
    #: branch name where examples can be found
    git_branch: str
    #: where to find the local repo
    local_dir: str
    #: where to find the example files
    example_dir: str
    #: full name of language
    language_name: str
    #: what to look for as a prefix in source examples
    src_comment_flag: str
    #: what file extensions to consider
    file_extension: str
    #: name for example source
    name: str

    def doc_comment_flag(self) -> str:
        return f"<!-- ==={self.name}_"

    def clone_url(self) -> str:
        return f"{self.github_url}.git"

    def file_url(self, file_name: str) -> str:
        if file_name.startswith(self.example_path()):
            file_name = file_name[len(self.example_path()) + 1 :]

        return (
            f"{self.github_url}/blob/{self.git_branch}/{self.example_dir}/{file_name}"
        )

    def example_path(self) -> str:
        return f"{self.local_dir}/{self.example_dir}"


@dataclass
class Example:
    """Represents a tagged example in source file"""

    path: str
    line_start: int
    lines: list[str]
    matches: int


@dataclass
class DocExampleMatch:
    """Represents a match between source and docs"""

    name: str
    line_start: int
    line_stop: int

    @staticmethod
    def empty() -> "DocExampleMatch":
        return DocExampleMatch("", 0, 0)


# Example Name => source lines
SDKExamples = dict[str, Example]

sources: list[ExampleSource] = [
    ExampleSource(
        github_url="https://github.com/barnjamin/py-algorand-sdk",
        git_branch="doc-examples",
        local_dir="../../py-algorand-sdk",
        example_dir="_examples",
        language_name="python",
        src_comment_flag="# example: ",
        name="PYSDK",
        file_extension=".py",
    ),
    ExampleSource(
        github_url="https://github.com/joe-p/js-algorand-sdk",
        git_branch="doc-examples",
        local_dir="../../js-algorand-sdk",
        example_dir="examples",
        language_name="javascript",
        src_comment_flag="// example: ",
        name="JSSDK",
        file_extension=".ts",
    ),
    ExampleSource(
        github_url="https://github.com/barnjamin/go-algorand-sdk",
        git_branch="examples",
        local_dir="../../go/src/github.com/algorand/go-algorand-sdk",
        example_dir="_examples",
        language_name="go",
        src_comment_flag="\t// example: ",
        name="GOSDK",
        file_extension=".go",
    ),
    ExampleSource(
        github_url="https://github.com/barnjamin/java-algorand-sdk",
        git_branch="examples",
        local_dir="../../java-algorand-sdk",
        example_dir="examples",
        language_name="java",
        src_comment_flag="// example: ",
        name="JAVASDK",
        file_extension=".java",
    ),
    ExampleSource(
        github_url="https://github.com/algorand-devrel/algorand-teal-examples",
        git_branch="master",
        local_dir="../../algorand-teal-examples",
        example_dir="_examples",
        language_name="teal",
        src_comment_flag="// example: ",
        name="TEAL",
        file_extension=".teal",
    ),
    ExampleSource(
        github_url="https://github.com/barnjamin/pyteal",
        git_branch="examples-for-docs",
        local_dir="../../pyteal",
        example_dir="_examples",
        language_name="python",
        src_comment_flag="# example: ",
        name="PYTEAL",
        file_extension=".py",
    ),
    ExampleSource(
        github_url="https://github.com/algorand-devrel/beaker",
        git_branch="master",
        local_dir="../../beaker",
        example_dir="examples",
        language_name="python",
        src_comment_flag="# example: ",
        name="BEAKER",
        file_extension=".py",
    ),
]


def find_examples_in_sdk(dir: str, prefix: str, lang: str, ext: str) -> SDKExamples:
    directory = os.listdir(dir)

    name_to_src: SDKExamples = {}
    for fname in directory:
        if fname in SKIP_DIRS:
            continue

        path = os.path.join(dir, fname)
        if not os.path.isfile(path):
            name_to_src |= find_examples_in_sdk(path, prefix, lang, ext)
        elif os.path.splitext(path)[-1] == ext:
            local_example: list[str] = []
            with open(path, "r") as f:
                content = f.read()
                if prefix not in content:
                    continue

                lines = content.splitlines()
                for lno, line in enumerate(lines):
                    if prefix in line:
                        name = line.strip(prefix)
                        name_to_src[name] = Example(
                            path=path,
                            line_start=lno - len(local_example),
                            lines=local_example,
                            matches=0,
                        )
                        local_example = []
                    else:
                        local_example.append(line)

    return name_to_src


def replace_matches_in_docs(
    dir: str, prefix: str, examples: SDKExamples, src: ExampleSource
):
    """recursively search in directory for string prefix"""
    directory = os.listdir(dir)
    for fname in directory:
        path = os.path.join(dir, fname)
        if not os.path.isfile(path):
            # recurse through directories
            replace_matches_in_docs(path, prefix, examples, src)
            continue
        elif path[-2:] != "md":
            continue

        page_lines: list[str] = []
        matches: list[DocExampleMatch] = []
        current_match = DocExampleMatch.empty()

        with open(path, "r") as f:
            content = f.read()
            if prefix not in content:
                continue

            page_lines = content.splitlines()
            for lno, line in enumerate(page_lines):
                if prefix not in line:
                    continue

                # First time finding seeing this one
                if current_match.name == "":
                    current_match.name = line.strip()[len(prefix) :].strip("= ->_")
                    current_match.line_start = lno + 1
                # Second time finding it, add it to matches and wipe current
                else:
                    current_match.line_stop = lno
                    matches.append(current_match)
                    current_match = DocExampleMatch.empty()

        if len(matches) == 0:
            continue

        # Need to track the offset here so we dont write to the
        # wrong spot in the doc file if the example is longer or shorter
        # than the current set of lines in the docs
        offset = 0
        for match in matches:

            if match.name not in examples:
                print(
                    f"Missing {match.name} in {prefix.strip(' -<!=_')} "
                    f"examples (in {path}:{match.line_start})"
                )
                continue

            src_example = examples[match.name]

            example_link = (
                src.file_url(src_example.path)
                + f"#L{src_example.line_start}-"
                + f"L{src_example.line_start + len(src_example.lines)}"
            )

            example_lines = [
                "```" + src.language_name,
                *src_example.lines,
                "```",
                f"[Snippet Source]({example_link})"
            ]

            page_lines[
                match.line_start + offset : match.line_stop + offset
            ] = example_lines

            offset += len(example_lines) - (match.line_stop - match.line_start)

            examples[match.name].matches += 1

        with open(path, "w") as f:
            f.write("\n".join(page_lines))

    return examples


def ensure_source(src: ExampleSource):
    import git

    if not os.path.isdir(src.local_dir):
        git.Repo.clone_from(src.clone_url(), src.local_dir, branch=src.git_branch)
    else:
        repo = git.Repo(src.local_dir)
        repo.git.checkout(src.git_branch)


if __name__ == "__main__":

    names = [src.name for src in sources]

    import argparse

    parser = argparse.ArgumentParser(description="Gather examples from source repos")
    parser.add_argument(
        "--src",
        metavar="name",
        type=str,
        nargs="*",
        choices=names,
        help="source names to pull (default: all)",
    )

    args = parser.parse_args()
    choices = args.src
    if choices is None:
        choices = names

    for src in sources:
        if src.name not in choices:
            continue

        ensure_source(src)

        sdk_examples = find_examples_in_sdk(
            src.example_path(),
            src.src_comment_flag,
            src.language_name,
            src.file_extension,
        )

        replace_matches_in_docs("../docs", src.doc_comment_flag(), sdk_examples, src)

        for name, example in sdk_examples.items():
            if example.matches == 0:
                print(
                    f"Missing {name} for {src.language_name} in docs "
                    f"(in: {example.path}:{example.line_start})"
                )
