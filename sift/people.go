package sift

import (
	"context"

	"github.com/bernardolm/octo-batch/debug"
)

type person struct {
	// CustomPictureUrl   string
	// DirectoryId        string
	// Email              string
	// IsTeamLeader       bool
	// LastName           string
	// OfficialPictureUrl string
	// PictureUrl         string
	// ReportingPath      []string
	DirectReportCount int
	FirstName         string
	Id                string
	TeamLeaderId      string
	TotalReportCount  int
}

// resp, err := getClient().Get(
//		fmt.Sprintf("%s/people/aacuna5f@tychoengineering.com",
//		apiURL))
// resp, err := getClient().Get(
//		fmt.Sprintf("%s/search/people?department=Human%%20Resources&firstName=Pris",
//		apiURL))
// resp, err := getClient().Get(
//		fmt.Sprintf("%s/search/people?department=Human%%20Resources",
//		apiURL))
// resp, err := getClient().Get(
//		fmt.Sprintf("%s/search/people?pageSize=100",
//		apiURL))

func PeopleListAll(ctx context.Context) ([]person, error) {
	return []person{}, nil

	pg := page{
		Page:     1,
		PageSize: 100,
	}

	var people []person
	var stop bool
	var total int

	for keepRunning := true; keepRunning; keepRunning = !stop {
		res, err := paginate(ctx, "search/people", &pg)
		if err != nil {
			return nil, err
		}

		if total == 0 {
			total = res.Meta.TotalLength
		}

		result := struct {
			Data []person
		}{}

		if err := json.Unmarshal(res.Data, &result); err != nil {
			return nil, err
		}

		people = append(people, result.Data...)

		pg.Page = res.Meta.Pages.Next

		if len(people) >= total {
			stop = true
		}
	}

	debug.Print("PeopleListAll.people", len(people))

	return people, nil
}
