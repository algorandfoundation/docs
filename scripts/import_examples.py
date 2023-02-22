#!/usr/bin/env python3

import os
from dataclasses import dataclass

SKIP_DIRS = [".venv", "node_modules"]


@dataclass
class ExampleSource:
    """Represents a source for examples"""

    example_dir: str
    language_name: str
    src_comment_flag: str
    doc_comment_flag: str
    file_extension: str


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
        example_dir="../../py-algorand-sdk/_examples",
        language_name="python",
        src_comment_flag="# example: ",
        doc_comment_flag="<!-- ===PYSDK_",
        file_extension=".py",
    ),
    ExampleSource(
        example_dir="../../js-algorand-sdk/examples",
        language_name="javascript",
        src_comment_flag="// example: ",
        doc_comment_flag="<!-- ===JSSDK_",
        file_extension=".js",
    ),
    ExampleSource(
        example_dir="../../go/src/github.com/algorand/go-algorand-sdk/examples",
        language_name="go",
        src_comment_flag="// example: ",
        doc_comment_flag="<!-- ===GOSDK_",
        file_extension=".go",
    ),
    ExampleSource(
        example_dir="../../java-algorand-sdk/examples",
        language_name="java",
        src_comment_flag="// example: ",
        doc_comment_flag="<!-- ===JAVASDK_",
        file_extension=".java",
    ),
    ExampleSource(
        example_dir="../../algorand-teal-examples/_examples",
        language_name="teal",
        src_comment_flag="// example: ",
        doc_comment_flag="<!-- ===TEAL_",
        file_extension=".teal",
    ),
    ExampleSource(
        example_dir="../../pyteal/examples",
        language_name="python",
        src_comment_flag="# example: ",
        doc_comment_flag="<!-- ===PYTEAL_",
        file_extension=".py",
    ),
    ExampleSource(
        example_dir="../../beaker/examples",
        language_name="python",
        src_comment_flag="# example: ",
        doc_comment_flag="<!-- ===BEAKER_",
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
                            lines=[
                                "```" + lang,
                                *local_example,
                                "```",
                            ],
                            matches=0,
                        )
                        local_example = []
                    else:
                        local_example.append(line)

    return name_to_src


def replace_matches_in_docs(dir: str, prefix: str, examples: SDKExamples):
    """recursively search in directory for string prefix"""
    directory = os.listdir(dir)
    for fname in directory:
        path = os.path.join(dir, fname)
        if not os.path.isfile(path):
            # recurse through directories
            replace_matches_in_docs(path, prefix, examples)
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

            page_lines[match.line_start + offset : match.line_stop + offset] = examples[
                match.name
            ].lines

            offset += len(examples[match.name].lines) - (
                match.line_stop - match.line_start
            )

            examples[match.name].matches += 1

        with open(path, "w") as f:
            f.write("\n".join(page_lines))

    return examples


if __name__ == "__main__":
    for src in sources:
        if not os.path.isdir(src.example_dir):
            print(
                f"Missing {src.doc_comment_flag.strip(' -<!=_')} "
                f"example directory ({src.example_dir})"
            )
            continue

        sdk_examples = find_examples_in_sdk(
            src.example_dir, src.src_comment_flag, src.language_name, src.file_extension
        )
        replace_matches_in_docs("../docs", src.doc_comment_flag, sdk_examples)

        for name, example in sdk_examples.items():
            if example.matches == 0:
                print(
                    f"Missing {name} for {src.language_name} in docs "
                    f"(in: {example.path}:{example.line_start})"
                )
