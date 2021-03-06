package governance

import (
	"github.com/CyberMiles/travis/utils"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
	"golang.org/x/crypto/ripemd160"
)

const TRANSFER_FUND_PROPOSAL = "transfer_fund"
const CHANGE_PARAM_PROPOSAL = "change_param"
const DEPLOY_LIBENI_PROPOSAL = "deploy_libeni"

type Proposal struct {
	Id           string
	Type         string
	Proposer     *common.Address
	BlockHeight  int64
	ExpireTimestamp       int64
	ExpireBlockHeight     int64
	CreatedAt    string
	Result       string
	ResultMsg    string
	ResultBlockHeight    int64
	ResultAt     string
	Detail       map[string]interface{}
}

func (p *Proposal) Hash() []byte {
	var status interface{}
	if p.Detail != nil {
		if status = p.Detail["status"]; status != nil {
			delete(p.Detail, "status")
		}
	}
	pp, err := json.Marshal(struct {
		Id           string
		Type         string
		Proposer     *common.Address
		BlockHeight  int64
		ExpireTimestamp       int64
		ExpireBlockHeight     int64
		Result       string
		ResultMsg    string
		ResultBlockHeight    int64
		Detail       map[string]interface{}
	}{
		p.Id,
		p.Type,
		p.Proposer,
		p.BlockHeight,
		p.ExpireTimestamp,
		p.ExpireBlockHeight,
		p.Result,
		p.ResultMsg,
		p.ResultBlockHeight,
		p.Detail,
	})
	if err != nil {
		panic(err)
	}
	if status != nil {
		p.Detail["status"] = status
	}
	hasher := ripemd160.New()
	hasher.Write(pp)
	return hasher.Sum(nil)
}

func NewTransferFundProposal(id string, proposer *common.Address, blockHeight int64, from *common.Address, to *common.Address, amount string, reason string, expireTimestamp, expireBlockHeight int64) *Proposal {
	now := utils.GetNow()
	return &Proposal {
		id,
		TRANSFER_FUND_PROPOSAL,
		proposer,
		blockHeight,
		expireTimestamp,
		expireBlockHeight,
		now,
		"",
		"",
		0,
		"",
		map[string]interface{}{
			"from": from,
			"to": to,
			"amount": amount,
			"reason": reason,
		},
	}
}

func NewChangeParamProposal(id string, proposer *common.Address, blockHeight int64, name, value, reason string, expireTimestamp, expireBlockHeight int64) *Proposal {
	now := utils.GetNow()
	return &Proposal {
		id,
		CHANGE_PARAM_PROPOSAL,
		proposer,
		blockHeight,
		expireTimestamp,
		expireBlockHeight,
		now,
		"",
		"",
		0,
		"",
		map[string]interface{}{
			"name": name,
			"value": value,
			"reason": reason,
		},
	}
}

func NewDeployLibEniProposal(id string, proposer *common.Address, blockHeight int64, name, version, fileurl, md5, reason, status string, expireTimestamp, expireBlockHeight int64) *Proposal {
	now := utils.GetNow()
	return &Proposal {
		id,
		DEPLOY_LIBENI_PROPOSAL,
		proposer,
		blockHeight,
		expireTimestamp,
		expireBlockHeight,
		now,
		"",
		"",
		0,
		"",
		map[string]interface{}{
			"name": name,
			"version": version,
			"fileurl": fileurl,
			"md5": md5,
			"reason": reason,
			"status": status,
		},
	}
}

type Vote struct {
	ProposalId     string
	Voter          common.Address
	BlockHeight    int64
	Answer         string
	CreatedAt      string
}

func (v *Vote) Hash() []byte {
	vote, err := json.Marshal(struct {
		ProposalId     string
		Voter          common.Address
		BlockHeight    int64
		Answer         string
	}{
		v.ProposalId,
		v.Voter,
		v.BlockHeight,
		v.Answer,
	})
	if err != nil {
		panic(err)
	}
	hasher := ripemd160.New()
	hasher.Write(vote)
	return hasher.Sum(nil)
}

func NewVote(proposalId string, voter common.Address, blockHeight int64, answer string) *Vote {
	now := utils.GetNow()
	return &Vote {
		proposalId,
		voter,
		blockHeight,
		answer,
		now,
	}
}

