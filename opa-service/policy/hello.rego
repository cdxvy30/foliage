package hello

import rego.v1

msg if {
  msg := data.greetings[input.lang]
}
