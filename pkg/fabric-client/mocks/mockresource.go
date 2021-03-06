/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"bytes"

	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

// MockResource ...
type MockResource struct {
	errorScenario bool
}

// NewMockInvalidResource ...
func NewMockInvalidResource() *MockResource {
	c := &MockResource{errorScenario: true}
	return c
}

// NewMockResource ...
func NewMockResource() *MockResource {
	return &MockResource{}
}

// ExtractChannelConfig ...
func (c *MockResource) ExtractChannelConfig(configEnvelope []byte) ([]byte, error) {
	if bytes.Compare(configEnvelope, []byte("ExtractChannelConfigError")) == 0 {
		return nil, errors.New("Mock extract channel config error")
	}

	return configEnvelope, nil
}

// SignChannelConfig ...
func (c *MockResource) SignChannelConfig(config []byte, signer fab.IdentityContext) (*common.ConfigSignature, error) {
	if bytes.Compare(config, []byte("SignChannelConfigError")) == 0 {
		return nil, errors.New("Mock sign channel config error")
	}
	return nil, nil
}

// CreateChannel ...
func (c *MockResource) CreateChannel(request fab.CreateChannelRequest) (fab.TransactionID, error) {
	if c.errorScenario {
		return fab.TransactionID{}, errors.New("Create Channel Error")
	}

	return fab.TransactionID{}, nil
}

//QueryChannels ...
func (c *MockResource) QueryChannels(peer fab.Peer) (*pb.ChannelQueryResponse, error) {
	return nil, errors.New("Not implemented yet")
}

//QueryInstalledChaincodes mocks query installed chaincodes
func (c *MockResource) QueryInstalledChaincodes(peer fab.Peer) (*pb.ChaincodeQueryResponse, error) {
	if peer == nil {
		return nil, errors.New("Generate Error")
	}
	ci := &pb.ChaincodeInfo{Name: "name", Version: "version", Path: "path"}
	response := &pb.ChaincodeQueryResponse{Chaincodes: []*pb.ChaincodeInfo{ci}}
	return response, nil
}

// InstallChaincode mocks install chaincode
func (c *MockResource) InstallChaincode(req fab.InstallChaincodeRequest) ([]*fab.TransactionProposalResponse, string, error) {
	if req.Name == "error" {
		return nil, "", errors.New("Generate Error")
	}

	if req.Name == "errorInResponse" {
		result := fab.TransactionProposalResult{Endorser: "http://peer1.com", Status: 10}
		response := &fab.TransactionProposalResponse{TransactionProposalResult: result, Err: errors.New("Generate Response Error")}
		return []*fab.TransactionProposalResponse{response}, "1234", nil
	}

	result := fab.TransactionProposalResult{Endorser: "http://peer1.com", Status: 0}
	response := &fab.TransactionProposalResponse{TransactionProposalResult: result}
	return []*fab.TransactionProposalResponse{response}, "1234", nil
}
