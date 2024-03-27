#!/bin/bash

total_commands=1000

# Temps (en secondes) entre chaque commande
interval=0.06 # 60 secondes / 1000 commandes

for ((i=1; i<=total_commands; i++))
do
	curl 127.0.0.1:80
    sleep $interval
done

