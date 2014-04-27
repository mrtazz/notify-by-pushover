# notify-by-pushover

## Overview
Simple script to notify about nagios alerts via pushover

## Usage
The script can be called standalone like this:
```
% ./notify-by-pushover --help
Usage of ./notify-by-pushover:
  -H="": the host to alert for
  -S="": the service to alert for
  -a="": the pushover auth token
  -d="": description of the alert
  -s="": the state
  -t="": the alert type
  -u="": pushover recipient ID to send to
```
But you probably want to use it in Nagios. For this define commands like this:
```
define command{
  command_name	notify-host-pushover
  command_line  $USER1$/notify-by-pushover -H="$HOSTNAME$" -s="$HOSTSTATE$" -a="pushoverKEY" -t="$NOTIFICATIONTYPE$" -d="$HOSTOUTPUT$" -T="$LONGDATETIME$" -u="$CONTACTEMAIL$"
}

define command{
  command_name	notify-service-pushover
  command_line  $USER1$/notify-by-pushover -H="$HOSTNAME$" -S="$SERVICEDESC$" -s="$SERVICESTATE$" -a="pushoverKEY" -t="$NOTIFICATIONTYPE$" -d="$SERVICEOUTPUT$" -T="$LONGDATETIME$" -u="$CONTACTEMAIL$"
}
```
And then set up a contact to use the commands for notifications like this:
```
define contact{
        contact_name                   pushover
        use                            generic-contact
        alias                          pushover contact
        email                          $PUSHOVER_USER_ID
        service_notification_commands  notify-service-pushover
        host_notification_commands     notify-host-pushover
        }
```

## How to contribute
1. Fork the repo
2. Hack away
3. Push the branch up to GitHub
4. Send a pull request

