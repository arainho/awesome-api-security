#!/bin/sh
#
# Copyright 2020 Banco Bilbao Vizcaya Argentaria, S.A.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

#!/bin/sh

# Eventually remove tmpfile on exit
tmpfile="/tmp/content.oas3"
trap "{ [ -f $tmpfile ] && rm -f $tmpfile; }" EXIT

# Process arguments. TODO more error checking.
if [ "$1" == "-r" ]
then
  rules="-r $2"
  shift; shift;
fi
file="$1"

# File from stdin
if [ -z "$file" ]; then
    cat /dev/stdin > "$tmpfile"
    file="$tmpfile"
# File from the web
elif [ -z "${file##http://*}" ] || [ -z "${file##https://*}" ]; then
    wget "$file" -qo "$tmpfile"
    file="$tmpfile"
# Existing local file
elif [ -f "$file" ]; then
  cp "$file" "$tmpfile"
else
  echo "File $file not found"
  exit 1
fi

spectral lint ${rules} "$tmpfile"
statusCode=$?

if [ "$statusCode" -eq 0 ]
then
  cat "$tmpfile"
fi

exit $statusCode
