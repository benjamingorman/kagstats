// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fixtures

import (
	"github.com/Harrison-Miller/kagstats/collector/database"
	"github.com/Harrison-Miller/kagstats/common/models"
	"sync"
)

// Ensure, that DatabaseMock does implement database.Database.
// If this is not the case, regenerate this file with moq.
var _ database.Database = &DatabaseMock{}

// DatabaseMock is a mock implementation of database.Database.
//
// 	func TestSomethingThatUsesDatabase(t *testing.T) {
//
// 		// make and configure a mocked database.Database
// 		mockedDatabase := &DatabaseMock{
// 			CommitFunc: func(kills []models.Kill) error {
// 				panic("mock out the Commit method")
// 			},
// 			CommitFlagCaptureFunc: func(capture models.FlagCapture) error {
// 				panic("mock out the CommitFlagCapture method")
// 			},
// 			CommitMapStatsFunc: func(stats models.MapStats) error {
// 				panic("mock out the CommitMapStats method")
// 			},
// 			CommitMapVotesFunc: func(votes models.MapVotes) error {
// 				panic("mock out the CommitMapVotes method")
// 			},
// 			CommitPlayerFunc: func(player *models.Player) error {
// 				panic("mock out the CommitPlayer method")
// 			},
// 			InitDBFunc: func() error {
// 				panic("mock out the InitDB method")
// 			},
// 			UpdatePlayerInfoFunc: func(player *models.Player) error {
// 				panic("mock out the UpdatePlayerInfo method")
// 			},
// 			UpdateServerInfoFunc: func(server *models.Server) error {
// 				panic("mock out the UpdateServerInfo method")
// 			},
// 			UpdateServerStatusFunc: func(server models.Server, status bool) error {
// 				panic("mock out the UpdateServerStatus method")
// 			},
// 		}
//
// 		// use mockedDatabase in code that requires database.Database
// 		// and then make assertions.
//
// 	}
type DatabaseMock struct {
	// CommitFunc mocks the Commit method.
	CommitFunc func(kills []models.Kill) error

	// CommitFlagCaptureFunc mocks the CommitFlagCapture method.
	CommitFlagCaptureFunc func(capture models.FlagCapture) error

	// CommitMapStatsFunc mocks the CommitMapStats method.
	CommitMapStatsFunc func(stats models.MapStats) error

	// CommitMapVotesFunc mocks the CommitMapVotes method.
	CommitMapVotesFunc func(votes models.MapVotes) error

	// CommitPlayerFunc mocks the CommitPlayer method.
	CommitPlayerFunc func(player *models.Player) error

	// InitDBFunc mocks the InitDB method.
	InitDBFunc func() error

	// UpdatePlayerInfoFunc mocks the UpdatePlayerInfo method.
	UpdatePlayerInfoFunc func(player *models.Player) error

	// UpdateServerInfoFunc mocks the UpdateServerInfo method.
	UpdateServerInfoFunc func(server *models.Server) error

	// UpdateServerStatusFunc mocks the UpdateServerStatus method.
	UpdateServerStatusFunc func(server models.Server, status bool) error

	// calls tracks calls to the methods.
	calls struct {
		// Commit holds details about calls to the Commit method.
		Commit []struct {
			// Kills is the kills argument value.
			Kills []models.Kill
		}
		// CommitFlagCapture holds details about calls to the CommitFlagCapture method.
		CommitFlagCapture []struct {
			// Capture is the capture argument value.
			Capture models.FlagCapture
		}
		// CommitMapStats holds details about calls to the CommitMapStats method.
		CommitMapStats []struct {
			// Stats is the stats argument value.
			Stats models.MapStats
		}
		// CommitMapVotes holds details about calls to the CommitMapVotes method.
		CommitMapVotes []struct {
			// Votes is the votes argument value.
			Votes models.MapVotes
		}
		// CommitPlayer holds details about calls to the CommitPlayer method.
		CommitPlayer []struct {
			// Player is the player argument value.
			Player *models.Player
		}
		// InitDB holds details about calls to the InitDB method.
		InitDB []struct {
		}
		// UpdatePlayerInfo holds details about calls to the UpdatePlayerInfo method.
		UpdatePlayerInfo []struct {
			// Player is the player argument value.
			Player *models.Player
		}
		// UpdateServerInfo holds details about calls to the UpdateServerInfo method.
		UpdateServerInfo []struct {
			// Server is the server argument value.
			Server *models.Server
		}
		// UpdateServerStatus holds details about calls to the UpdateServerStatus method.
		UpdateServerStatus []struct {
			// Server is the server argument value.
			Server models.Server
			// Status is the status argument value.
			Status bool
		}
	}
	lockCommit             sync.RWMutex
	lockCommitFlagCapture  sync.RWMutex
	lockCommitMapStats     sync.RWMutex
	lockCommitMapVotes     sync.RWMutex
	lockCommitPlayer       sync.RWMutex
	lockInitDB             sync.RWMutex
	lockUpdatePlayerInfo   sync.RWMutex
	lockUpdateServerInfo   sync.RWMutex
	lockUpdateServerStatus sync.RWMutex
}

// Commit calls CommitFunc.
func (mock *DatabaseMock) Commit(kills []models.Kill) error {
	if mock.CommitFunc == nil {
		panic("DatabaseMock.CommitFunc: method is nil but Database.Commit was just called")
	}
	callInfo := struct {
		Kills []models.Kill
	}{
		Kills: kills,
	}
	mock.lockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	mock.lockCommit.Unlock()
	return mock.CommitFunc(kills)
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//     len(mockedDatabase.CommitCalls())
func (mock *DatabaseMock) CommitCalls() []struct {
	Kills []models.Kill
} {
	var calls []struct {
		Kills []models.Kill
	}
	mock.lockCommit.RLock()
	calls = mock.calls.Commit
	mock.lockCommit.RUnlock()
	return calls
}

// CommitFlagCapture calls CommitFlagCaptureFunc.
func (mock *DatabaseMock) CommitFlagCapture(capture models.FlagCapture) error {
	if mock.CommitFlagCaptureFunc == nil {
		panic("DatabaseMock.CommitFlagCaptureFunc: method is nil but Database.CommitFlagCapture was just called")
	}
	callInfo := struct {
		Capture models.FlagCapture
	}{
		Capture: capture,
	}
	mock.lockCommitFlagCapture.Lock()
	mock.calls.CommitFlagCapture = append(mock.calls.CommitFlagCapture, callInfo)
	mock.lockCommitFlagCapture.Unlock()
	return mock.CommitFlagCaptureFunc(capture)
}

// CommitFlagCaptureCalls gets all the calls that were made to CommitFlagCapture.
// Check the length with:
//     len(mockedDatabase.CommitFlagCaptureCalls())
func (mock *DatabaseMock) CommitFlagCaptureCalls() []struct {
	Capture models.FlagCapture
} {
	var calls []struct {
		Capture models.FlagCapture
	}
	mock.lockCommitFlagCapture.RLock()
	calls = mock.calls.CommitFlagCapture
	mock.lockCommitFlagCapture.RUnlock()
	return calls
}

// CommitMapStats calls CommitMapStatsFunc.
func (mock *DatabaseMock) CommitMapStats(stats models.MapStats) error {
	if mock.CommitMapStatsFunc == nil {
		panic("DatabaseMock.CommitMapStatsFunc: method is nil but Database.CommitMapStats was just called")
	}
	callInfo := struct {
		Stats models.MapStats
	}{
		Stats: stats,
	}
	mock.lockCommitMapStats.Lock()
	mock.calls.CommitMapStats = append(mock.calls.CommitMapStats, callInfo)
	mock.lockCommitMapStats.Unlock()
	return mock.CommitMapStatsFunc(stats)
}

// CommitMapStatsCalls gets all the calls that were made to CommitMapStats.
// Check the length with:
//     len(mockedDatabase.CommitMapStatsCalls())
func (mock *DatabaseMock) CommitMapStatsCalls() []struct {
	Stats models.MapStats
} {
	var calls []struct {
		Stats models.MapStats
	}
	mock.lockCommitMapStats.RLock()
	calls = mock.calls.CommitMapStats
	mock.lockCommitMapStats.RUnlock()
	return calls
}

// CommitMapVotes calls CommitMapVotesFunc.
func (mock *DatabaseMock) CommitMapVotes(votes models.MapVotes) error {
	if mock.CommitMapVotesFunc == nil {
		panic("DatabaseMock.CommitMapVotesFunc: method is nil but Database.CommitMapVotes was just called")
	}
	callInfo := struct {
		Votes models.MapVotes
	}{
		Votes: votes,
	}
	mock.lockCommitMapVotes.Lock()
	mock.calls.CommitMapVotes = append(mock.calls.CommitMapVotes, callInfo)
	mock.lockCommitMapVotes.Unlock()
	return mock.CommitMapVotesFunc(votes)
}

// CommitMapVotesCalls gets all the calls that were made to CommitMapVotes.
// Check the length with:
//     len(mockedDatabase.CommitMapVotesCalls())
func (mock *DatabaseMock) CommitMapVotesCalls() []struct {
	Votes models.MapVotes
} {
	var calls []struct {
		Votes models.MapVotes
	}
	mock.lockCommitMapVotes.RLock()
	calls = mock.calls.CommitMapVotes
	mock.lockCommitMapVotes.RUnlock()
	return calls
}

// CommitPlayer calls CommitPlayerFunc.
func (mock *DatabaseMock) CommitPlayer(player *models.Player) error {
	if mock.CommitPlayerFunc == nil {
		panic("DatabaseMock.CommitPlayerFunc: method is nil but Database.CommitPlayer was just called")
	}
	callInfo := struct {
		Player *models.Player
	}{
		Player: player,
	}
	mock.lockCommitPlayer.Lock()
	mock.calls.CommitPlayer = append(mock.calls.CommitPlayer, callInfo)
	mock.lockCommitPlayer.Unlock()
	return mock.CommitPlayerFunc(player)
}

// CommitPlayerCalls gets all the calls that were made to CommitPlayer.
// Check the length with:
//     len(mockedDatabase.CommitPlayerCalls())
func (mock *DatabaseMock) CommitPlayerCalls() []struct {
	Player *models.Player
} {
	var calls []struct {
		Player *models.Player
	}
	mock.lockCommitPlayer.RLock()
	calls = mock.calls.CommitPlayer
	mock.lockCommitPlayer.RUnlock()
	return calls
}

// InitDB calls InitDBFunc.
func (mock *DatabaseMock) InitDB() error {
	if mock.InitDBFunc == nil {
		panic("DatabaseMock.InitDBFunc: method is nil but Database.InitDB was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInitDB.Lock()
	mock.calls.InitDB = append(mock.calls.InitDB, callInfo)
	mock.lockInitDB.Unlock()
	return mock.InitDBFunc()
}

// InitDBCalls gets all the calls that were made to InitDB.
// Check the length with:
//     len(mockedDatabase.InitDBCalls())
func (mock *DatabaseMock) InitDBCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInitDB.RLock()
	calls = mock.calls.InitDB
	mock.lockInitDB.RUnlock()
	return calls
}

// UpdatePlayerInfo calls UpdatePlayerInfoFunc.
func (mock *DatabaseMock) UpdatePlayerInfo(player *models.Player) error {
	if mock.UpdatePlayerInfoFunc == nil {
		panic("DatabaseMock.UpdatePlayerInfoFunc: method is nil but Database.UpdatePlayerInfo was just called")
	}
	callInfo := struct {
		Player *models.Player
	}{
		Player: player,
	}
	mock.lockUpdatePlayerInfo.Lock()
	mock.calls.UpdatePlayerInfo = append(mock.calls.UpdatePlayerInfo, callInfo)
	mock.lockUpdatePlayerInfo.Unlock()
	return mock.UpdatePlayerInfoFunc(player)
}

// UpdatePlayerInfoCalls gets all the calls that were made to UpdatePlayerInfo.
// Check the length with:
//     len(mockedDatabase.UpdatePlayerInfoCalls())
func (mock *DatabaseMock) UpdatePlayerInfoCalls() []struct {
	Player *models.Player
} {
	var calls []struct {
		Player *models.Player
	}
	mock.lockUpdatePlayerInfo.RLock()
	calls = mock.calls.UpdatePlayerInfo
	mock.lockUpdatePlayerInfo.RUnlock()
	return calls
}

// UpdateServerInfo calls UpdateServerInfoFunc.
func (mock *DatabaseMock) UpdateServerInfo(server *models.Server) error {
	if mock.UpdateServerInfoFunc == nil {
		panic("DatabaseMock.UpdateServerInfoFunc: method is nil but Database.UpdateServerInfo was just called")
	}
	callInfo := struct {
		Server *models.Server
	}{
		Server: server,
	}
	mock.lockUpdateServerInfo.Lock()
	mock.calls.UpdateServerInfo = append(mock.calls.UpdateServerInfo, callInfo)
	mock.lockUpdateServerInfo.Unlock()
	return mock.UpdateServerInfoFunc(server)
}

// UpdateServerInfoCalls gets all the calls that were made to UpdateServerInfo.
// Check the length with:
//     len(mockedDatabase.UpdateServerInfoCalls())
func (mock *DatabaseMock) UpdateServerInfoCalls() []struct {
	Server *models.Server
} {
	var calls []struct {
		Server *models.Server
	}
	mock.lockUpdateServerInfo.RLock()
	calls = mock.calls.UpdateServerInfo
	mock.lockUpdateServerInfo.RUnlock()
	return calls
}

// UpdateServerStatus calls UpdateServerStatusFunc.
func (mock *DatabaseMock) UpdateServerStatus(server models.Server, status bool) error {
	if mock.UpdateServerStatusFunc == nil {
		panic("DatabaseMock.UpdateServerStatusFunc: method is nil but Database.UpdateServerStatus was just called")
	}
	callInfo := struct {
		Server models.Server
		Status bool
	}{
		Server: server,
		Status: status,
	}
	mock.lockUpdateServerStatus.Lock()
	mock.calls.UpdateServerStatus = append(mock.calls.UpdateServerStatus, callInfo)
	mock.lockUpdateServerStatus.Unlock()
	return mock.UpdateServerStatusFunc(server, status)
}

// UpdateServerStatusCalls gets all the calls that were made to UpdateServerStatus.
// Check the length with:
//     len(mockedDatabase.UpdateServerStatusCalls())
func (mock *DatabaseMock) UpdateServerStatusCalls() []struct {
	Server models.Server
	Status bool
} {
	var calls []struct {
		Server models.Server
		Status bool
	}
	mock.lockUpdateServerStatus.RLock()
	calls = mock.calls.UpdateServerStatus
	mock.lockUpdateServerStatus.RUnlock()
	return calls
}