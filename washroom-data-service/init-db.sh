#!/bin/sh
sqlite3 washrooms.db "$(cat ./repository/sqlite/schema.sql)"