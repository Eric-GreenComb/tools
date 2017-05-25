#!/bin/bash

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBank.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerFund.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerChannel.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor01.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor02.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor03.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor04.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor05.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor06.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor07.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor08.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor09.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerDonor10.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract01.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract02.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract03.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract04.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerSmartContract05.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain01.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain02.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain03.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain04.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain05.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain06.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain07.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain08.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain09.json" http://192.168.1.20:7050/chaincode
echo ""

curl -H "Content-Type: application/json" -X POST --raw --data "@registerBargain10.json" http://192.168.1.20:7050/chaincode
echo ""

sleep 2

curl -H "Content-Type: application/json" -X POST --raw --data "@coinbase.json" http://192.168.1.20:7050/chaincode
echo ""
