package intchess

type ChessGame struct {
	Id          int64           `json:"game_id"`
	CreatedAt   NullTime        `json:"commenced_at"`
	FinishedAt  NullTime        `json:"finished_at"`
	Status      string          `json:"game_status"`
	BoardStatus [][][]byte      `json:"board_status"`
	GameMoves   []GameMove      `json:"game_moves"`
	WhitePlayer *User           `json:"player1"`
	BlackPlayer *User           `json:"player2"`
	Winner      *UserRankChange `json:"winner",omitempty`
	Loser       *UserRankChange `json:"loser","omitempty`
}

type UserRankChange struct {
	Id         int64    `json:"-"`
	UserId     int64    `json:"user_id"`
	RankChange int      `json:"rank_change"`
	CreatedAt  NullTime `json:"-"`
}

type GameMove struct {
	Id        int64    `json:"id"`
	Move      []byte   `json:"move"`
	CreatedAt NullTime `json:"created_at"`
}

func NewGame(white *User, black *User) ChessGame {
	g := Game{
		Status:      "new",
		BoardStatus: NewBoardStatus(),
		WhitePlayer: white,
		BlackPlayer: black,
	}
	return g
}
