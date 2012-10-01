Nolo
========

Nolo is a simple, flexible metric API to multiple metric collectors

Most of the time, all a metrics system really needs is an identifier
and a value. So the simplest possible plugin should just do that:

    % ./app_plugin
    connections_open 216
    requests 153

Of course many collector systems support extra metadata, so it should
be easy to include this as well:

    % ./app_plugin
    # host=darkstar
    connections_open 216 time=1200928930 type=uint16
    requests 153 release=2.24.2.4271 type=uint8

This is a shorthand way of writing:

    % ./app_plugin
    connections_open 216 time=1200928930 type=uint16 host=darkstar
    requests 153 release=2.24.2.4271 type=uint8 host=darkstar

Collector adaptors are free to use these metadata pairs as they like.
Some will ignore them entirely, others will only use the ones that are
meaningful, and a some will store them all. Refer to the individual
adapter for how your system works.

Installation
------------

Extract the repo into the directory of your choice:

    git clone https://github.com/josephholsten/nolo /opt/nolo

Adapter scripts will be found in `bin/`, and plugins will be found
in `plugins/`.

Ganglia
-------

Assuming you've installed nolo into `/opt/nolo`, you should be able to
run the nolo `load` plugin and pass the input into gmetric by running:

    /opt/nolo/bin/nolo-gmetric /opt/nolo/plugins/load

You can run this plugin regularly by entering adding it to your crontab:

    * * * * * /opt/nolo/bin/nolo-gmetric /opt/nolo/plugins/load

Please note: the ganglia adapter does not yet support any metadata
flags, so it will only pass along the identifier and value.

Wanted Adapters
---------------

- cacti
- munin
- collectd
- scout
- server density

About the name
--------------

I ran `grep '^....$' < /usr/share/dict/words` and skimmed. Nolo
reminded me of YOLO and is the first part of "nolo contendere", a plea
of no contest. Which is how I feel when I have to choose between
plugin formats.

See also
--------

- [Graphite](http://graphite.wikidot.com): whose plain text protocol is impossibly simple yet impressively flexible
- [Nagios](http://www.nagios.org): whose plugin format is the de facto standard in alerting

LICENSE
-------

Nolo is Copyright (c) 2012 [Joseph Anthony Pasquale
Holsten](http://josephholsten.com) and distributed under the ISC
license. See the `LICENSE` file for more info.
