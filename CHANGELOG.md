# Changelog

## v0.1.0-rc.1

_July 5th, 2018_

### FEATURES

- Gas fee: Charge Validator for declaring candidacy, updating candidate information and proposing transactions. 
- Governance Transactions: Change system parameters through governance transactions.

### IMPROVEMENTS

- Update tendermit to v0.20.0.
- Add Candidate information fields: name, email, profile.
- Change parameters of ChainId:  18: mainnet, 19: testnet, 20: staging.
- Add cmt.syncing to get node syncing status.

### FIXES

- If the maximum staking amount decreases, Validator self-staked CMTs won't be charged.
- Correct non-running Validators won’t get block awards.
- Fix Block Award calculation error.
- Correct delegator address when a Validator withdraws candidacy.
