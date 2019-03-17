#!/bin/bash

for x in {1..100}; do
    test-client $1 &
    sleep 1
done

wait
