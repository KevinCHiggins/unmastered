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
	announcements = []Announcement{
		{1, "First Post", "And we're live!", time.Date(2022, time.April, 0, 0, 0, 0, 0, myDate)},
		{2, "More Content", "I've added some content from Dan Soro's trio project, check ittttt!", time.Now()},
	}

	t, _ := time.Parse("02/01/06", "29/06/22")
	projects = []Project{
		{1, "Dan Soro Trio", t},
		{2, "Farpark", t},
		{2, "Denatured", t},
	}

	t, _ = time.Parse("02/01/06", "29/06/22")
	t2, _ := time.Parse("02/01/06", "02/07/22")
	contributors = []Contributor{
		{1, "Kevin Higgins", "kevin@kevinhiggins.dev", "spoon", t},
		{2, "Dan Soro", "dan.soro@example.com", "fork", t2},
		{2, "Tommy Gray", "info@tommygraycomposer.com", "knife", t2},
	}

	// yes, this will be in the DB too one day
	hardcodedMediaUploadTypes = []MediaUploadType{
		{1, "score", "pdf"},
		{2, "sample pack", "zip"},
		{3, "multitrack", "wav"},
	}

	t, _ = time.Parse("02/01/06", "30/06/22")
	hardcodedMediaUploads = []MediaUpload{
		{1, 1, 1, 1, "I'll Be Seeing You (Arrangement)", "Guitar and bass arrangement by Kevin", true, t},
		{2, 2, 1, 3, "Nardub", "A dubby, layered take on the jam session staple 'Nardis'", true, t},
	}
	hardcodedUploadedFiles = []UploadedFile{
		{"1/hashy.pdf", 1},
	}
}
