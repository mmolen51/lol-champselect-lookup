#!/bin/sh

help="Welcome to lol lookup! This allows you to lookup a champion on the enemy team 
to see its counters and your mastery points on those counters.
Don't use any apostrophes or spaces.

Before using, be sure to set your summoner name with
    lol ign <your summoner name>

To search counters, type
    lol <champ name>

examples:
    lol lux
    lol kaisa
    lol nunu"

function lol() {
    if [ -z "$1" ]
    then 
        echo "$help"
    elif [ $1 = "ign" ]
    then
        if [ -z "$2" ]
        then
        echo "To set your summoner name, type
    lol ign <your summoner name>
If your name has spaces in it, please wrap it in quotation marks"
        else
        echo '{ "SummonerName": ''"'${2}'" }' > configs.json
        fi
    else
        ./main $1
    fi
}
