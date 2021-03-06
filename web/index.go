/*
Copyright 2013 CoreOS Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package web

// This file was generated from index.html.

var index_html string = "<html>\n<head>\n<title>etcd Web Interface</title>\n<script type=\"text/javascript\" src=\"//ajax.googleapis.com/ajax/libs/jquery/1.10.1/jquery.min.js\"></script>\n<script type=\"text/javascript\">\n    $(function() {\n\n    var conn;\n    var content = $(\"#content\");\n\n    function update(response) {\n        // if set\n        if (response.action == \"SET\") {\n\n            if (response.expiration > \"1970\") {\n                t = response.key + \"=\" + response.value\n                        + \"  \" + response.expiration\n            } else {\n                t = response.key + \"=\" + response.value\n            }\n\n            id = response.key.replace(new RegExp(\"/\", 'g'), \"\\\\/\");\n\n            if ($(\"#store_\" + id).length == 0) {\n                if (response.expiration > \"1970\") {\n                    t = response.key + \"=\" + response.value\n                        + \"  \" + response.expiration\n                } else {\n                    t = response.key + \"=\" + response.value\n                }\n\n                var e = $('<div id=\"store_' + response.key + '\"/>')\n                    .text(t)\n                e.appendTo(content)\n            }\n            else {\n\n                $(\"#store_\" + id)\n                    .text(t)\n            }\n        }\n        // if delete\n        else if (response.action == \"DELETE\") {\n            id = response.key.replace(new RegExp(\"/\", 'g'), \"\\\\/\");\n\n            $(\"#store_\" + id).remove()\n        }\n    }\n\n\n    if (window[\"WebSocket\"]) {\n        conn = new WebSocket(\"ws://{{.Address}}/ws\");\n        conn.onclose = function(evt) {\n\n        }\n        conn.onmessage = function(evt) {\n            var response = JSON.parse(evt.data)\n            update(response)\n        }\n    } else {\n        appendLog($(\"<div><b>Your browser does not support WebSockets.</b></div>\"))\n    }\n    });\n</script>\n</head>\n<body>\n    <div id=\"leader\">Leader: {{.Leader}}</div>\n    <div id=\"content\"></div>\n</body>\n</html>\n"
