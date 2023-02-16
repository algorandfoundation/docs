import os
from dataclasses import dataclass


@dataclass
class ExampleSource:
    example_dir: str
    language_name: str
    src_comment_flag: str
    doc_comment_flag: str


sources: list[ExampleSource] = [
    ExampleSource(
        example_dir="../../py-algorand-sdk/_examples",
        language_name="python",
        src_comment_flag="# example: ",
        doc_comment_flag="<!-- ===PYSDK_",
    ),
    # ExampleSource(
    #    example_dir="../../js-algorand-sdk/examples",
    #    language_name='javascript',
    #    src_comment_flag="// example: ",
    #    doc_comment_flag="<!-- ===JSSDK_"
    # ),
    # ExampleSource(
    #    example_dir="../../go/src/github.com/algorand/go-algorand-sdk/examples",
    #    language_name='go',
    #    src_comment_flag="// example: ",
    #    doc_comment_flag="<!-- ===GOSDK_"
    # ),
    # ExampleSource(
    #    example_dir="../../java-algorand-sdk/examples",
    #    language_name='java',
    #    src_comment_flag="// example: ",
    #    doc_comment_flag="<!-- ===JAVASDK_"
    # ),
    # ExampleSource(
    #    example_dir="../../pyteal/examples",
    #    language_name='python',
    #    src_comment_flag="# example: ",
    #    doc_comment_flag="<!-- ===PYTEAL_"
    # ),
    # ExampleSource(
    #    example_dir="../../beaker/examples",
    #    language_name='python',
    #    src_comment_flag="# example: ",
    #    doc_comment_flag="<!-- ===BEAKER_"
    # )
]


# Example Name => source string
SDKExamples = dict[str, list[str]]


def find_examples_in_sdk(dir: str, prefix: str, lang: str) -> SDKExamples:
    directory = os.listdir(dir)

    name_to_src: SDKExamples = {}
    for fname in directory:
        path = os.path.join(dir, fname)
        if not os.path.isfile(path):
            name_to_src |= find_examples_in_sdk(path, prefix, lang)
        elif path[-2:] in ["py", "js", "go", "va"]:
            local_example: list[str] = []
            with open(path, "r") as f:
                content = f.read()
                if prefix not in content:
                    continue

                lines = content.splitlines()
                for line in lines:
                    if prefix in line:
                        name_to_src[line.strip(prefix)] = [
                            "```" + lang,
                            *local_example,
                            "```",
                        ]
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
        elif path[-2:] == "md":
            start, stop = 0, 0
            example_name = ""
            page_lines = []
            with open(path, "r") as f:
                content = f.read()
                if prefix not in content:
                    continue

                page_lines = content.splitlines()
                for lno, line in enumerate(page_lines):
                    if prefix in line:
                        if example_name == "":
                            example_name = line.strip(prefix + "->")
                            start = lno + 1
                        else:
                            stop = lno

            if start != 0 and stop != 0:
                page_lines[start:stop] = examples[example_name]
                with open(path, "w") as f:
                    f.write("\n".join(page_lines))


if __name__ == "__main__":
    for src in sources:
        sdk_examples = find_examples_in_sdk(
            src.example_dir, src.src_comment_flag, src.language_name
        )
        replace_matches_in_docs("../docs", src.doc_comment_flag, sdk_examples)
