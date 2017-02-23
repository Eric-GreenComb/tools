#!/bin/bash

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBank.json" http://localhost:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerFund.json" http://localhost:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerChannel.json" http://localhost:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor01.json" http://localhost:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract01.json" http://localhost:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain01.json" http://localhost:7050/chaincode
echo ""

sleep 2

curl -H "Content-Type: application/json" -X POST --raw --data "@coinbase.json" http://localhost:7050/chaincode
echo ""
