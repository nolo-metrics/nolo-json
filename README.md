Nolo-JSON
=========

Nolo is a simple, flexible metric API to multiple metric collectors. For
more information about nolo, see https://github.com/nolo-metrics/nolo-core

This is the JSON output adapter. It is designed to provide a simple way
to write nolo adapters without having to write the parsing yourself.

Assuming you have a plugin that serves the following metrics:

    % ./app_plugin
    connections_open 216 time=1200928930 type=uint16 host=darkstar
    requests 153 release=2.24.2.4271 type=uint8 host=darkstar

Then `nolo-json` lets you access the data in json format:

    % nolo-json ./app_plugin
    {
      "app-plugin": [
        {
          "identifier": "connections_open",
          "value": "216",
          "time": "1200928930",
          "type": "uint16",
          "host": "darkstar"
        },
        {
          "identifier": "requests",
          "value": "153",
          "release": "2.24.2.4271",
          "type": "uint8",
          "host": "darkstar"
        }
      ]
    }

LICENSE
-------

Nolo-JSON is Copyright (c) 2012 [Joseph Anthony Pasquale
Holsten](http://josephholsten.com) and distributed under the ISC
license. See the `LICENSE` file for more info.
