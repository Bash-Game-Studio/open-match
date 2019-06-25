// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	pb "open-match.dev/open-match/pkg/pb"
	"testing"
)

func TestServiceHealth(t *testing.T) {
	om, closer := New(t)
	defer closer()
	if err := om.HealthCheck(); err != nil {
		t.Errorf("cannot create ticket, %s", err)
	}
}

func TestGetClients(t *testing.T) {
	om, closer := New(t)
	defer closer()
	fe, fec := om.MustFrontendGRPC()
	defer fec()
	if fe == nil {
		t.Error("cannot get frontend client")
	}
	be, bec := om.MustBackendGRPC()
	defer bec()
	if be == nil {
		t.Error("cannot get backend client")
	}
	mml, mmlc := om.MustMmLogicGRPC()
	defer mmlc()
	if mml == nil {
		t.Error("cannot get mmlogic client")
	}
}

func TestCreateTicket(t *testing.T) {
	om, closer := New(t)
	defer closer()
	fe, cc := om.MustFrontendGRPC()
	defer cc()
	resp, err := fe.CreateTicket(om.Context(), &pb.CreateTicketRequest{
		Ticket: &pb.Ticket{
			Properties: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"test-property": {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
				},
			},
		},
	})
	if err != nil {
		t.Errorf("cannot create ticket, %s", err)
	}
	t.Logf("created ticket %+v", resp)
}