package main

// Input string - can be a real value or can be blank
// Output to another value represented by the string e.g. int, string, date, float.

func conv (s string, t string) *interface{} {
	if s != nil {
		switch t {
			case "INT":
			case "STR":
		}
	} else {
		return nil
	}
}
	
