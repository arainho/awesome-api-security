
while read LINE; do

  # Get URL
  TARGET=$(echo $LINE | jq -c '.request.url')

  wfuzz "$@" "$TARGET"

done
