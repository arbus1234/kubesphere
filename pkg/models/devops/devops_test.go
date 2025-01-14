package devops

import "testing"

func Test_parseCronJobTime(t *testing.T) {
	type Except struct {
		Last string
		Next string
	}

	Items := []struct {
		Input    string
		Expected Except
	}{
		{"上次运行的时间 Tuesday, September 10, 2019 8:59:09 AM UTC; 下次运行的时间 Tuesday, September 10, 2019 9:14:09 AM UTC.", Except{Last: "2019-09-10T08:59:09Z", Next: "2019-09-10T09:14:09Z"}},
		{"上次运行的时间 Thursday, January 3, 2019 11:56:30 PM UTC; 下次运行的时间 Friday, January 3, 2020 12:11:30 AM UTC.", Except{Last: "2019-01-03T23:56:30Z", Next: "2020-01-03T00:11:30Z"}},
		{"上次运行的时间 Tuesday, September 10, 2019 8:41:34 AM UTC; 下次运行的时间 Tuesday, September 10, 2019 9:41:34 AM UTC.", Except{Last: "2019-09-10T08:41:34Z", Next: "2019-09-10T09:41:34Z"}},
		{"上次运行的时间 Tuesday, September 10, 2019 9:15:26 AM UTC; 下次运行的时间 Tuesday, September 10, 2019 10:03:26 AM UTC.", Except{Last: "2019-09-10T09:15:26Z", Next: "2019-09-10T10:03:26Z"}},
	}

	for _, item := range Items {
		last, next, err := parseCronJobTime(item.Input)
		if err != nil {
			t.Fatalf("should not get error %+v", err)
		}

		if last != item.Expected.Last {
			t.Errorf("got %#v, expected %#v", last, item.Expected.Last)
		}

		if next != item.Expected.Next {
			t.Errorf("got %#v, expected %#v", next, item.Expected.Next)
		}

	}
}
