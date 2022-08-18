#!/bin/bash

PS3='Please enter your choice: '
options=("Dev" "Build" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Dev")
            echo "Starting Dev server"
            concurrently -n client,backend -c "green,cyan" "cd client && yarn dev" "cd backend && go run ."
            break
            ;;
        "Build")
            echo "Building client and backend"
            concurrently -n client,backend -c "green,cyan" "cd client && yarn build" "cd backend && go build"
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done
