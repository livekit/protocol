package livekit

import (
	"testing"

	"github.com/stretchr/testify/require"
	proto "google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

func TestUnmarshallRoomConfiguration(t *testing.T) {
	y := `
a:
  name: room_name
  egress:
    room:
      room_name: room_name
  agent:
    dispatches:
      - {}
      - agent_name: ag
        metadata: mm
  min_playout_delay: 42
`

	obj := make(map[string]*RoomConfiguration)

	err := yaml.Unmarshal([]byte(y), &obj)
	require.NoError(t, err)
	require.Equal(t, 1, len(obj))

	re := obj["a"]
	require.NotNil(t, re)
	require.Equal(t, re.Name, "room_name")
	require.Equal(t, re.MinPlayoutDelay, uint32(42))
	require.Equal(t, re.Egress.Room.RoomName, "room_name")
	require.Equal(t, len(re.Agent.Dispatches), 2)
	require.Equal(t, "ag", re.Agent.Dispatches[1].AgentName)
	require.Equal(t, "mm", re.Agent.Dispatches[1].Metadata)

}

func TestMarshallRoomConfiguration(t *testing.T) {
	r := &RoomConfiguration{
		Name:             "name",
		MaxParticipants:  42,
		EmptyTimeout:     12,
		DepartureTimeout: 13,
		MinPlayoutDelay:  14,
		MaxPlayoutDelay:  15,
		Egress: &RoomEgress{
			Room: &RoomCompositeEgressRequest{
				AudioOnly: true,
				RoomName:  "room name",
			},
		},
		Agent: &RoomAgent{
			Dispatches: []*RoomAgentDispatch{
				&RoomAgentDispatch{
					AgentName: "agent name",
				},
			},
		},
	}

	b, err := yaml.Marshal(r)
	require.NoError(t, err)

	var ur RoomConfiguration
	err = yaml.Unmarshal(b, &ur)
	require.NoError(t, err)
	require.True(t, proto.Equal(r, &ur))
}

func TestUnmarshallRoomEgress(t *testing.T) {
	y := `
a:
  room:
    room_name: room name
b:
  participant:
    file_outputs:
        - s3:
            access_key: key 
`

	obj := make(map[string]*RoomEgress)

	err := yaml.Unmarshal([]byte(y), &obj)
	require.NoError(t, err)
	require.Equal(t, 2, len(obj))

	re := obj["a"]
	require.NotNil(t, re)
	require.Equal(t, re.Room.RoomName, "room name")

	re = obj["b"]
	require.NotNil(t, re)
	require.Equal(t, 1, len(re.Participant.FileOutputs))
	require.Equal(t, "key", re.Participant.FileOutputs[0].Output.(*EncodedFileOutput_S3).S3.AccessKey)
}

func TestMarshallRoomEgress(t *testing.T) {
	e := &RoomEgress{
		Room: &RoomCompositeEgressRequest{
			AudioOnly: true,
			RoomName:  "room name",
		},
	}

	b, err := yaml.Marshal(e)
	require.NoError(t, err)

	var ue RoomEgress
	err = yaml.Unmarshal(b, &ue)
	require.NoError(t, err)
	require.True(t, proto.Equal(e, &ue))
}

func TestUnmarshallRoomAgent(t *testing.T) {
	y := `
a:
  dispatches:
    - {}
    - agent_name: ag
      metadata: mm
`
	obj := make(map[string]*RoomAgent)

	err := yaml.Unmarshal([]byte(y), &obj)
	require.NoError(t, err)
	require.Equal(t, 1, len(obj))

	re := obj["a"]
	require.NotNil(t, re)
	require.Equal(t, len(re.Dispatches), 2)
	require.Equal(t, "ag", re.Dispatches[1].AgentName)
	require.Equal(t, "mm", re.Dispatches[1].Metadata)
}

func TestMarshallRoomAgent(t *testing.T) {
	a := &RoomAgent{
		Dispatches: []*RoomAgentDispatch{
			&RoomAgentDispatch{
				AgentName: "agent name",
			},
		},
	}

	b, err := yaml.Marshal(a)
	require.NoError(t, err)

	var ua RoomAgent
	err = yaml.Unmarshal(b, &ua)
	require.NoError(t, err)
	require.True(t, proto.Equal(a, &ua))
}
