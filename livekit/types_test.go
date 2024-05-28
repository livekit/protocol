package livekit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

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

func TestUnmarshallRoomAgent(t *testing.T) {
	y := `
a:
  agents:
    - type: 0 
      room: test room
    - type: 1
      participant_identity:
        - id1
        - id2
      namespace: nm
`
	obj := make(map[string]*RoomAgent)

	err := yaml.Unmarshal([]byte(y), &obj)
	require.NoError(t, err)
	require.Equal(t, 1, len(obj))

	re := obj["a"]
	require.NotNil(t, re)
	require.Equal(t, len(re.Agents), 2)
	require.Equal(t, JobType_JT_ROOM, re.Agents[0].Type)
	require.Equal(t, "test room", re.Agents[0].Room)
	require.Equal(t, JobType_JT_PUBLISHER, re.Agents[1].Type)
	require.Equal(t, "nm", re.Agents[1].Namespace)
	require.Equal(t, 2, len(re.Agents[1].ParticipantIdentity))
	require.Equal(t, "id1", re.Agents[1].ParticipantIdentity[0])
	require.Equal(t, "id2", re.Agents[1].ParticipantIdentity[1])
}
