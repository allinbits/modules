# poa module specification

## Abstract

This paper specifies a `poa` module of the Cosmos-SDK. The `poa` module will define a Proof-Of-Authority (POA) system that will allow the Cosmos-SDK to leverage a POA algorithm within the ecosystem.

The `poa` module will have the following characteristics:

- Trusted set of validators on a network
- Produces blocks at a reliable rate
- Not as many adversarial conditions on the network
- Potentially higher performance than other consensus algorithms

Two main data structures are used to allow the `poa` module to function correctly:

- `Validator`: Represents consensus validators that are used to create and confirm blocks
- `Vote`: Allows verified validators to vote if new validators can join the consensus

## Contents

1. **[State](01_state.md)**
    - [Validator](01_state.md#validator)
    - [Vote](01_state.md#vote)
1. **[State Transitions](02_state_transitions.md)**
    - [Validator](02_state_transitions.md#validator)
    - [Vote](02_state_transitions.md#vote)
1. **[Messages](03_messages.md)**
    - [MsgCreateValidator](03_messages.md#MsgCreateValidator)
    - [MsgVoteVote](03_messages.md#MsgCastVote)
1. **[Begin-Block](04_begin_block.md)**
    - [Vote changes](04_begin_block#Vote-changes)
1. **[End-Block](05_end_block.md)**
    - [Validator set changes](05_end_block#Validator-set-changes)
1. **[Events](06_events.md)**
1. **[Parameters](07_parameters.md)**
