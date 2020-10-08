#!/bin/sh

###############################################################################
###                           FUNCTIONS		                            ###
###############################################################################

# Creates a validator for a given node
# Take 1 arg the name of the node e.g auth-chaindnode0
createValidator() {
	echo "Creating validator for node $1\n"

	# Create the validator
	docker exec -e MONIKER=$1 $1 /bin/sh -c 'auth-chaincli tx poa create-validator $(auth-chaincli keys show validator --bech val -a --keyring-backend test) $(auth-chaind tendermint show-validator) $(echo $MONIKER) identity website security@contact details -y --trust-node --from validator --chain-id auth --keyring-backend test'

	sleep 5
}

# Votes for a perspecitve canidate
# Take 2 args the name of the node voting and the candidate node e.g auth-chaindnode0 auth-chaindnode1
voteForValidator() {
	eval CANDIDATE=$(docker exec $2 /bin/sh -c "auth-chaincli keys show validator --bech val -a --keyring-backend test")
	echo "Voter $1 is voting for candidate $2"
	docker exec -e CANDIDATE=$CANDIDATE $1 /bin/sh -c 'auth-chaincli tx poa vote-validator $(echo $CANDIDATE) -y --trust-node --from validator --chain-id auth --keyring-backend test'

	sleep 5
}

# Kicks for a perspecitve canidate
# Take 2 args the name of the node voting and the candidate node e.g auth-chaindnode0 auth-chaindnode1
kickValidator() {
	eval CANDIDATE=$(docker exec $2 /bin/sh -c "auth-chaincli keys show validator --bech val -a --keyring-backend test")
	echo "Votee $1 is voting to kick candidate $2"
	docker exec -e CANDIDATE=$CANDIDATE $1 /bin/sh -c 'auth-chaincli tx poa kick-validator $(echo $CANDIDATE) -y --trust-node --from validator --chain-id auth --keyring-backend test'

	sleep 5
}
###############################################################################
###                           STEP 1		                            ###
###############################################################################

# Import the exported key for the first node
docker exec auth-chaindnode0 /bin/sh -c "echo -e 'password1234\n' | auth-chaincli keys import validator /root/validator --keyring-backend test"

## Create the validator
voteForValidator auth-chaindnode0 auth-chaindnode0

###############################################################################
###                           STEP 2		                            ###
###############################################################################

# Create the keys for each node
for var in auth-chaindnode1 auth-chaindnode2 auth-chaindnode3
do
	echo "Creating key for node $var\n"
	docker exec $var /bin/sh -c "auth-chaincli keys add validator --keyring-backend test"
done


## Send tokens to each validator
for node in auth-chaindnode1 auth-chaindnode2 auth-chaindnode3
do
	eval ADDRESS=$(docker exec $node /bin/sh -c "auth-chaincli keys show validator -a --keyring-backend test")
	echo "Sending tokens to $ADDRESS\n"
	docker exec -e ADDRESS=$ADDRESS auth-chaindnode0 /bin/sh -c 'auth-chaincli tx send $(auth-chaincli keys show validator -a --keyring-backend test) $(echo $ADDRESS) 100000stake -y --trust-node --from validator --chain-id auth --keyring-backend test'
	sleep 5
done

###############################################################################
###                           STEP 3		                            ###
###############################################################################

# Create validator for validator set
for var in auth-chaindnode1 auth-chaindnode2 auth-chaindnode3
do
	createValidator $var
done

###############################################################################
###                           STEP 4		                            ###
###############################################################################

# Adding new validators to the set

# Vote for validator1 to join the set
voteForValidator auth-chaindnode0 auth-chaindnode1

# auth-chaindnode1 votes for auth-chaindnode0 to prove the node is in the consensus
voteForValidator auth-chaindnode1 auth-chaindnode0

# auth-chaindnode1 votes for auth-chaindnode1 to stay relevant in the consensus
voteForValidator auth-chaindnode1 auth-chaindnode1

# auth-chaindnode1 and poanode0 votes for auth-chaindnode2 to join the consensus
voteForValidator auth-chaindnode0 auth-chaindnode2
voteForValidator auth-chaindnode1 auth-chaindnode2

# auth-chaindnode2 votes for auth-chaindnode2 to stay relevant in the consensus
voteForValidator auth-chaindnode2 auth-chaindnode2

# auth-chaindnode2 votes for auth-chaindnode1 to prove the node is in the consensus
voteForValidator auth-chaindnode2 auth-chaindnode1

# auth-chaindnode2 votes for auth-chaindnode0 to prove the node is in the consensus
voteForValidator auth-chaindnode2 auth-chaindnode0

# kick auth-chaindnode2 out of the consensus
kickValidator auth-chaindnode0 auth-chaindnode2
kickValidator auth-chaindnode1 auth-chaindnode2

echo "POA Consensus started with 2 nodes :thumbs_up:\n"

sleep 5

## Verify valdiators are in the set by checking the validator set
docker exec auth-chaindnode0 /bin/sh -c "curl -X GET 'localhost:26657/validators'"
