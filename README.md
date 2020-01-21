About
----

Hundun

Features
---


## Incidents


1. List all incidents

eg. one that came in last 4hours

    hundun get incidents --since_relative "4h"


2.  describe an incident

    hundun describe incident --id "PJSQ2EJ"



## Resolve and Ack incidents


To ack incident

    hundun ack --incident_id "abcd"

## Add note to an incident


    hundun create note -i "abcd" -c "my note content goes here"


## List all notes of an incident

    hundun get notes -i "abcd"



## Alientvault

    hundun describe avalarm -a "3cc9e726-0edb-e178-2520-0aa8ba5236fb"


