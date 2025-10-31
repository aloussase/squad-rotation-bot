#!/bin/bash

API_BASE='http://localhost:3000'

list_members() {
  local members=`curl "$API_BASE/api/v1/rotation/members" 2>/dev/null | jq '.[] | .FullName + "," + .AvatarUrl' | tr -d '"'`
  echo "$members" | gum table --columns="Name,Avatar" --separator="," --print
}

create_member() {
  gum style \
    --foreground "#00e0ee" \
    --padding '1 1' \
    --bold \
    "Add a squad member"

  local full_name=`gum input --placeholder 'Full name' --padding '1 1'` 
  local avatar_url=`gum input --placeholder 'Avatar URL' --padding '1 1'`

  gum style \
    --foreground "#00e0ee" \
    --padding "1 1" \
    "The squad member '$full_name' is going to be created with avatar: '$avatar_url'"

  if gum confirm 'Proceed with creation?'; then
    local result=`curl -X POST -i -s "$API_BASE/api/v1/rotation/members" -d "{ \"FullName\": \"$full_name\", \"AvatarUrl\": \"$avatar_url\" }"`
    local status_code=`echo $result | head -n1 | cut -d' ' -f2`
    if [ "$status_code" = '201' ]; then
      gum style --foreground "#00ff00" --bold --padding "1 1" "User created successfully"
    else
      gum style --foreground "#ff0000" --bold --padding "1 1" "There was an error creating the user"
    fi
  else
    gum style --foreground "#ff0000" --bold --padding "1 1" 'Operation aborted'
  fi
}

read -r -d '' menu_options <<'EOF'
1. List squad members
2. Add squad member
3. Exit
EOF

while :; do
  choice=`echo "$menu_options" | gum choose --header "Menu Options" --padding "1 1" --header.bold | cut -d'.' -f1`
  case $choice in
    1)
      list_members
      ;;
    2)
      create_member
      ;;
    3)
      exit 0
      ;;
  esac
done
