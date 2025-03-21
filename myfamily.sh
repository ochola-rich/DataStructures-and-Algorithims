#!/bin/bash

curl https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ".[] | select( .id == $HERO_ID ) | .connections .relatives" | sed 's/"//g'
