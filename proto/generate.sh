#!/bin/bash

protoc -I. --go_out=com_ss_pb_proto *.proto