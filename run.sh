#!/bin/bash

PS3='Please enter your choice: '
options=("Dev" "Prod" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Dev")
            echo "Starting Dev server"
            cd client && npm run dev && cd .. &&  cd backend && go run ./ && cd ..
            break
            ;;
        "Prod")
            echo "Starting Prod server"
             cd client && npm run build && npm run preview && cd .. &&  cd backend && go build && ./backend && cd .. 
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done


# if [$2 == "dev"] 
# then 
#     cd backend && go run ./ && cd ..
#     cd client && npm run dev && cd ..

# elif [$2 == "build"]
# then
#     cd backend && go build && ./backend && cd ..
#     cd client && npm run build && npm run preview && cd ..
# fi