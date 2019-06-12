#!/bin/bash

psql nutrition < createDerivation.sql
psql nutrition < createProducts.sql
psql nutrition < createServing.sql
psql nutrition < createNutrient.sql

