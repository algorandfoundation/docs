import os

# Example Name => source string
SDKExamples = dict[str, list[str]]
def find_examples_in_sdk(dir: str, prefix: str, lang: str )->SDKExamples:
    directory = os.listdir(dir)

    name_to_src: SDKExamples = {} 
    for fname in directory:
        path = os.path.join(dir, fname)
        if not os.path.isfile(path):
            name_to_src |= find_examples_in_sdk(path, prefix, lang)
        elif path[-2:] in ["py", "js", "go", "va"]:
            local_example: list[str] = []
            with open(path, 'r') as f:
                content = f.read()
                if prefix not in content:
                    continue

                lines = content.splitlines()
                for line in lines:
                    if prefix in line:
                        name_to_src[line.strip(prefix)] = ["```"+lang, *local_example, "```"]
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
            start, stop = 0,0
            example_name = ""
            page_lines = []
            with open(path, 'r') as f:
                content = f.read()
                if prefix not in content:
                    continue
                    
                page_lines = content.splitlines()
                for lno, line in enumerate(page_lines):
                    if prefix in line:
                        if example_name == "":
                            example_name = line.strip(prefix + "->")
                            start = lno+1
                        else:
                            stop = lno

            if start != 0 and stop != 0:
                page_lines[start:stop] = examples[example_name]
                with open(path, "w") as f:
                    f.write("\n".join(page_lines))




def import_pysdk_examples():
    sdk_examples = find_examples_in_sdk("/home/ben/py-algorand-sdk/_examples", "# example: ", 'python')
    replace_matches_in_docs("../docs", "<!-- ===PYSDK_", sdk_examples)

if __name__ == "__main__":
    import_pysdk_examples()