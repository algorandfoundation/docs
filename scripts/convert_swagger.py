#!/usr/bin/env python3

import re
import sys
	
def main(filename, outputfilename):
	"""
	Takes output algod and kmd md files and reformats headings to look better with mkdocs.
	"""

	raw = open(filename, 'r').read()

	# Change headers to bolded so they do not show up in ToC
	step1 = re.sub('\#\#\#\# (\S+)', r'**\1**', raw)

	# Make the header the description text and the request path the header so it shows up in ToC
	step2 = re.sub('\#\#\# (.+?)\n```\n((GET|PUT|POST|DELETE|PATCH) \S+)\n```', r'### \2\n\1\n```\n\2\n```', step1)
	
	# Bring every header up a level
	matches = re.finditer('#+', step2)
	subtract = 0
	for match in matches:
		replace = match.group()
		start = match.start()
		end = match.end()
		step2 = step2[:start - subtract] + replace[:-1] + step2[end-subtract:]
		subtract += 1
	outputfile = open(outputfilename, 'w')
	outputfile.write(step2)
	
main(sys.argv[1], sys.argv[2])

