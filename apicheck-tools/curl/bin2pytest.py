import sys


print(repr(open(sys.argv[1], mode="rb").read()))
