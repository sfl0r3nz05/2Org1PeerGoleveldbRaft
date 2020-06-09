# 2Org1PeerGoleveldbRaft

- benchmarks folder
  Contains the configuration files in order to test the smart contract.
- network folder
  Contains:
  network options to deploy in docker-compose files
  crypto configuration files (certificates, channel, genesis block, configtx.yaml, crypto-config.yaml)
  connection-profile to enable caliper to recognize the network
  prometheus/grafana configuration
- src folder contains chaincode source code

1. npm i @hyperledger/caliper-cli
2. npm init -y
3. npm install --only=prod @hyperledger/caliper-cli
4. npx caliper bind --caliper-bind-sut fabric --caliper-bind-sdk 2.0.0

- Run a network of 2CA, 2ORG, 1PEER/ORG, 3ORDERER, RAFT CONSENSUS, GOLEVELDB STATEDB

5. npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/fabcar/config.yaml --caliper-networkconfig networks/connection-profile/2org1peergoleveldb_raft/fabric-go-tls.yaml

- Run a network of 2CA, 2ORG, 1PEER/ORG, 3ORDERER, RAFT CONSENSUS, COUCHDB STATEDB

6. npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/fabcar/config.yaml --caliper-networkconfig networks/connection-profile/2org1peercouchdb_raft/fabric-go-tls-prometheus.yaml

- Run a network of CA, 2ORG, 1PEER/ORG, 2ORDERER, 3 ZOOKEEPEER, 4 KAFKA BROKERS, CONSENSUS KAFKA, GOLEVELDB STATEDB

7. npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/fabcar/config.yaml --caliper-networkconfig networks/connection-profile/2org1peergoleveldb_kafka/fabric-go-tls-prometheus.yaml

- Run a network of 2CA, 2ORG, 1PEER/ORG, 2ORDERER, 3 ZOOKEEPEER, 4 KAFKA BROKERS, CONSENSUS KAFKA, COUCHDB STATEDB

8. npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/fabcar/config.yaml --caliper-networkconfig networks/connection-profile/2org1peercouchdb_kafka/fabric-go-tls-prometheus.yaml

9. docker rm -f \$(docker ps -a -q)
10. docker volume rm \$(docker volume ls -q)
