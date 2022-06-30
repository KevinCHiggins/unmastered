package model

import "time"

func test() {
	print(" ")
}

func LoadTestData() {

	myDate, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err) // this is  hardcoded dummy data so no point handling error
	}
	hardcodedAnnouncements = []Announcement{
		{1, "First Post", "And we're live!", time.Date(2022, time.April, 0, 0, 0, 0, 0, myDate)},
		{2, "More Content", "I've added some content from my duo project with Dan Soro, check ittttt!", time.Now()},
	}

	t, _ := time.Parse("02/01/06", "29/06/22")
	hardcodedProjects = []Project{{1, "Kev and Soro", t}}

	hardcodedMediaUploadTypes = []MediaUploadType{
		{1, "score", "pdf"},
		{2, "samplepack", "zip"},
		{3, "multitrack", "wav"},
	}
	t, _ = time.Parse("02/01/06", "30/06/22")
	hardcodedMediaUploads = []MediaUpload{
		{1, 1, 1, 1, "I'll Be Seeing You (Arrangement)", "Guitar and bass arrangement by Kevin", true, t},
	}
	hardcodedUploadedFiles = []UploadedFile{
		{"1/hashy.pdf", 1},
	}
}
