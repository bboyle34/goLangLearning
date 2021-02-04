#!/bin/bash

cd /home/bboyle/github/goLangLearning/cockroachDB/
#./startDB.sh $

./connectDB.sh

use apnews_db;

select * from article;
