#!/usr/bin/env bash

if [ -z "$1" ]
  then
    echo "Usage $0 <appName>"
    exit -1
fi

appName=$1

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  find ./ -type f -not -path ".//.git/*" -not -path ".//rename.sh" -exec sed -i "s/myApi/$appName/g" {} \;
elif [[ "$OSTYPE" == "darwin"* ]]; then
  find ./ -type f -not -path ".//.git/*" -not -path ".//rename.sh" -exec sed -i '' "s/myApi/$appName/g" {} \;
else
  echo "Unsupported OS"
  exit -1
fi

cd ../ ; mv myApi $appName; cd $appName

echo "The app has renamed to $appName"

rm -- "$0"
