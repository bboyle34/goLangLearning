#!/bin/bash

# start the db
./startDB.sh &
wait
sleep 3
echo ""
echo "####################################"
echo "---DATABASE STARTED---"
echo "####################################"
sleep 3
# create and populate tables
cockroach sql --insecure < creates.sql
wait
echo ""
echo "####################################"
echo "---TABLES CRATED AND POPPULATED---"
echo "####################################"
sleep 3
# run scraper
./articleScrapeDB
wait
echo "####################################"
echo "--- WEBSITE SCRAPED AND DATA UPLOADED---"
echo "####################################"
echo ""

cockroach sql --execute="use apnews_db; select * from article;" --insecure
