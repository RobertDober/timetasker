Timetasker CLI
==============

Provides a nice command line interface to submit time sheets using [Time
Task](https://timetask.com). This command will submit a single time entry.

The project improves upon a [Chrome
extension](https://github.com/teddywing/chrome-timetasker) that auto-fills the
time sheet form on the website.


## Usage
This will submit a time entry for the "example" project on the current day with
a duration of 7 hours and an empty description:

	$ timetasker --project example

Here we set a custom time of 4.5 hours:

	$ timetasker --project example --time 4.5

Now we specify a date and add a description:

	$ timetasker --project example --date 2017-05-31 --description "Worked on Timetasker"


## Configuration
Timetasker relies on a configuration file in order to work properly.

If this is your first time running Timetasker, use this command to generate a
skeleton config:

	$ timetasker --write-config

This will generate a `$HOME/.config/timetasker/config.toml` file that looks like
this:

	[auth]
	username = ""
	password_cmd = ""


	[profile]
	person_id = # ADD PERSON ID


	[projects.example]
	client =    # ADD CLIENT ID
	project =   # ADD PROJECT ID
	module =    # ADD MODULE ID
	task = 0
	work_type = # ADD WORK TYPE ID
	billable = true

Fill in the `username` with your TimeTask username. The `password_cmd` should be
a shell command that will output your TimeTask password to STDOUT.

Notice the `[projects.example]` line? That's a project alias. If we, for
instance, changed it to `[projects.my-cool-project]`, we could post a time entry
to that project like this:

	$ timetasker --project my-cool-project

You say you have more than one project? No problem, just copy-paste the entire
`[projects.example]` section and give it a new name.

To fill in the other configuration options, we're going to have to take a trip
to the TimeTask website (relax, we won't be using it much after this).

1. Visit `https://*.timetask.com/time/add/`
2. Fill in a single entry, but don't submit it yet
3. Open the Network console in your browser's developer tools
4. Submit your time entry
5. Open the POST request to `https://*.timetask.com/index.php`
6. Take a look at the form data of the request

You should see something like this:

	module:time
	action:submitmultipletime
	f_entryIndexes:0
	f_personID0:111111
	f_clientID0:22222
	f_projectID0:333333
	f_moduleID0:444444
	f_taskID0:0
	f_worktypeID0:555555
	f_date0:04/06/17
	f_time0:7
	f_billable0:t
	f_description0:

Copy the numbers into their corresponding fields in your `config.toml` file.

Once you have a complete config file, you should be all set to start posting
time entries!


## Install
Visit the [releases][3] page, download the version corresponding to your
platform, and put the resulting `timetasker` binary on your PATH.

To install from source, use:

	$ go install github.com/teddywing/timetasker


## License
Copyright © 2017 Teddy Wing. Licensed under the GNU GPLv3+ (see the included
COPYING file).


[3]: https://github.com/teddywing/timetasker/releases
