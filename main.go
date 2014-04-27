// simple notification script to use pushover from nagios
//
//     notify-by-pushover -H=$HOSTNAME$ -s=$HOSTSTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$HOSTOUTPUT$ -u=$CONTACTEMAIL$
//     notify-by-pushover -H=$HOSTNAME$ -S=$SERVICEDESC$ -s=$SERVICESTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$SERVICEOUTPUT$ -u=$CONTACTEMAIL$
//
package main

import (
	"bitbucket.org/kisom/gopush/pushover"
	"flag"
	"fmt"
)

func main() {
	// note, that variables are pointers
	host := flag.String("H", "", "the host to alert for")
	service := flag.String("S", "", "the service to alert for")
	state := flag.String("s", "", "the state")
	key := flag.String("a", "", "the pushover auth token")
	alert_type := flag.String("t", "", "the alert type")
	recipient := flag.String("u", "", "pushover recipient ID to send to")
	description := flag.String("d", "", "description of the alert")

	flag.Parse()

	identity := pushover.Authenticate(*key, *recipient)

	var label, message string
	if *service != "" {
		label = fmt.Sprintf("%s/%s", *host, *service)
		message = fmt.Sprintf("%s on %s is %s\n%s", *service, *host, *state,
			*description)
	} else {
		label = *host
		message = fmt.Sprintf("%s is %s\n%s", *host, *state, *description)
	}

	pushover.Notify_titled(identity, message, fmt.Sprintf("%s: %s", *alert_type, label))
}
