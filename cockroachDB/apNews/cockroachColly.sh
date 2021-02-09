#!/bin/bash

# start the db
./startDB.sh &
wait
echo ""
echo "####################################"
echo "---DATABASE STARTED---"
echo "####################################"
sleep 3

# create and populate tables
cockroach sql --insecure < creates.sql
echo ""
echo "####################################"
echo "---TABLES CRATED AND POPPULATED---"
echo "####################################"

# run scraper
./articleScrapeDB
echo "####################################"
echo "--- WEBSITE SCRAPED AND DATA UPLOADED---"
echo "####################################"
echo ""

cockroach sql --execute="use apnews_db; select * from article;" --insecure
