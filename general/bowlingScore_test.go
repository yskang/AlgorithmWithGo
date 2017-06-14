package general

import "testing"

func TestOneInputReturn0(t *testing.T) {
	expectScore := 0

	score := BowlingScore("1")

	if expectScore != score {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestTowInputNotSpareReturnSimpleSum(t *testing.T) {
	expectScore := 7

	score := BowlingScore("34")

	if expectScore != score {
		t.Error("Wrong answer! expect score is ", expectScore, "but return ", score)
	}
}

func TestThreeInputNotStrikeReturnLastScore(t *testing.T) {
	expectScore := 1

	score := BowlingScore("013")

	if expectScore != score {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestTwoInputSpareReturnPreviousScore(t *testing.T) {
	expectScore := 0

	score := BowlingScore("1/")

	if expectScore != score {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestOneAfterSpareReturnScore(t *testing.T) {
	expectScore := 19

	score := BowlingScore("0/9")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestTwoAfterSpareReturnScore(t *testing.T) {
	expectScore := 28

	score := BowlingScore("0/90")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestThreeAfterSpareReturnScore(t *testing.T) {
	expectScore := 28

	score := BowlingScore("0/909")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestFourAfterSpareReturnScore(t *testing.T) {
	expectScore := 36

	score := BowlingScore("0/9035")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestThreeStrikeReturnTenPlusBonusTwenty(t *testing.T) {
	expectScore := 30

	score := BowlingScore("XXX")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}

func TestFreeScore(t *testing.T) {
	expectScore := 60

	score := BowlingScore("XXXX")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}

	expectScore = 104

	score = BowlingScore("XXXX4/")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}


	expectScore = 133

	score = BowlingScore("XXXX4/327/018/")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}

	expectScore = 300

	score = BowlingScore("XXXXXXXXXXXX")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}

	expectScore = 145

	score = BowlingScore("5/5/5/5/5/5/5/5/5/5/0")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}

	expectScore = 150

	score = BowlingScore("908/72X9/09X0/90XX8")

	if score != expectScore {
		t.Error("Wrong answer! expect score is", expectScore, "but return ", score)
	}
}
