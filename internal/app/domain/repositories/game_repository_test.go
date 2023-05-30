package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abdil1234/test-golang/config"
	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/commons/appcontext"
	"github.com/abdil1234/test-golang/internal/app/domain/entities"
	"github.com/abdil1234/test-golang/internal/app/domain/repositories/models"
	mockRepo "github.com/abdil1234/test-golang/mocks/domain/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GameRepositorySuite struct {
	suite.Suite
	mockRepo *mockRepo.MockIGameRepository

	MockDatabase sqlmock.Sqlmock
	Option       commons.Option

	GameRepository IGameRepository
	games          []entities.GameEntity
}

const gamesJson = `[
  {
    "id": 1,
    "mdate": "12 June 2020",
    "stadium": "GBK",
    "team1": "INA",
    "team2": "MAS"
  }
]`

func TestGameRepositorySuite(t *testing.T) {

	suite.Run(t, new(GameRepositorySuite))
}

func (ts *GameRepositorySuite) SetupSuite() {

	mockCtrl := gomock.NewController(ts.T())

	defer mockCtrl.Finish()

	var err error
	var dbClient *sql.DB

	cfg := config.GetConfig()
	app := appcontext.NewAppContext(cfg)

	dbClient, ts.MockDatabase, err = sqlmock.New()
	if err != nil {
		_ = fmt.Errorf("sqlmock error: %v", err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: dbClient, SkipInitializeWithVersion: true}))

	if err != nil {
		_ = fmt.Errorf("gorm error: %v", err)
	}

	opt := commons.Option{
		AppCtx:   app,
		Config:   cfg,
		Database: db,
	}

	ts.Option = opt

	ts.GameRepository = NewGameRepository(opt)

	err = json.Unmarshal([]byte(gamesJson), &ts.games)

	log.Println(ts.games)
	if err != nil {
		_ = fmt.Errorf("unmarshall error: %v", err)
	}
}

func dateToUnix(date string) *int64 {
	mDate, _ := time.Parse("02 January 2006", date)
	unixDate := mDate.Unix()
	return &unixDate
}

func (ts *GameRepositorySuite) TestGameRepository_Insert() {

	type args struct {
		game models.Game
	}

	tests := []struct {
		name    string
		args    args
		mock    func(game models.Game)
		wantErr bool
		want    error
	}{
		{
			name: "Insert case",
			args: args{
				game: models.Game{
					Id:      1,
					Mdate:   dateToUnix("01 June 1998"),
					Stadium: "GBK",
					Team1:   "INA",
					Team2:   "MAS",
				},
			},
			mock: func(game models.Game) {
				ts.MockDatabase.ExpectBegin()
				ts.MockDatabase.ExpectExec("INSERT INTO `game`").WillReturnResult(sqlmock.NewResult(1, 1))
				ts.MockDatabase.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args.game)

			got := ts.GameRepository.Insert(tt.args.game)

			if tt.wantErr == false {
				assert.Nil(t, got)
			}
			assert.Equalf(t, tt.want, got, "Insert()")
		})
	}
}

func (ts *GameRepositorySuite) TestGameRepository_Gets() {

	type args struct {
		game models.Game
	}
	mDate, _ := time.Parse("02 January 2006", "12 June 2020")
	mDateUnix := mDate.Unix()

	tests := []struct {
		name      string
		args      args
		mock      func(game []entities.GameEntity)
		wantErr   bool
		wantGames []models.Game
	}{
		{
			name: "Insert case",
			args: args{
				game: models.Game{
					Id:      1,
					Mdate:   &mDateUnix,
					Stadium: "GBK",
					Team1:   "INA",
					Team2:   "MAS",
				},
			},
			mock: func(games []entities.GameEntity) {
				rows := sqlmock.NewRows([]string{"id", "stadium", "team1", "team2"})
				for _, game := range games {
					rows = rows.AddRow(
						game.Id, game.Stadium, game.Team1, game.Team2,
					)
				}

				ts.MockDatabase.ExpectQuery("SELECT (.+) FROM `game`").WillReturnRows(rows)
			},
			wantGames: []models.Game{
				{
					Id:      1,
					Stadium: "GBK",
					Team1:   "INA",
					Team2:   "MAS",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock(ts.games)
			gotGames, err := ts.GameRepository.Gets()
			if tt.wantErr == false {
				assert.Nil(t, err)
			}
			if !reflect.DeepEqual(gotGames, tt.wantGames) {
				t.Errorf("Gets() gotGames = %v, want %v", gotGames, tt.wantGames)
			}
		})
	}
}

func int64Ptr(value int64) *int64 {
	return &value
}
