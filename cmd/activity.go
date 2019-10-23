/**
 *2848869
 */

package cmd

type Activity struct {
	//file?
	Type    string
	Comment string
}

func newActivity(activityType string, comment string) Activity {
	return Activity{activityType, comment}
}
