package processor

//func TestRanker2(t *testing.T) {
//	matchlocations := map[string][][]int{}
//	matchlocations["cat"] = [][]int{
//		{0, 4},
//	}
//
//	results := []*fileJob{}
//
//	for i := 0; i < 300; i++ {
//		results = append(results, &fileJob{
//			Score:          0,
//			MatchLocations: matchlocations,
//		})
//	}
//
//	for i := 0; i < 9_700; i++ {
//		results = append(results, &fileJob{
//			Score:          0,
//			MatchLocations: map[string][][]int{},
//		})
//	}
//
//
//
//	ranked := rankResultsTFIDF(results)
//	sortResults(ranked)
//
//	if len(ranked) != 1 {
//		t.Error("Should be one results")
//	}
//
//	if ranked[0].Score <= 0 {
//		t.Error("Score should be greater than 0")
//	}
//}

//
//func TestRankResultsLocation(t *testing.T) {
//	results := []*fileJob{
//		{
//			Filename: "test.go",
//			Location: "/this/matches/something/test.go",
//			Score:    0,
//		},
//	}
//	ranked := rankResultsLocation([][]byte{[]byte("something")}, results)
//
//	if ranked[0].Score == 0 {
//		t.Error("Expect rank to be > 0 got", ranked[0].Score)
//	}
//}
//
//func TestRankResultsLocationScoreCheck(t *testing.T) {
//	results := []*fileJob{
//		{
//			Filename: "test1.go",
//			Location: "/this/matches/something/test1.go",
//			Score:    0,
//		},
//		{
//			Filename: "test2.go",
//			Location: "/this/matches/something/test2.go",
//			Score:    0,
//		},
//	}
//	ranked := rankResultsLocation([][]byte{[]byte("something"), []byte("test1")}, results)
//
//	if ranked[0].Score <= ranked[1].Score {
//		t.Error("Expect first to get higher match", ranked[0].Score, ranked[1].Score)
//	}
//}
