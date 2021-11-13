from setuptools import setup

requirements = open("requirements.txt", "r").readlines()

setup(install_requires=requirements)
