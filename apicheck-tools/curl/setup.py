from os import path
from setuptools import setup, find_packages

here = path.abspath(path.dirname(__file__))

setup(
    name="gurl",
    packages=find_packages(),
    install_requires=["httptools"]
)
