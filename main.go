// Name:Dona Maria
// NetID: dmari21
// Description: creates a password manager that handles a list of entries, each with the required project information

package main

import (
	"fmt"
	"os"
)

// Global variables are allowed (and encouraged) for this project.

// a struct called Entry to store the username and password
type Entry struct {
  Username string
  Password string
}

// Declares EntrySlice as a slice of Entry
type EntrySlice []Entry

var passwordMap map[string]EntrySlice

// _______________________________________________________________________
// initialize before main()
// initializes the map and calls the pmRead()
// _______________________________________________________________________
func init(){
  passwordMap = make(map[string]EntrySlice)
  pmRead()
}

// _______________________________________________________________________
// This function finds the matching entry slice, if it exists, and returns it.
// takes in the site and EntrySlice as the parameter
// returns the matching EntrySlice and a bool
// _______________________________________________________________________
func findEntrySlice(site string) (EntrySlice, bool) {
	entry, found := passwordMap[site]
	return entry, found
}

// _______________________________________________________________________
// set the entrySlice for site
// takes in the site and EntrySlice as the parameter
// does not return anything
// _______________________________________________________________________
func setEntrySlice(site string, entrySlice EntrySlice) {
  passwordMap[site] = entrySlice
}
  

// _______________________________________________________________________
// to find the matching entry, if it exists, and returns it.
// takes in user and EntrySlice as the parameter
// returns an interger and a bool
// _______________________________________________________________________
func find(user string, entrySlice EntrySlice) (int, bool) {
  for i, entry := range entrySlice {
    if entry.Username == user {
      return i, true
    }
  }
  return 0, false
}

// _______________________________________________________________________
// print the list in columns
// _______________________________________________________________________
func pmList() {
  for site, entrySlice := range passwordMap {
    for _, entry := range entrySlice {
      fmt.Printf("%-20s %-20s %-20s\n", site, entry.Username, entry.Password)
    }
  }
}

// _______________________________________________________________________
//	add an entry if the site, user is not already found
// takes in site, username and password as the arguments and returns nothing
// _______________________________________________________________________
func pmAdd(site, user, password string) {
  entrySlice, found := findEntrySlice(site)

  if !found {
    // If the site is not found, initialize the entrySlice with an empty slice
    entrySlice = EntrySlice{}
  }

  _, find := find(user, entrySlice)

  if !find {
    // If the user is not found, append a new entry
    entry := Entry{Username: user, Password: password}
    entrySlice = append(entrySlice, entry)
  } else {
    // If the user is found, update the password
    fmt.Printf("add: duplicate entry")
  }

  // Set the entrySlice for the site
  setEntrySlice(site, entrySlice)
  pmWrite()
}

// _______________________________________________________________________
// This function remove by site and user
// the function takes site and user as paremeter and returns nothing
// _______________________________________________________________________
func pmRemove(site, user string) {
  entrySlice, found := findEntrySlice(site)
  if !found {
    fmt.Printf("remove: site not found")
    return
  }
  i, found := find(user, entrySlice)
  if !found {
    fmt.Printf("No such user: %s\n", user)
    return
  }
  entrySlice = append(entrySlice[:i], entrySlice[i+1:]...)

  setEntrySlice(site, entrySlice)
  pmWrite()
}

// _______________________________________________________________________
// remove the whole site if there is a single user at that site
// _______________________________________________________________________
func pmRemoveSite(site string) {
  entrySlice, found := findEntrySlice(site)
  if !found {
    fmt.Printf("remove: site not found")
    return
  }
  if len(entrySlice) == 1 {
    delete(passwordMap, site)
    pmWrite()
    return
  }
  fmt.Printf("Attempted to remove multiple users.")
}

// _______________________________________________________________________
// to read from the passwordVault file
// _______________________________________________________________________
func pmRead() {
	filename := "passwordVault"
	data := make(map[string]interface{})

	file, _ := os.Open(filename)

	var site, info, username, password string
	for {
		_, err := fmt.Fscanf(file, `"%s": "%s",`, &site, &info)
		if err != nil {
			break
		}
		data[site] = Entry{Username: username, Password: password}
	}
}

// _______________________________________________________________________
// to write the entries to the passwordVault file
// _______________________________________________________________________
func pmWrite() {
  filename := "passwordVault"
  file, _ := os.Create(filename)
  for site, entrySlice := range passwordMap {
    for _, entry := range entrySlice {
      fmt.Fprintf(file, `"%s": `, site)
      fmt.Fprintf(file, `"%s"`, entry.Username)
      fmt.Fprintf(file, `"%s"`, entry.Password)
      fmt.Fprintf(file, `%s`, "\n")
    }
  }
  file.Close()
}

// _______________________________________________________________________
// do forever loop reading the following commands
//
//	  l
//	  a s u p
//	  r s
//	  r s u
//	  x
//	where l,a,r,x are list, add, remove, and exit
//	and s,u,p are site, user, and password
//
// _______________________________________________________________________
func loop() {
  for {
    var command, site, user, password string
    fmt.Scanln(&command, &site, &user, &password)
    switch command {
      case "l":
      pmList()
      case "a":
      pmAdd(site, user, password)
      case "r":
      if len(user) == 0 {
        pmRemoveSite(site)
      } else {
          pmRemove(site, user)
      }
      case "x":
      return
      default:
      fmt.Printf("Unknown command: %s\n", command)
    }
    fmt.Printf("\n")
  }
}



// _______________________________________________________________________
func main() {
  loop()
}

